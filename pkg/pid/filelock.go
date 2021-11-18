package pid

type fileLock interface {
	path() string
	release() error
}
