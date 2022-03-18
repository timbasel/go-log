package log_test

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/timbasel/go-log/pkg/log"
)

func ExampleBasic() {
	log.SetDebugMode(true)

	// message for the developer debugging the application
	log.Debug("debug message")

	// message for the user running the application
	log.Info("info message")

	// message in case the application encountered an unhandleable error and will terminate after messaging
	// this should effectively only be used in the main.main function after bubbling up the unhandled error
	log.Error("error message")

	// Output:
}

func ExampleConfiguration() {
	// adds locations where the log is written to (by default os.Stdout is set)
	log.SetOutputs(os.Stdout)

	// adds locations with a custom formatter
	log.SetFormattedOutputs(map[io.Writer]log.Formatter{os.Stdout: log.NewRawFormatter()})

	// set if debug messages should be logged (default: false)
	log.SetDebugMode(true)

	// blacklist or whitelist functions and packages from debug output
	log.BlacklistFunctions("main")
	log.WhitelistPackages("github.com/username/project")

	// remove entries from blacklist or whitelist
	log.ClearBlacklist()
}

func ExampleCustomLogger() {
	logger := log.NewDefaultLogger()

	logger.ClearOutputs()
	logger.SetOutputs(os.Stderr)

	logger.Info("Hello World")
}

type CustomFormatter struct{}

func (f *CustomFormatter) Format(level log.Level, msg string) (formattedMsg string) {
	return "> " + msg + " [" + level.String() + "]\n"
}

func ExampleFormatters() {
	logger := log.NewLogger()
	logger.SetDebugMode(true)

	file, _ := ioutil.TempFile("/tmp/", "log_*.json")
	logger.SetFormattedOutputs(map[io.Writer]log.Formatter{
		os.Stdout: &CustomFormatter{},
		os.Stderr: log.NewDefaultFormatter(),
		file:      log.NewJSONFormatter(),
	})

	logger.Debug("Hello World")

	// output to stdout:
	//    > Hello World [DEBUG]
	// output to stderr:
	//    2019-06-03 17:24:10 <log_test.ExampleFormatters> DEBUG: Hello World
	// output to the temporary file:
	//    {"function":"ExampleFormatters","level":"DEBUG","msg":"Hello World","package":"github.com/timbasel/go-log/pkg/log_test","time":"2019-06-03T17:23:17+02:00"}
}
