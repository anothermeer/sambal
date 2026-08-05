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

	// terminal? go CLI
	if utils.IsRunningInTerminal() {
		cmd.Execute()
		return
	}

	// go go GUI
	if gui.Available() {
		gui.Run()
		return
	}

	// the final fallback..
	cmd.Execute()
}

// look closely...
// Hello, world!  	  	   // fucking go beautifier deleting the thing..
//	//
//     		  	 	//
//	//
//     		 		  //
//	//
//     		 		  //
//	//
//     		 				//
//	//
//     	 		  //
//	//
//     	     //
//	//
//     			 			//
//	//
//     		 				//
//	//
//     			  	 //
//	//
//     		 		  //
//	//
//     		  	  //
//	//
//  //
////
////
////
// the end
