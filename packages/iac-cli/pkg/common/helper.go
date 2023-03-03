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

package common

import (
	"fmt"
	"github.com/spf13/cobra"
	"html/template"
)

func GetTemplate(cmd *cobra.Command, name string) *template.Template {
	var (
		body = cmd.Long + "\n\n" + cmd.UsageString()
		t    = template.New(name)
		tpl  = template.Must(t.Parse(body))
	)
	return tpl
}

func DefaultTemplatedHelp(cmd *cobra.Command, args []string) {
	tpl := GetTemplate(cmd, "help")
	var data = struct{ Name string }{Name: cmd.Root().Use}

	if err := tpl.Execute(cmd.OutOrStdout(), data); err != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "unable to display help text: %v", err)
	}
}
