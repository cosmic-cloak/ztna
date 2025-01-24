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

package xgress_proxy

import (
	"errors"
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/cosmic-cloak/ztna/router/env"
	"github.com/cosmic-cloak/ztna/router/xgress"
	"github.com/openziti/identity"
	"github.com/openziti/transport/v2"
)

func newListener(id *identity.TokenId, ctrl env.NetworkControllers, options *xgress.Options, tcfg transport.Configuration, service string) xgress.Listener {
	return &listener{
		id:          id,
		ctrl:        ctrl,
		options:     options,
		tcfg:        tcfg,
		service:     service,
		closeHelper: &xgress.CloseHelper{},
	}
}

func (listener *listener) Listen(address string, bindHandler xgress.BindHandler) error {
	if address == "" {
		return errors.New("address must be specified for proxy listeners")
	}
	txAddress, err := transport.ParseAddress(address)
	if err != nil {
		return fmt.Errorf("cannot listen on invalid address [%s] (%s)", address, err)
	}

	acceptF := func(peer transport.Conn) {
		go listener.handleConnect(peer, bindHandler)
	}
	go listener.closeHelper.Init(txAddress.MustListen("tcp", listener.id, acceptF, listener.tcfg))

	return nil
}

func (listener *listener) handleConnect(peer transport.Conn, bindHandler xgress.BindHandler) {
	conn := &proxyXgressConnection{peer}
	log := pfxlog.ContextLogger(conn.LogContext())
	request := &xgress.Request{ServiceId: listener.service}
	response := xgress.CreateCircuit(listener.ctrl, conn, request, bindHandler, listener.options)
	if !response.Success {
		log.Errorf("error creating circuit (%s)", response.Message)
		_ = peer.Close()
	}
}

type listener struct {
	id          *identity.TokenId
	ctrl        env.NetworkControllers
	options     *xgress.Options
	tcfg        transport.Configuration
	service     string
	closeHelper *xgress.CloseHelper
}

func (listener *listener) Close() error {
	return listener.closeHelper.Close()
}
