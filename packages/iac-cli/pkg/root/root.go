/*
 * Copyright 2023 Red Hat, Inc. and/or its affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package root

import (
	"fmt"
	"github.com/kiegroup/iac-tools/packages/iac-cli/pkg/command"
	"github.com/kiegroup/iac-tools/packages/iac-cli/pkg/common"
	"github.com/ory/viper"
	"github.com/spf13/cobra"
)

type RootCmdConfig struct {
	Name    string
	Version string
}

func NewRootCommand(cfg RootCmdConfig) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   cfg.Name,
		Short: "IaC - Ansible Plans CLI",
		Long:  "Infrastructure as Code - Ansible Plans",
	}

	viper.AutomaticEnv()

	cmd.Version = cfg.Version
	cmd.SetVersionTemplate(`{{printf "%s\n" .Version}}`)

	cmd.AddCommand(command.NewPlanCommand())
	cmd.AddCommand(command.NewApplyCommand())

	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		runRootHelp(cmd, args)
	})

	return cmd
}

func runRootHelp(cmd *cobra.Command, args []string) {
	tpl := common.GetTemplate(cmd, "root")
	var data = struct {
		Name string
	}{
		Name: cmd.Root().Use,
	}

	if err := tpl.Execute(cmd.OutOrStdout(), data); err != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "unable to display help text: %v", err)
	}
}
