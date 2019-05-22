<!-- markdownlint-disable MD001 -->

# log

a simple logging library for Go inspired by Dave Cheney's [Let's talk about logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging) blog post.

## installation

```bash
go get github.com/timbasel/go-log
```

## usage

#### basic

```go
import "github.com/timbasel/go-log"

func main() {
  // message for the developer debugging the application
  log.Debug("debug message")

  // message for the user running the application
  log.Info("info message")

  // message in case the application encountered an unhandleable error and will terminate after messaging
  // this should effectively only be used in the main.main function after bubbling up the unhandled error
  log.Error("error message")
}
```

the logging functions can handle different types of messages defined by a postfix identifier (e.g. `log.DebugE` for passing an error type or `log.InfoF` for a formatted string)

#### configuration

```go
import "github.com/timbasel/go-log"

func main() {
  // set the location(s) where the log is written to (default: os.Stdout)
  log.SetOutput(os.Stdout, someIOWriter)

  // CONCEPT: set the output format of the log (default: &log.DefaultFormatter{})
  log.SetFormatter(&log.Formatter{})

  // set if debug messages should be logged (default: false)
  log.SetDebug(true)

  // CONCEPT: set the list of packages/functions that should be excluded from debug logging (default: [])
  log.SetDebugBlacklist("main.Hello")

  // CONCEPT: set the list of packages/functions that should exclusivly have debug logging (default: [])
  log.SetDebugBlacklist("main.World")
}
```

## license

apache license 2.0 © Tim Basel
