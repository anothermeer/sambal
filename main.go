/*
Copyright © 2026 anothermeer <me@melons.cc>
*/
package main

import (
	"os"

	"github.com/anothermeer/sambal/cmd"
	"github.com/anothermeer/sambal/internal/core/utils"
	"github.com/anothermeer/sambal/internal/gui"
)

func main() {
	// --gui flag
	for _, arg := range os.Args[1:] {
		if arg == "--gui" {
			gui.Run()
			return
		}
	}

	// double click detection
	if len(os.Args) == 1 && !utils.IsRunningInTerminal() {
		gui.Run()
		return
	}

	//finally CLI
	cmd.Execute()
}
