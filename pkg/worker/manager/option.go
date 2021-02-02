package manager

import (
	"fmt"
	"time"
)

type MessageOption struct {
	ReadChan         chan []byte
	WriteChan        chan []byte
	writeTimeoutInMs int64
}

func (mo *MessageOption) Read() <-chan []byte {
	return mo.ReadChan
}

func (mo *MessageOption) Write(in []byte) error {
	select {
	case <-time.After(time.Millisecond * time.Duration(mo.writeTimeoutInMs)):
		return fmt.Errorf("MessageOption Write failed due to write to channel timeout:%v ms", mo.writeTimeoutInMs)
	case mo.WriteChan <- in:
	}
	return nil
}

func (mo *MessageOption) writeToReadChan(in []byte) error {
	select {
	case <-time.After(time.Millisecond * time.Duration(mo.writeTimeoutInMs)):
		return fmt.Errorf("MessageOption writeToReadChan failed due to write to channel timeout:%v ms", mo.writeTimeoutInMs)
	case mo.ReadChan <- in:
	}
	return nil
}
