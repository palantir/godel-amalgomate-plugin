// Copyright (c) 2018 Palantir Technologies Inc. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/palantir/godel/v2/framework/pluginapi/v2/pluginapi"

	"github.com/palantir/godel-amalgomate-plugin/cmd"
)

func main() {
	if ok := pluginapi.InfoCmd(os.Args, os.Stdout, cmd.PluginInfo); ok {
		return
	}
	os.Exit(cmd.Execute())
}
