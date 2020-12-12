# lllidan
Distributed push server written in go


## todo
##### logic
- manage all the gateway client connections
- push notifications and discoveries to all the gateway clients
- provide websocket api for demonstrating the status of the cluster at the real time

##### gateway
- manage websocket app clients
- report status to the remote logic server
- transport notifications to the other gateway clients