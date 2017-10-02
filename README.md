## Midprocrunner

"midprocrunner" utilizes the [midproc](https://github.com/foresthoffman/midproc) package to build an executable intermediary process runner. This runner is meant to be executed by other Go programs, in order to create detached processes for them.

### Installation

Run `go get github.com/foresthoffman/midprocrunner`

In order for the `midprocrunner` executable to work, you must have `$GOPATH/bin` in your `$PATH` environment variable.

### Importing

This package is not a library. It provides an executable that can be used by other programs, therefore importing it doesn't do anything.

### Usage

Here's a simple example of the syntax:

```Go
package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
)

func main() {

	// prepare a buffer, to which the PID will be written
	var stdout bytes.Buffer

	// prepare the command
	sleepCmd := exec.Command("midprocrunner", "-cmd='sleep'", "-args='30'")
	sleepCmd.Stdout = &stdout

	// run the command
	err := sleepCmd.Run()
	if nil != err {
		panic(err)
	}

	// convert the PID string to a valid integer
	pidInt, err := strconv.ParseInt(stdout.String(), 10, 64)
	if nil != err {
		panic(err)
	}
	pid := int(pidInt)
	fmt.Printf("Created a detached process with an ID of %d!\n", pid)
}
```

_That's all, enjoy!_
