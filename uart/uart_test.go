package uart

import (
	"io"
	"github.com/jacobsa/go-serial/serial"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockReadWriteCloser struct {
	mockRead func()
	mockWrite func()
	mockClose func()
}

func (s MockReadWriteCloser) Read([]byte) (int, error) {
	if s.mockRead != nil {
		s.mockRead()
	}
	return 0, nil
}
func (s MockReadWriteCloser) Write([]byte) (int, error) {
	if s.mockWrite != nil {
		s.mockWrite()
	}
	return 0, nil
}
func (s MockReadWriteCloser) Close() error {
    if s.mockClose != nil {
		s.mockClose()
	}
	return nil
}

func TestFindDevicePathNotFound(t *testing.T) {
	originalDevicePath := devicePath
    defer func () { devicePath = originalDevicePath }()
    devicePath = "/does/not/exist"
	
	port, err := findDevice()
	assert.Nil(t, port)
	assert.Equal(t, "Device path not found", err.Error())
}

func TestFindDeviceNotFound(t *testing.T) {
	port, err := findDevice()
	assert.Nil(t, port)
	assert.Equal(t, "Device not found", err.Error())
}

func TestFindDeviceFound(t *testing.T) {
	originalSerialOpen := serialOpen
	defer func () { serialOpen = originalSerialOpen }()
	mockPort := MockReadWriteCloser{}
    serialOpen = func(options serial.OpenOptions) (io.ReadWriteCloser, error) { return mockPort, nil }
	
	port, err := findDevice()
	assert.Equal(t, mockPort, port)
	assert.Nil(t, err)
}