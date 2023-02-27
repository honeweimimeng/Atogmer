package conf

type EngineConf struct {
	path      string
	ProtoName string `yaml:"proto"`
	ConnType  string `yaml:"conn"`
}

func BuildConf(provider Provider) *EngineConf {
	provider.Init()
	re := provider.GetResult()
	return provider.Format(re)
}

func BuildConf0(provider Provider, formatter Formatter) *EngineConf {
	provider.Init()
	re := provider.GetResult()
	return formatter.Format(re)
}
