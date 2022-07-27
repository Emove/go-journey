package conf

type ServiceConfig struct {
	Conn conn `yaml:"conn"`
}

type conn struct {
	MaxCon  int32 `yaml:"max_con"`
	MaxIdle int32 `yaml:"max_idle"`
}
