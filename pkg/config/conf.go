package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"k8s.io/klog/v2"
)

type KubernetesService struct {
	Name     string `json:"name" yaml:"name"`
	Port     int    `json:"port" yaml:"port"`
	NodePort int    `json:"nodePort" yaml:"nodePort"`
}

type Logic struct {
	ListenPort        int    `json:"listenPort" yaml:"listenPort"`
	KubernetesService string `json:"kubernetesService" yaml:"kubernetesService"`
}

type Gateway struct {
	ListenPort        int    `json:"listenPort" yaml:"listenPort"`
	KubernetesService string `json:"kubernetesService" yaml:"kubernetesService"`
}

type Config struct {
	Logic   Logic   `yaml:"GameService,flow"`
	Gateway Gateway `yaml:"GmtService,flow"`
}

func Init(configFile string) *Config {
	c := &Config{}
	var data []byte
	var err error
	if data, err = ioutil.ReadFile(configFile); err != nil {
		klog.Fatal(err)
	}
	if err := yaml.Unmarshal(data, c); err != nil {
		klog.Fatal(err)
	}
	return c
}