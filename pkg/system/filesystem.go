package system

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
)

const (
	statFailedErr    = "stat path %s failed: %v"
	pathIsEmptyErr   = "given path is empty"
	pathNotExistsErr = "given path %s is not exist"
	notDirectoryErr  = "given path %s is not a directory"
	notFileErr       = "given path %s is not a file"
)

func IsDirExists(dir string) (bool, error) {
	if len(dir) == 0 {
		return false, fmt.Errorf(pathIsEmptyErr)
	}

	dirInfo, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, fmt.Errorf(statFailedErr, dir, err)
	}
	if !dirInfo.IsDir() {
		return false, fmt.Errorf(notDirectoryErr, dir)
	}

	return true, nil
}

func IsDirWritable(dir string) (bool, error) {
	if len(dir) == 0 {
		return false, fmt.Errorf(pathIsEmptyErr)
	}

	exists, err := IsDirExists(dir)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, fmt.Errorf(pathNotExistsErr, dir)
	}

	_, err = os.Stat(dir)
	if err != nil {
		return false, fmt.Errorf(statFailedErr, dir, err)
	}

	return unix.Access(dir, unix.W_OK) == nil, nil
}

func IsFileExists(file string) (bool, error) {
	if len(file) == 0 {
		return false, fmt.Errorf(pathIsEmptyErr)
	}

	fileInfo, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, fmt.Errorf(statFailedErr, file, err)
	}

	if fileInfo.IsDir() {
		return false, fmt.Errorf(notFileErr, file)
	}

	return true, nil
}
