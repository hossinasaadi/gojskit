package main

import (
	"encoding/json"
	"fmt"

	gojs "github.com/hossinasaadi/gojskit"
)

func main() {
	context := gojs.Core{}
	context.EvaluateScript(`
    function print(a,b,c) {
        console.log(a,b,c)
    }
    function printAndReturn(a,b,c) {
        console.log(a,b,c)
        return [a,b,c].join(",");
    }
	
	`)
	params := []interface{}{"hello", 42, true}

	// Convert to JSON
	jsonBytes, err := json.Marshal(params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert bytes to string
	jsonStringParams := string(jsonBytes)

	result := context.CallFunc("printAndReturn", jsonStringParams)
	fmt.Println("function output:", result)
}
