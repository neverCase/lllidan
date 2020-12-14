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
	APIPing            ServiceAPI = "ping"
	APIGatewayRegister ServiceAPI = "gatewayRegister"
)
