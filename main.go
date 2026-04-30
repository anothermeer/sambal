/*
Copyright © 2026 anothermeer <me@melons.cc>
*/
package main

import (
	"os"

	"github.com/anothermeer/sambal/cmd"
)

func main() {
	// default CLI behavior (send files)
	if len(os.Args) > 1 {
		arg := os.Args[1]

		if arg != "send" && arg != "recv" && arg != "list" {
			os.Args = append([]string{os.Args[0], "send"}, os.Args[1:]...)
		}
	}

	cmd.Execute()
}
