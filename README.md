# GoJSKit

GoJSKit is a Go package that implements JavaScript functionality in Go and exports it as an XCFramework for the Swift platform. It enables seamless JavaScript execution in Swift by leveraging Go’s efficiency and flexibility.

## Installation

### Clone the repository:

```sh
git clone https://github.com/hossinasaadi/gojskit.git
cd gojskit
```

### Build the XCFramework:

```sh
make build_apple
```

Or follow the build instructions in the repository.

## Usage

GoJSKit is designed to be used in Swift via SwiftJSKit. However, if you want to run it directly in Go:

```go
import (
	"encoding/json"
	"fmt"

	gojs "github.com/hossinasaadi/gojskit"
)

func main() {
	script := `
          function print(a,b,c) {
            console.log(a,b,c)
          }
          function printAndReturn(a,b,c) {
            console.log(a,b,c)
            return [a,b,c].join(",");
          }
        `
	context := gojs.Core{}
	context.EvaluateScript(script)
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

```

For Swift integration, refer to the [SwiftJSKit documentation](https://github.com/hossinasaadi/SwiftJSKit/).

## Credits

GoJSKit is powered by [goja](https://github.com/dop251/goja), a pure Go JavaScript engine.

## Contributing

Contributions are welcome! Feel free to open issues or pull requests.

