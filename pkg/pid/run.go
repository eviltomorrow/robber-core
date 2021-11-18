package pid

import "os"

var lock fileLock

func CreatePidFile(path string) (err error) {
	if lock != nil {
		panic("pid file lock is not nil")
	}
	lock, err = newFileLock(path, false)
	if err != nil {
		return
	}
	return nil
}

func DestroyFile() error {
	if lock == nil {
		return nil
	}
	if err := lock.release(); err != nil {
		return err
	}
	return os.Remove(lock.path())
}
