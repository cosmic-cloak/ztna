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

package env

import (
	"github.com/openziti/channel/v3"
	"ztna-core/ztna/common/config"
)

// An Xrctrl allows adding handlers to the router <-> controller connection
// on the router side. This means you can support additional message
// types/flows to extend the basic fabric functionality.
//
// There is a corresponding Xctrl interface for extending communication on
// the controller side
type Xrctrl interface {
	config.Subconfig
	channel.BindHandler
	Enabled() bool
	Run(env RouterEnv) error
	NotifyOfReconnect(ch channel.Channel)
}
