package handler

type Interface interface {
	RegisterConnWriteChan(ch chan<- []byte)
	RegisterConnClose(do func())
	RegisterConnPing(do func())
	Handler(data []byte) (res []byte, err error)
	Close()
}
