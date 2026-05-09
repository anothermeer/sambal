package utils

import (
	"os"

	"github.com/mattn/go-isatty"
)

func IsRunningInTerminal() bool {
	return isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
}
