package rayUtils

import (
	"fmt"
	"os"
)

func Exit() {
	fmt.Print("Exiting program...\n")
	os.Exit(0)
}
