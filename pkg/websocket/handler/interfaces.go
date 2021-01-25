package handler

type Handler interface {
	Close()
	Handler(data []byte) (res []byte, err error)
	RegisterConnWriteChan(ch chan<- []byte)
	RegisterConnClose(do func())
	RegisterConnPing(do func())
}
