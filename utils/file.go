package utils

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	// "golang.org/x/sys/windows"
	// "unsafe"
)

// "golang.org/x/sys/windows"
// DirSize 获取一个目录的大小
func DirSize(dirPath string) (int64, error) {
	var size int64
	err := filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size, err
}

type DiskStatus struct {
	All  uint64
	Used uint64
	Free uint64
}

// AvailableDiskSize 获取磁盘剩余可用空间大小
// func AvailableDiskSize() (uint64, error) {
// 	var disk DiskStatus
// 	wd, err := syscall.Getwd()

// 	if err != nil {
// 		return 0, err
// 	}

// 	h := windows.MustLoadDLL("kernel32.dll")
// 	c := h.MustFindProc("GetDiskFreeSpaceExW")
// 	lpFreeBytesAvailable := uint64(0)
// 	lpTotalNumberOfBytes := uint64(0)
// 	lpTotalNumberOfFreeBytes := uint64(0)
// 	r1, r2, err := c.Call(uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(wd))),
// 		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
// 		uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
// 		uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)))
// 	if r1 == r2 {

// 	}
// 	// 这里正常返回情况下 err 也不为 nil
// 	// if err != nil {

// 	// }

// 	disk.All = lpTotalNumberOfBytes
// 	disk.Free = lpTotalNumberOfFreeBytes
// 	disk.Used = lpFreeBytesAvailable
// 	return disk.Free, nil
// }

// // AvailableDiskSize 获取磁盘剩余可用空间大小
func AvailableDiskSize() (uint64, error) {
	wd, err := syscall.Getwd()
	if err != nil {
		return 0, err
	}
	var stat syscall.Statfs_t
	if err = syscall.Statfs(wd, &stat); err != nil {
		return 0, err
	}
	return stat.Bavail * uint64(stat.Bsize), nil
}

// CopyDir 拷贝数据目录
func CopyDir(src, dest string, exclude []string) error {
	// 目标目标不存在则创建
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		if err := os.MkdirAll(dest, os.ModePerm); err != nil {
			return err
		}
	}

	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		fileName := strings.Replace(path, src, "", 1)
		if fileName == "" {
			return nil
		}

		for _, e := range exclude {
			matched, err := filepath.Match(e, info.Name())
			if err != nil {
				return err
			}
			if matched {
				return nil
			}
		}

		if info.IsDir() {
			return os.MkdirAll(filepath.Join(dest, fileName), info.Mode())
		}

		data, err := os.ReadFile(filepath.Join(src, fileName))
		if err != nil {
			return err
		}
		return os.WriteFile(filepath.Join(dest, fileName), data, info.Mode())
	})
}
