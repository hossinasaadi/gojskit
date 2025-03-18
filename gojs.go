package gojs

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
)

type Core struct {
	VM *goja.Runtime
}

func (c *Core) EvaluateScript(script string) {

	loop := eventloop.NewEventLoop()
	loop.Run(func(vm *goja.Runtime) {
		c.VM = vm
		_, err := c.VM.RunString(script)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Script runs successfuly.")
		}
	})
}

func (c *Core) CallFunc(function string, jsonParams string) string {
	assert, ok := goja.AssertFunction(c.VM.Get(function))
	if !ok {
		fmt.Println(function + " Not a function")
		return "err"
	}
	gojaValueList := c.ProcessParams(jsonParams)

	res, err := assert(goja.Undefined(), gojaValueList...)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return res.String()
}

func (c *Core) ProcessParams(jsonParams string) []goja.Value {
	var params []any // Decode JSON into a slice
	err := json.Unmarshal([]byte(jsonParams), &params)
	if err != nil {
		return nil
	}
	gojaValueList := []goja.Value{}
	for _, v := range params {
		gojaValueList = append(gojaValueList, c.VM.ToValue(v))
	}

	return gojaValueList
}
