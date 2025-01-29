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

package pki

import (
	"io"

	"ztna-core/ztna/ztna/cmd/common"
	cmdhelper "ztna-core/ztna/ztna/cmd/helpers"
	"ztna-core/ztna/ztna/cmd/lets_encrypt"
	"ztna-core/ztna/ztna/cmd/templates"
	"ztna-core/ztna/ztna/pki/pki"
	"github.com/spf13/cobra"
)

// PKIOptions contains the command line options
type PKIOptions struct {
	common.CommonOptions

	Flags PKIFlags
}

type PKIFlags struct {
	PKIRoot               string
	PKIOrganization       string
	PKIOrganizationalUnit string
	PKICountry            string
	PKILocality           string
	PKIProvince           string
	CAFile                string
	CAName                string
	CommonName            string
	CAExpire              int
	CAMaxPath             int
	CAPrivateKeySize      int
	IntermediateFile      string
	IntermediateName      string
	ServerFile            string
	ServerName            string
	ClientFile            string
	ClientName            string
	KeyFile               string
	CSRFile               string
	CSRName               string
	KeyName               string
	DNSName               []string
	IP                    []string
	Email                 []string
	PKI                   *pki.ZitiPKI
	SpiffeID              string
	AllowOverwrite        bool
	EcCurve               string
}

var (
	pkiLong = templates.LongDesc(`
Provide the components needed to manage a Ziti PKI.
	`)
)

// NewCmdPKI PKIs a command object for the "PKI" command
func NewCmdPKI(out io.Writer, errOut io.Writer) *cobra.Command {
	options := &PKIOptions{
		CommonOptions: common.CommonOptions{
			Out: out,
			Err: errOut,
		},
	}

	cmd := &cobra.Command{
		Use:   "pki",
		Short: "Manage a Ziti PKI",
		Long:  pkiLong,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			cmdhelper.CheckErr(err)
		},
	}

	cmd.AddCommand(NewCmdPKICreate(out, errOut))

	cmd.AddCommand(lets_encrypt.NewCmdLE(out, errOut))

	return cmd
}

// Run implements this command
func (o *PKIOptions) Run() error {
	return o.Cmd.Help()

}
