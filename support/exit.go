package support

import (
	"fmt"
	"os"
)

// Exit exits the application.
func Exit(ExitCode int) {
	exit, _ := os.OpenFile(".exit", os.O_RDWR|os.O_CREATE, 0666)
	exit.WriteString(fmt.Sprintf("%d", ExitCode))
	os.Exit(ExitCode)
}
