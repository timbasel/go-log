package log

import (
	"runtime"
	"strings"

	"github.com/benbjohnson/clock"
)

func getTimestamp(clock clock.Clock, layout string) (timestamp string) {
	return clock.Now().Format(layout)
}

func getFullFunctionName() (name string) {
	functionName, packageName := findInitialCaller()
	if functionName == "" || packageName == "" {
		return ""
	}

	index := strings.LastIndex(packageName, "/")
	return packageName[index+1:] + "." + functionName
}

func getFunctionName() (name string) {
	functionName, _ := findInitialCaller()
	if functionName == "" {
		return ""
	}

	return functionName
}

func getPackageName() (name string) {
	_, packageName := findInitialCaller()
	if packageName == "" {
		return ""
	}

	return packageName
}

func findInitialCaller() (functionName string, packageName string) {
	_, currentPackageName := getCallerInformation(0)
	i := 1
	for {
		functionName, packageName = getCallerInformation(i)
		if packageName != currentPackageName || functionName == "" || packageName == "" {
			return functionName, packageName
		}
		i++
	}
}

func getCallerInformation(index int) (functionName string, packageName string) {
	pc, _, _, ok := runtime.Caller(index)
	details := runtime.FuncForPC(pc)
	if !ok && details == nil {
		return "", ""
	}
	parts := strings.Split(details.Name(), ".")
	packageName = ""
	functionName = parts[len(parts)-1]

	if parts[len(parts)-2][0] == '(' {
		functionName = parts[len(parts)-2] + "." + functionName
		packageName = strings.Join(parts[0:len(parts)-2], ".")
	} else {
		packageName = strings.Join(parts[0:len(parts)-1], ".")
	}

	return functionName, packageName
}
