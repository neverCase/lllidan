package proto

const (
	RouterDashboard = "/dashboard"
	RouterGateway   = "/gateway"
	RouterWorker    = "/worker"
)

const (
	RouterLogicAPICloseClient = "/logic/api/close/:uid"
)

type ServiceAPI string

const (
	ServiceAPIPing            ServiceAPI = "ping"
	ServiceAPIPong            ServiceAPI = "Pong"
	ServiceAPIGatewayRegister ServiceAPI = "gatewayRegister"
	ServiceAPIGatewayList     ServiceAPI = "gatewayList"
	ServiceAPIGatewayConnect  ServiceAPI = "gatewayConnect"
	ServiceAPIWorkerRegister  ServiceAPI = "workerRegister"
	ServiceAPIWorkerList      ServiceAPI = "workerList"
	ServiceAPIMultiplex       ServiceAPI = "multiplex"
	// ServiceAPIKickAddress was only used in client
	ServiceAPIKickAddress ServiceAPI = "kick"
)
