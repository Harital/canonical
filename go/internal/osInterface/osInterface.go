package osInterface

import "os"

//go:generate mockery --name OsInterface -inpkg
type OsInterface interface {
	OpenFile(filePath string, flag int, perm os.FileMode) (*os.File, error)
	Remove(filepath string) error
}

type RealOs struct {
}

func (r *RealOs) OpenFile(filePath string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(filePath, flag, perm)
}

func (r *RealOs) Remove(filepath string) error {
	return os.Remove(filepath)
}