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
	ServiceAPIGatewayRegister ServiceAPI = "gatewayRegister"
	ServiceAPIGatewayList     ServiceAPI = "gatewayList"
	ServiceAPIWorkerRegister  ServiceAPI = "workerRegister"
	ServiceAPIWorkerList      ServiceAPI = "workerList"
	ServiceAPIMultiplex       ServiceAPI = "multiplex"
	// ServiceAPIKickAddress was only used in client
	ServiceAPIKickAddress ServiceAPI = "kick"
)
