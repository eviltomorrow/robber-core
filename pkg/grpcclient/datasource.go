package grpcclient

import (
	"context"
	"fmt"
	"time"

	"github.com/eviltomorrow/robber-core/pkg/grpclb"
	pbdatasource "github.com/eviltomorrow/robber-datasource/pkg/pb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
)

var (
	KeyDatasource = "grpclb/service/database"
)

func NewClientForDatasource() (pbdatasource.ServiceClient, func(), error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   EtcdEndpoints,
		DialTimeout: 5 * time.Second,
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
		return nil, nil, err
	}

	builder := &grpclb.Builder{
		Client: cli,
	}
	resolver.Register(builder)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	target := fmt.Sprintf("etcd:///%s", KeyDatasource)
	conn, err := grpc.DialContext(
		ctx,
		target,
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, nil, err
	}

	return pbdatasource.NewServiceClient(conn), func() { conn.Close() }, nil
}
