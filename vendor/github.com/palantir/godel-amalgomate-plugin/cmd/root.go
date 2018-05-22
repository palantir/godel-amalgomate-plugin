// Copyright (c) 2018 Palantir Technologies Inc. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package cmd

import (
	"github.com/palantir/godel/framework/pluginapi"
	"github.com/palantir/pkg/cobracli"
	"github.com/spf13/cobra"
)

const VerifyFlagName = "verify"

var (
	debugFlagVal      bool
	projectDirFlagVal string
	configFileFlagVal string
)

var rootCmd = &cobra.Command{
	Use:   "amalgomate-plugin",
	Short: "Run amalgomate based on project configuration",
}

func Execute() int {
	return cobracli.ExecuteWithDebugVarAndDefaultParams(rootCmd, &debugFlagVal)
}

func init() {
	pluginapi.AddDebugPFlagPtr(rootCmd.PersistentFlags(), &debugFlagVal)
	pluginapi.AddProjectDirPFlagPtr(rootCmd.PersistentFlags(), &projectDirFlagVal)
	if err := rootCmd.MarkPersistentFlagRequired(pluginapi.ProjectDirFlagName); err != nil {
		panic(err)
	}
	pluginapi.AddConfigPFlagPtr(rootCmd.PersistentFlags(), &configFileFlagVal)
	if err := rootCmd.MarkPersistentFlagRequired(pluginapi.ConfigFlagName); err != nil {
		panic(err)
	}
}
