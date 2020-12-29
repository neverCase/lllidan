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
	ListenPort        int               `json:"listenPort" yaml:"listenPort"`
	KubernetesService KubernetesService `json:"kubernetesService" yaml:"kubernetesService"`
}

type Gateway struct {
	ListenPort        int               `json:"listenPort" yaml:"listenPort"`
	KubernetesService KubernetesService `json:"kubernetesService" yaml:"kubernetesService"`
	Routes            []Route           `json:"routes" yaml:"routes"`
}

type Route struct {
	Name       string      `json:"name" yaml:"name"`
	Scheme     string      `json:"scheme" yaml:"scheme"`
	Host       string      `json:"host" yaml:"host"`
	Path       string      `json:"path" yaml:"path"`
	Parameters []Parameter `json:"parameters" yaml:"parameters"`
}

type Parameter struct {
	Key          string `json:"key" yaml:"key"`
	IsRequired   bool   `json:"isRequired" yaml:"isRequired"`
	IsDynamic    bool   `json:"isDynamic" yaml:"isDynamic"`
	DefaultValue string `json:"defaultValue" yaml:"defaultValue"`
}

type Config struct {
	Logic   Logic   `yaml:"Logic"`
	Gateway Gateway `yaml:"Gateway"`
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
