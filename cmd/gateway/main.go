package main

import (
	"flag"
	"github.com/nevercase/k8s-controller-custom-resource/pkg/signals"
	"github.com/nevercase/lllidan/pkg/config"
	"github.com/nevercase/lllidan/pkg/gateway"
	"k8s.io/klog/v2"
)

func main() {
	var conf = flag.String("configPath", "conf.yml", "configuration file path")
	klog.InitFlags(nil)
	flag.Parse()
	stopCh := signals.SetupSignalHandler()
	handler := make(chan []byte, 1024)
	gateway.NewClientWithRecover(config.Init(*conf), stopCh, handler)
}