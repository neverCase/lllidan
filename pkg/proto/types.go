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

type GatewayList struct {
	Items []Gateway `json:"items" protobuf:"bytes,1,opt,name=items"`
}

type Gateway struct {
	Hostname string `json:"hostname" protobuf:"bytes,1,opt,name=hostname"`
	Ip       string `json:"ip" protobuf:"bytes,2,opt,name=ip"`
	Port     int32  `json:"port" protobuf:"varint,3,opt,name=port"`
	NodePort int32  `json:"nodePort" protobuf:"varint,4,opt,name=nodePort"`
}
