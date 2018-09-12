package rayUtils

import (
	"fmt"
	"os"
)

func Exit() {
	fmt.Print("Exiting program...\n")
	os.Exit(0)
}

func ExitFatal(str string, args ...interface{}) {
	rayUtilsGlobals.Logger.LogFatal(str, args...)
	Exit()
}

func ExitPrint(str string) {
	fmt.Println(str)
	Exit()
}

func ExitPrintf(str string, args ...interface{}) {
	fmt.Printf(str, args...)
	Exit()
}
