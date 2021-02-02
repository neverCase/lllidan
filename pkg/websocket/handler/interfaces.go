package handler

type Interface interface {
	RegisterId(id int32)
	RegisterRemoveChan(ch chan<- int32)
	RegisterConnWriteChan(ch chan<- []byte)
	RegisterConnClose(do func())
	RegisterConnPing(do func())
	Handler(data []byte) (res []byte, err error)
	Run()
	Close()
}
