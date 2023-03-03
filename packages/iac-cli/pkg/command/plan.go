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
	"strings"
	"time"
)

type CreatePlanConfig struct {
	ResourceName string
}

func NewPlanCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "plan",
		Short: "Generate the plan for a resource",
		Long: `
	 Generate the changes plan required by a given resource configuration.
		 `,
		Example: `
	 # Generate the plan for a resource
	 {{.Name}} plan --resource my-resource.iac.yaml
		 `,
		SuggestFor: []string{"pln", "plans"}, //nolint:misspell
		PreRunE:    common.BindEnv("resource"),
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runPlan(cmd, args)
	}
	cmd.Flags().StringP("resource", "r", "", "Generate the plan for a resource.")
	cmd.SetHelpFunc(common.DefaultTemplatedHelp)

	return cmd
}

func runPlan(cmd *cobra.Command, args []string) error {
	start := time.Now()

	cfg, err := createPlanConfig(cmd)
	if err != nil {
		return fmt.Errorf("initializing create config: %w", err)
	}

	fmt.Printf("Creating a plan for: %s \n", cfg.ResourceName)

	if err := generatePlan(cfg); err != nil {
		return err
	}

	finish := time.Since(start)

	fmt.Printf("🚀 Plan creation took: %f seconds \n", finish.Seconds())
	return nil
}

func generatePlan(cfg CreatePlanConfig) error {
	exists, err := common.CheckIfPathExists(cfg.ResourceName)
	if err != nil || !exists {
		return fmt.Errorf("resource file with name \"%s\" not found", cfg.ResourceName)
	}

	planName := getPlanName(cfg.ResourceName)
	CreatePlan(planName)
	return nil
}

func getPlanName(resourceName string) string {
	name := strings.TrimSuffix(resourceName, ".yaml")
	currentDate := time.Now().Format(time.RFC3339)
	newFilename := name + "_" + currentDate + ".plan.sw.json"
	return newFilename
}

func createPlanConfig(cmd *cobra.Command) (cfg CreatePlanConfig, err error) {
	if !viper.IsSet("resource") {
		err = fmt.Errorf("missing required flag: --resource")
		return CreatePlanConfig{}, err
	}
	cfg = CreatePlanConfig{
		ResourceName: viper.GetString("resource"),
	}
	return cfg, nil
}
