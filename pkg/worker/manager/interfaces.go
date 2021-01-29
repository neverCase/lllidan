package manager

import (
	"fmt"
	"time"
)

type MessageHandler interface {
	Read() <-chan interface{}
	Write(in interface{}) error
}

type MessageOption struct {
	ReadChan         chan interface{}
	WriteChan        chan interface{}
	writeTimeoutInMs int64
}

func (mo *MessageOption) Read() <-chan interface{} {
	return mo.ReadChan
}

func (mo *MessageOption) Write(in interface{}) error {
	select {
	case <-time.After(time.Millisecond * time.Duration(mo.writeTimeoutInMs)):
		return fmt.Errorf("MessageOption Write failed due to write to channel timeout:%v ms", mo.writeTimeoutInMs)
	case mo.WriteChan <- in:
	}
	return nil
}

func (mo *MessageOption) writeToReadChan(in interface{}) error {
	select {
	case <-time.After(time.Millisecond * time.Duration(mo.writeTimeoutInMs)):
		return fmt.Errorf("MessageOption writeToReadChan failed due to write to channel timeout:%v ms", mo.writeTimeoutInMs)
	case mo.ReadChan <- in:
	}
	return nil
}
