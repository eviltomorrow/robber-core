package grpclb

import (
	"context"
	"fmt"
	"time"

	"github.com/eviltomorrow/robber-core/pkg/zlog"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

var timeout = 5 * time.Second

func init() {

}

func Register(service string, host string, port int, endpoints []string, ttl int64) (func(), error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: timeout,
		LogConfig: &zap.Config{
			Level:            zap.NewAtomicLevelAt(zap.ErrorLevel),
			Development:      false,
			Encoding:         "json",
			EncoderConfig:    zap.NewProductionEncoderConfig(),
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
		},
	})
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for _, endpoint := range endpoints {
		_, err = client.Status(ctx, endpoint)
		if err != nil {
			return nil, err
		}
	}

	leaseResp, err := client.Grant(context.Background(), ttl)
	if err != nil {
		return nil, err
	}
	var leaseID = &leaseResp.ID

	key, value := fmt.Sprintf("/%s/%s:%d", service, host, port), fmt.Sprintf("%s:%d", host, port)
	_, err = client.Put(context.Background(), key, value, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		return nil, err
	}

	keepAlive, err := client.KeepAlive(context.Background(), leaseResp.ID)
	if err != nil {
		return nil, err
	}
	go func() {
	keep:
		for {
			select {
			case <-client.Ctx().Done():
				return
			case k, ok := <-keepAlive:
				if !ok {
					break keep
				}
				if k != nil {
					_ = k
				}
			}
		}

	release:
		zlog.Error("Etcd status is abnormal: register service info retrying", zap.String("key", key), zap.String("value", value))
		leaseResp, err := client.Grant(context.Background(), ttl)
		if err != nil {
			goto release
		}

		key, value := fmt.Sprintf("/%s/%s:%d", service, host, port), fmt.Sprintf("%s:%d", host, port)
		_, err = client.Put(context.Background(), key, value, clientv3.WithLease(leaseResp.ID))
		if err != nil {
			goto release
		}

		keepAlive, err = client.KeepAlive(context.Background(), leaseResp.ID)
		if err != nil {
			goto release
		}
		leaseID = &leaseResp.ID
		zlog.Info("Etcd status is ok: register service info complete", zap.String("key", key), zap.String("value", value))
		goto keep
	}()
	close := func() {
		_, _ = client.Revoke(ctx, *leaseID)
	}

	return close, nil
}
