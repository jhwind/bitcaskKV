package fio

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func destroyFile(name string) {
	if err := os.RemoveAll(name); err != nil {
		panic(err)
	}
}

func TestNewFileIOManager(t *testing.T) {
	// 要么 ./ 要么加绝对路径
	path := filepath.Join("./", "a.data")
	// path := filepath.Join("c:/Users/thinkbook/Nutstore/1/Project/go/bitcaskKV/tmp", "a.data")
	fio, err := NewFileIOManager(path)
	// defer destroyFile(path)

	assert.Nil(t, err)
	assert.NotNil(t, fio)
}

func TestFileIO_Write(t *testing.T) {
	path := filepath.Join("./", "a.data")
	fio, err := NewFileIOManager(path)
	// defer destroyFile(path)

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	n, err := fio.Write([]byte(""))
	assert.Equal(t, 0, n)
	assert.Nil(t, err)

	n, err = fio.Write([]byte("bitcaskKV"))
	assert.Equal(t, 9, n)
	assert.Nil(t, err)

	n, err = fio.Write([]byte("storage"))
	assert.Equal(t, 7, n)
	assert.Nil(t, err)
}

func TestFileIO_Read(t *testing.T) {
	path := filepath.Join("./", "b.data")
	fio, err := NewFileIOManager(path)
	// defer destroyFile(path)

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	_, err = fio.Write([]byte("key-a"))
	assert.Nil(t, err)

	_, err = fio.Write([]byte("key-b"))
	assert.Nil(t, err)

	b := make([]byte, 5)
	n, err := fio.Read(b, 0)

	assert.Equal(t, 5, n)
	assert.Equal(t, []byte("key-a"), b)

	b2 := make([]byte, 5)
	n, err = fio.Read(b2, 5)
	assert.Equal(t, []byte("key-b"), b2)
}

func TestFileIO_Sync(t *testing.T) {
	path := filepath.Join("./", "b.data")
	fio, err := NewFileIOManager(path)
	// defer destroyFile(path)

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	// err = fio.Sync()
	assert.Nil(t, err)
}

func TestFileIO_Close(t *testing.T) {
	path := filepath.Join("./", "b.data")
	fio, err := NewFileIOManager(path)
	// defer destroyFile(path)

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	err = fio.Close()
	assert.Nil(t, err)
}
