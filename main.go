/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package main

import (
	"runtime"

	"github.com/IBAX-io/go-ibax/cmd"
)

func main() {
	runtime.LockOSThread()
	cmd.Execute()
}
