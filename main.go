/**
 * main.go
 *
 * Copyright (c) 2017 Forest Hoffman. All Rights Reserved.
 * License: MIT License (see the included LICENSE file) or download at
 *     https://raw.githubusercontent.com/foresthoffman/reaper/master/LICENSE
 */

package main

import (
	"flag"
	"fmt"
	"github.com/foresthoffman/reap"
)

const version string = "1.0.0"

func main() {

	// setup CLI flags
	cmdFlag := flag.String("cmd", "", "The command that will be run by the middle-man process.")
	argsFlag := flag.String("args", "", "The arguments for the command provided by the -cmd flag. Optional.")
	helpFlag := flag.Bool("h", false, "Displays the help menu.")
	versionFlag := flag.Bool("v", false, "Displays the executable's current version.")
	flag.Parse()

	if *helpFlag {
		flag.Usage()
		return
	}
	if *versionFlag {
		fmt.Printf("reaper version: %s\n", version)
		return
	}

	if "" == *cmdFlag {
		fmt.Print("--cmd flag was empty")
		return
	}

	// run the provided command with the provided arguments
	pid, err := reap.Exec(*cmdFlag, *argsFlag)

	// prints an error or a valid PID for the parent process to pick up from standard output
	if nil != err {
		fmt.Print(err)
	} else {
		fmt.Print(pid)
	}
}
