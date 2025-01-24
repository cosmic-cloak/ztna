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

package agentcli

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/openziti/agent"
	"github.com/cosmic-cloak/ztna/ziti/cmd/common"
	cmdhelper "github.com/cosmic-cloak/ztna/ziti/cmd/helpers"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

type AgentSetChannelLogLevelAction struct {
	AgentOptions
}

func NewSetChannelLogLevelCmd(p common.OptionsProvider) *cobra.Command {
	action := &AgentSetChannelLogLevelAction{
		AgentOptions: AgentOptions{
			CommonOptions: p(),
		},
	}

	cmd := &cobra.Command{
		Use:   "set-channel-log-level channel log-level (panic, fatal, error, warn, info, debug, trace)",
		Short: "Sets a channel-specific log level in the target application",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			action.Cmd = cmd
			action.Args = args
			err := action.Run()
			cmdhelper.CheckErr(err)
		},
	}

	action.AddAgentOptions(cmd)

	return cmd
}

// Run implements the command
func (self *AgentSetChannelLogLevelAction) Run() error {
	if self.Cmd.Flags().Changed("timeout") {
		time.AfterFunc(self.timeout, func() {
			fmt.Println("operation timed out")
			os.Exit(-1)
		})
	}

	channelArg := self.Args[0]
	levelArg := self.Args[1]

	var level logrus.Level
	var found bool
	for _, l := range logrus.AllLevels {
		if strings.EqualFold(l.String(), levelArg) {
			level = l
			found = true
		}
	}

	if !found {
		return errors.Errorf("invalid log level %v", levelArg)
	}

	lenBuf := make([]byte, 8)
	lenLen := binary.PutVarint(lenBuf, int64(len(channelArg)))
	buf := &bytes.Buffer{}
	buf.Write(lenBuf[:lenLen])
	buf.Write([]byte(channelArg))
	buf.WriteByte(byte(level))

	return self.MakeRequest(agent.SetChannelLogLevel, buf.Bytes(), self.CopyToWriter(os.Stdout))
}
