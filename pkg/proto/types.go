package proto

type Request struct {
	ServiceAPI ServiceAPI `json:"serviceApi" protobuf:"bytes,1,opt,name=serviceApi"`
	Data       [][]byte   `json:"data" protobuf:"bytes,2,opt,name=data"`
}

type Response struct {
	ServiceAPI ServiceAPI `json:"serviceApi" protobuf:"bytes,1,opt,name=serviceApi"`
	Code       int32      `json:"code" protobuf:"varint,2,opt,name=code"`
	Message    string     `json:"message" protobuf:"bytes,3,opt,name=message"`
	Data       [][]byte   `json:"data" protobuf:"bytes,4,opt,name=data"`
}

// Gateway
type Gateway struct {
	Hostname string `json:"hostname" protobuf:"bytes,1,opt,name=hostname"`
	Ip       string `json:"ip" protobuf:"bytes,2,opt,name=ip"`
	Port     int32  `json:"port" protobuf:"varint,3,opt,name=port"`
	NodePort int32  `json:"nodePort" protobuf:"varint,4,opt,name=nodePort"`
}

// GatewayList
type GatewayList struct {
	Items []Gateway `json:"items" protobuf:"bytes,1,opt,name=items"`
}

// Worker
type Worker struct {
	Hostname string `json:"hostname" protobuf:"bytes,1,opt,name=hostname"`
	Ip       string `json:"ip" protobuf:"bytes,2,opt,name=ip"`
	Port     int32  `json:"port" protobuf:"varint,3,opt,name=port"`
	NodePort int32  `json:"nodePort" protobuf:"varint,4,opt,name=nodePort"`
}

// WorkerList
type WorkerList struct {
	Items []Worker `json:"items" protobuf:"bytes,1,opt,name=items"`
}

type KickInfo struct {
	Hostname string `json:"hostname" protobuf:"bytes,1,opt,name=hostname"`
	ClientId string `json:"clientId" protobuf:"bytes,2,opt,name=clientId"`
	Id       int32  `json:"id" protobuf:"bytes,3,opt,name=id"`
}

// Notification
type Notification struct {
	Ids     []int32 `json:"ids" protobuf:"bytes,1,opt,name=ids"`
	Message []byte  `json:"message" protobuf:"bytes,2,opt,name=message"`
}
