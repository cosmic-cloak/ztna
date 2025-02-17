/*
	Copyright NetFoundry Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package create

import (
	"io"

	cmdhelper "ztna-core/ztna/ztna/cmd/helpers"
	"ztna-core/ztna/ztna/cmd/pki"
	"ztna-core/ztna/ztna/cmd/templates"
	"github.com/spf13/cobra"
)

var (
	createLong = templates.LongDesc(`
		Creates a new Ziti resource.

	`)
)

// NewCmdCreate creates a command object for the "create" command
func NewCmdCreate(out io.Writer, errOut io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new resource",
		Long:  createLong,
		Run: func(cmd *cobra.Command, args []string) {
			cmdhelper.CheckErr(cmd.Help())
		},
	}

	cmd.AddCommand(NewCmdCreateConfig())

	cmd.AddCommand(pki.NewCmdPKICreateCA(out, errOut))
	cmd.AddCommand(pki.NewCmdPKICreateIntermediate(out, errOut))
	cmd.AddCommand(pki.NewCmdPKICreateServer(out, errOut))
	cmd.AddCommand(pki.NewCmdPKICreateClient(out, errOut))
	cmd.AddCommand(pki.NewCmdPKICreateCSR(out, errOut))

	return cmd
}
