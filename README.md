# Logger Library

This is a simple and standardized logging library for Go applications.

## Installation

To install the library, use the following command:

```sh
go get github.com/danielpnjt/logger
```

## Usage

First, import the logger package in your Go application:

```go
import "github.com/danielpnjt/logger"
```

### Initializing the Logger

Initialize the logger with the desired log level:

```go
logger.Init(logger.LevelInfo)
```

### Logging Messages

You can log messages at different levels:

```go
logger.Debug("This is a debug message")
logger.Info("This is an info message")
logger.Warn("This is a warning message")
logger.Error("This is an error message")
```

### Log Levels

The available log levels are:

- `logger.LevelDebug`
- `logger.LevelInfo`
- `logger.LevelWarn`
- `logger.LevelError`

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.