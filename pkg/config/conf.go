package config

type Register struct {
	ListenPort int `json:"listenPort" yaml:"listenPort"`
}

type Gateway struct {
	ListenPort int `json:"listenPort" yaml:"listenPort"`
}

type Config struct {
	Register Register `yaml:"GameService,flow"`
	Gateway  Gateway  `yaml:"GmtService,flow"`
}
