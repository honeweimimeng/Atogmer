package driver

import (
	"litecluster/utils"
	"testing"
	"time"
)

type TestEvent struct {
	name string
	id   int
	msg  interface{}
}

func (c *TestEvent) Name() string {
	return c.name
}

func (c *TestEvent) Id() int {
	return c.id
}
func (c *TestEvent) Msg() interface{} {
	return c.msg
}

func TestExecutorGroup(t *testing.T) {
	eventCtx := UseEventContext()
	eventCtx.Boot().Execute()
	time.Sleep(20 * time.Second)
}

func TestSafeMap(t *testing.T) {
	mapInstance := utils.UseInterSafeMap[string, string]()
	mapInstance.Put("name", "张三")
	v := mapInstance.Get("name")
	println(v)
	go func() {
		mapInstance.Put("name", "张三")
	}()
	go func() {
		mapInstance.Put("name", "李四")
	}()
	go func() {
		v := mapInstance.Get("name")
		println(v)
	}()
}
