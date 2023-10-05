package console

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

func Success(msg string) {
	output(msg, "green")
}

func output(message, color string) {
	fmt.Fprintln(os.Stdout, ansi.Color(message, color))
}
