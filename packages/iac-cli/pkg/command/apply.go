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

package command

import (
	"fmt"
	"github.com/kiegroup/iac-tools/packages/iac-cli/pkg/common"
	"github.com/ory/viper"
	"github.com/spf13/cobra"
	"time"
)

type CreateApplyConfig struct {
	PlanName string
}

func NewApplyCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "apply",
		Short: "Create or update your infrastructure",
		Long: `
	 Apply a given plan to your target infrastructure.
		 `,
		Example: `
	 # Apply the plan for a target infrastructure
	 {{.Name}} apply --plan my-resource.iac.yaml
		 `,
		SuggestFor: []string{"appl", "app"}, //nolint:misspell
		PreRunE:    common.BindEnv("plan"),
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runApply(cmd, args)
	}
	cmd.Flags().StringP("plan", "r", "", "Apply the plan for a target infrastructure.")
	cmd.SetHelpFunc(common.DefaultTemplatedHelp)

	return cmd
}

func runApply(cmd *cobra.Command, args []string) error {
	start := time.Now()

	cfg, err := createApplyConfig(cmd)
	if err != nil {
		return fmt.Errorf("initializing create config: %w", err)
	}

	fmt.Printf("Executing plan for: %s \n", cfg.PlanName)

	if err := executeApply(cfg); err != nil {
		return err
	}

	finish := time.Since(start)

	fmt.Printf("🚀 Plan execution took: %f seconds \n", finish.Seconds())
	return nil
}

func executeApply(cfg CreateApplyConfig) error {
	time.Sleep(10 * time.Second)
	return nil
}

func createApplyConfig(cmd *cobra.Command) (cfg CreateApplyConfig, err error) {
	if !viper.IsSet("plan") {
		err = fmt.Errorf("missing required flag: --plan")
		return CreateApplyConfig{}, err
	}
	cfg = CreateApplyConfig{
		PlanName: viper.GetString("plan"),
	}
	return cfg, nil
}
