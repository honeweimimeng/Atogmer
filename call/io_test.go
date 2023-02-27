package call

import (
	"litecluster/call/chanl"
	"litecluster/data"
	"litecluster/data/conf"
	"strings"
	"testing"
)

func TestFileIOChannel(t *testing.T) {
	fileProto := &data.FileDataBuf{Path: "C:\\Users\\Administrator\\Desktop\\conf.yaml"}
	fileProto.Remove()
	nioChannel := &chanl.BufChannel{Adapter: fileProto}
	nioChannel.Init()
	info := "abcdef"
	nioChannel.Write([]byte(info))
	str := nioChannel.ReadString()
	if !strings.EqualFold(str, info) {
		t.Error("not equals")
	}
}

func TestDataConf(f *testing.T) {
	yaml := (&YamlConf{Name: "LITE_CLUSTER_ROOT"}).FromEnv()
	engine := conf.BuildConf(yaml)
	println(engine.ProtoName)
	println(engine.ConnType)
}
