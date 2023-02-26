package conf

type Formatter interface {
	Format([]byte) *EngineConf
}

type Provider interface {
	Init()
	GetResult() []byte
	Format([]byte) *EngineConf
}
