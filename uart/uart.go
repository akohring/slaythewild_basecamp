package uart

import (
	"time"
	"bytes"
	"errors"
	"strings"
	"io"
	"io/ioutil"
	"github.com/jacobsa/go-serial/serial"
	log "github.com/sirupsen/logrus"
)

var devicePath = "/dev"
var serialOpen = serial.Open

//Run the UART monitor
func Run() {
	for {
		port, name, err := findDevice()
		if(err != nil) {
			log.WithFields(log.Fields{
				"reason": err,
			  }).Warn("Failed to find device")
			time.Sleep(5000 * time.Millisecond)
		} else {
			log.WithFields(log.Fields{
				"device": name,
			  }).Info("Connected")
			readBytes(port)
		}
	}
}

func findDevice() (io.ReadWriteCloser, string, error) {
	devices, err := ioutil.ReadDir(devicePath)
	if(err != nil) {
		return nil, "", err
	}
	for _, f := range devices {
		if strings.HasPrefix(f.Name(), "ttyUSB") {
			portName := devicePath + "/" + f.Name()
			options := serial.OpenOptions {
				PortName:        portName,
				BaudRate:        57600,
				DataBits:        8,
				StopBits:        1,
				MinimumReadSize: 4,
			}
			port, err := serialOpen(options)
			if(err != nil) {
				return nil, "", err
			}
			return port, portName, nil
		}
	}
	return nil, "", errors.New("No devices available")
}

func readBytes(port io.ReadWriteCloser) {
	image := []byte{}
	buffer := []byte{}
	reconnect := false
	startCaptureTime := time.Now()
	for !reconnect {
		// TODO Does not appear to error on disconnect
		count, err := port.Read(buffer)
		if(count > 0) {
			keepBytes := true
			if(len(image) == 0) {
				if(count >= 2 && bytes.Equal(buffer[0:2], []byte{0xFF, 0xD8})) {
					startCaptureTime = time.Now()
				} else {
					keepBytes = false
				}
			}
			if(keepBytes) {
				image = append(image, buffer...)
			}
			if(len(image) > 2 && bytes.Equal(image[len(image)-2:], []byte{0xFF, 0xD9})) {
				// TODO Send notification
				image = []byte{}
			}
		}

		if(err != nil) {
			log.Error(err)
			reconnect = true
		} else if(len(image) > 0 && time.Since(startCaptureTime) > 3000) {
			image = []byte{}
			log.Warn("Abandoning partial image capture")
		} else if(len(image) == 0) {
			time.Sleep(1000 * time.Millisecond)
		}
	}
}