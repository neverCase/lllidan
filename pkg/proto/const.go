package proto

const (
	RouterDashboard = "/dashboard"
	RouterGateway   = "/gateway"
)

const (
	RouterLogicAPICloseClient = "/logic/api/close/:uid"
)

type ServiceAPI string

const (
	ServiceAPIPing            ServiceAPI = "ping"
	ServiceAPIGatewayRegister ServiceAPI = "gatewayRegister"
	ServiceAPIGatewayList     ServiceAPI = "gatewayList"
	ServiceAPIMultiplex       ServiceAPI = "multiplex"
)
