package handler

type Interface interface {
	RegisterId(id int32)
	Id() int32
	RegisterRemoveChan(ch chan<- int32)
	RegisterConnWriteChan(ch chan<- []byte)
	RegisterConnClose(do func())
	RegisterConnPing(do func())
	Ping()
	Handler(data []byte) (res []byte, err error)
	Run()
	Close()
}

type MessageHandler interface {
	Read() <-chan []byte
	Write(in []byte) error
}