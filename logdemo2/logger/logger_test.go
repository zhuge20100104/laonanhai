package logger

import (
	"fmt"
	"testing"
)

func TestLevel(t *testing.T) {
	t.Logf("%v %T\n", DebugLevel, DebugLevel)
	t.Logf("%v %T\n", FatalLevel, FatalLevel)
}

func TestGetCallerInfo(t *testing.T) {
	skip := 0
	fileName, line, funcName := GetCallerInfo(skip)
	fmt.Printf("Info of the %dth caller, %v %v %v\n", skip, fileName, line, funcName)
}
