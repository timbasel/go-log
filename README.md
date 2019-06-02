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
import log "github.com/timbasel/go-log/pkg"

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

#### configuration

```go
import log "github.com/timbasel/go-log/pkg"

func main() {
  // adds a location where the log is written to (by default os.Stdout is set)
  log.SetOutput(os.Stdout)
  
  // adds a location with a custom formatter
  log.SetFormattedOutput(os.Stdout, &log.CustomFormatter{})

  // set if debug messages should be logged (default: false)
  log.SetDebug(true)
```

## license

apache license 2.0 © Tim Basel
