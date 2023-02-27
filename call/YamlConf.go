package call

import (
	"gopkg.in/yaml.v3"
	"litecluster/call/chanl"
	"litecluster/data/conf"
	"os"
)

const FileYaml = "/conf/cluster.yaml"

type YamlConf struct {
	Name    string
	channel chanl.IOChannel
}

func (y *YamlConf) FromEnv() *YamlConf {
	ser := &chanl.ChannelService{}
	v := os.Getenv(y.Name) + FileYaml
	y.channel = ser.FileChannel(v).Build()
	return y
}

func (y *YamlConf) FromCommand() *YamlConf {
	_ = os.Args
	return y
}

func (y *YamlConf) Init() {}

func (y *YamlConf) GetResult() []byte {
	return y.channel.ReadBytes(-1)
}

func (y *YamlConf) Format(info []byte) *conf.EngineConf {
	c := &conf.EngineConf{}
	err := yaml.Unmarshal(info, c)
	if err != nil {
		panic(err)
	}
	return c
}
