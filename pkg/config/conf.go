package config

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
