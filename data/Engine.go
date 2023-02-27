package data

import (
	"litecluster/call"
	"litecluster/data/conf"
)

type Engine struct {
	conf *conf.EngineConf
}

func StartUp() {
	engine := &Engine{}
	yaml := (&call.YamlConf{Name: "LITE_CLUSTER_ROOT"}).FromEnv()
	engine.conf = conf.BuildConf(yaml)
}
