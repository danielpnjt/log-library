# Log Library

This is a simple logging library written in Go.

## Installation

To install the library, run:

```sh
go get github.com/danielpnjt/log-library
```

## Usage

Here is a basic example of how to use the library:

```go
package main

import (
	"context"
	"net/http"

	"github.com/danielpnjt/log-library/logger"
	"github.com/sirupsen/logrus"
)

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	// Create a context with log fields
	ctx := context.Background()
	ctx = logger.WithLogFields(ctx, logrus.Fields{
		"user_id":    "123",
		"request_id": "abcd-1234",
	})

	// Create a response
	response := logger.Response{
		Code: "200",
		Desc: "Success",
		Data: map[string]string{"message": "Hello, world!"},
	}

	// Add the response to the context
	ctx = logger.WithResponse(ctx, response)

	// Log and send the response
	logger.LogAndSendResponse(ctx, w)
}

func main() {
	// Initialize the logger
	config := &logger.Config{
		LogLevel: "info"
	}
	if err := logger.InitLogger(config); err != nil {
		panic(err)
	}

	http.HandleFunc("/", ExampleHandler)
	http.ListenAndServe(":8080", nil)
}

```

## Features

- Simple and easy to use
- Supports different log levels (Info, Warning, Error)
- Customizable output formats

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## Contact

For any questions or suggestions, feel free to contact me at [daniel.pnjt@gmail.com](mailto:daniel.pnjt@gmail.com).