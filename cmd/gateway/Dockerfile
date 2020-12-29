FROM ubuntu:18.04

LABEL maintainer="lunara-developer@lunara.net"

RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install tzdata

WORKDIR /server/
COPY lllidan-gateway /server/

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN dpkg-reconfigure -f noninteractive tzdata

CMD ["/server/lllidan-gateway", "-alsologtostderr=true", "-v=4"]
