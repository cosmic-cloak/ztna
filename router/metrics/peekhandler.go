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

package metrics

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/channel/v3"
	"github.com/cosmic-cloak/ztna/router/env"
	"github.com/cosmic-cloak/ztna/router/xgress"
	"github.com/openziti/metrics"
	"time"
)

// NewChannelPeekHandler creates a channel PeekHandler which tracks latency, message rate and message size distribution
func NewChannelPeekHandler(linkId string, registry metrics.UsageRegistry) channel.PeekHandler {
	appTxBytesMeter := registry.Meter("fabric.tx.bytesrate")
	appTxMsgMeter := registry.Meter("fabric.tx.msgrate")
	appTxMsgSizeHistogram := registry.Histogram("fabric.tx.msgsize")

	appRxBytesMeter := registry.Meter("fabric.rx.bytesrate")
	appRxMsgMeter := registry.Meter("fabric.rx.msgrate")
	appRxMsgSizeHistogram := registry.Histogram("fabric.rx.msgsize")

	linkTxBytesMeter := registry.Meter("link." + linkId + ".tx.bytesrate")
	linkTxMsgMeter := registry.Meter("link." + linkId + ".tx.msgrate")
	linkTxMsgSizeHistogram := registry.Histogram("link." + linkId + ".tx.msgsize")
	linkRxBytesMeter := registry.Meter("link." + linkId + ".rx.bytesrate")
	linkRxMsgMeter := registry.Meter("link." + linkId + ".rx.msgrate")
	linkRxMsgSizeHistogram := registry.Histogram("link." + linkId + ".rx.msgsize")

	usageCounter := registry.UsageCounter("fabricUsage", env.IntervalSize)

	return &channelPeekHandler{
		appTxBytesMeter:        appTxBytesMeter,
		appTxMsgMeter:          appTxMsgMeter,
		appTxMsgSizeHistogram:  appTxMsgSizeHistogram,
		appRxBytesMeter:        appRxBytesMeter,
		appRxMsgMeter:          appRxMsgMeter,
		appRxMsgSizeHistogram:  appRxMsgSizeHistogram,
		linkTxBytesMeter:       linkTxBytesMeter,
		linkTxMsgMeter:         linkTxMsgMeter,
		linkTxMsgSizeHistogram: linkTxMsgSizeHistogram,
		linkRxBytesMeter:       linkRxBytesMeter,
		linkRxMsgMeter:         linkRxMsgMeter,
		linkRxMsgSizeHistogram: linkRxMsgSizeHistogram,
		usageCounter:           usageCounter,
	}
}

type channelPeekHandler struct {
	appTxBytesMeter metrics.Meter
	appTxMsgMeter   metrics.Meter
	appRxBytesMeter metrics.Meter
	appRxMsgMeter   metrics.Meter

	appTxMsgSizeHistogram metrics.Histogram
	appRxMsgSizeHistogram metrics.Histogram

	linkTxBytesMeter       metrics.Meter
	linkTxMsgMeter         metrics.Meter
	linkRxBytesMeter       metrics.Meter
	linkRxMsgMeter         metrics.Meter
	linkTxMsgSizeHistogram metrics.Histogram
	linkRxMsgSizeHistogram metrics.Histogram

	usageCounter metrics.UsageCounter
}

func (h *channelPeekHandler) Connect(channel.Channel, string) {
}

func (h *channelPeekHandler) Rx(msg *channel.Message, _ channel.Channel) {
	msgSize := int64(len(msg.Body))
	h.linkRxBytesMeter.Mark(msgSize)
	h.linkRxMsgMeter.Mark(1)
	h.linkRxMsgSizeHistogram.Update(msgSize)
	h.appRxBytesMeter.Mark(msgSize)
	h.appRxMsgMeter.Mark(1)
	h.appRxMsgSizeHistogram.Update(msgSize)

	if msg.ContentType == int32(xgress.ContentTypePayloadType) {
		if payload, err := xgress.UnmarshallPayload(msg); err != nil {
			pfxlog.Logger().WithError(err).Error("Failed to unmarshal payload")
		} else {
			h.usageCounter.Update(circuitUsageSource(payload.CircuitId), "fabric.rx", time.Now(), uint64(len(payload.Data)))
		}
	}
}

func (h *channelPeekHandler) Tx(msg *channel.Message, _ channel.Channel) {
	msgSize := int64(len(msg.Body))
	h.linkTxBytesMeter.Mark(msgSize)
	h.linkTxMsgMeter.Mark(1)
	h.linkTxMsgSizeHistogram.Update(msgSize)
	h.appTxBytesMeter.Mark(msgSize)
	h.appTxMsgMeter.Mark(1)
	h.appTxMsgSizeHistogram.Update(msgSize)

	if msg.ContentType == int32(xgress.ContentTypePayloadType) {
		if payload, err := xgress.UnmarshallPayload(msg); err != nil {
			pfxlog.Logger().WithError(err).Error("Failed to unmarshal payload")
		} else {
			h.usageCounter.Update(circuitUsageSource(payload.CircuitId), "fabric.tx", time.Now(), uint64(len(payload.Data)))
		}
	}
}

func (h *channelPeekHandler) Close(channel.Channel) {
	// app level metrics and usageCounter are shared across all links, so we don't dispose of them
	h.linkTxBytesMeter.Dispose()
	h.linkTxMsgMeter.Dispose()
	h.linkTxMsgSizeHistogram.Dispose()
	h.linkRxBytesMeter.Dispose()
	h.linkRxMsgMeter.Dispose()
	h.linkRxMsgSizeHistogram.Dispose()
}

// NewXgressPeekHandler creates an xgress PeekHandler which tracks message rates and histograms as well as usage
func NewXgressPeekHandler(registry metrics.UsageRegistry) xgress.PeekHandler {
	ingressTxBytesMeter := registry.Meter("ingress.tx.bytesrate")
	ingressTxMsgMeter := registry.Meter("ingress.tx.msgrate")
	ingressRxBytesMeter := registry.Meter("ingress.rx.bytesrate")
	ingressRxMsgMeter := registry.Meter("ingress.rx.msgrate")
	egressTxBytesMeter := registry.Meter("egress.tx.bytesrate")
	egressTxMsgMeter := registry.Meter("egress.tx.msgrate")
	egressRxBytesMeter := registry.Meter("egress.rx.bytesrate")
	egressRxMsgMeter := registry.Meter("egress.rx.msgrate")

	ingressTxMsgSizeHistogram := registry.Histogram("ingress.tx.msgsize")
	ingressRxMsgSizeHistogram := registry.Histogram("ingress.rx.msgsize")
	egressTxMsgSizeHistogram := registry.Histogram("egress.tx.msgsize")
	egressRxMsgSizeHistogram := registry.Histogram("egress.rx.msgsize")

	return &xgressPeekHandler{
		ingressTxBytesMeter: ingressTxBytesMeter,
		ingressTxMsgMeter:   ingressTxMsgMeter,
		ingressRxBytesMeter: ingressRxBytesMeter,
		ingressRxMsgMeter:   ingressRxMsgMeter,
		egressTxBytesMeter:  egressTxBytesMeter,
		egressTxMsgMeter:    egressTxMsgMeter,
		egressRxBytesMeter:  egressRxBytesMeter,
		egressRxMsgMeter:    egressRxMsgMeter,

		ingressTxMsgSizeHistogram: ingressTxMsgSizeHistogram,
		ingressRxMsgSizeHistogram: ingressRxMsgSizeHistogram,
		egressTxMsgSizeHistogram:  egressTxMsgSizeHistogram,
		egressRxMsgSizeHistogram:  egressRxMsgSizeHistogram,

		usageCounter: registry.UsageCounter("usage", env.IntervalSize),
	}
}

type circuitUsageSource string

func (c circuitUsageSource) GetIntervalId() string {
	return string(c)
}

func (c circuitUsageSource) GetTags() map[string]string {
	return nil
}

type xgressPeekHandler struct {
	ingressTxBytesMeter metrics.Meter
	ingressTxMsgMeter   metrics.Meter
	ingressRxBytesMeter metrics.Meter
	ingressRxMsgMeter   metrics.Meter
	egressTxBytesMeter  metrics.Meter
	egressTxMsgMeter    metrics.Meter
	egressRxBytesMeter  metrics.Meter
	egressRxMsgMeter    metrics.Meter

	ingressTxMsgSizeHistogram metrics.Histogram
	ingressRxMsgSizeHistogram metrics.Histogram
	egressTxMsgSizeHistogram  metrics.Histogram
	egressRxMsgSizeHistogram  metrics.Histogram

	usageCounter metrics.UsageCounter
}

func (handler *xgressPeekHandler) Rx(x *xgress.Xgress, payload *xgress.Payload) {
	msgSize := int64(len(payload.Data))
	if x.Originator() == xgress.Initiator {
		handler.usageCounter.Update(x, "ingress.rx", time.Now(), uint64(msgSize))
		handler.ingressRxMsgMeter.Mark(1)
		handler.ingressRxBytesMeter.Mark(msgSize)
		handler.ingressRxMsgSizeHistogram.Update(msgSize)
	} else {
		handler.usageCounter.Update(x, "egress.rx", time.Now(), uint64(msgSize))
		handler.egressRxMsgMeter.Mark(1)
		handler.egressRxBytesMeter.Mark(msgSize)
		handler.egressRxMsgSizeHistogram.Update(msgSize)
	}
}

func (handler *xgressPeekHandler) Tx(x *xgress.Xgress, payload *xgress.Payload) {
	msgSize := int64(len(payload.Data))
	if x.Originator() == xgress.Initiator {
		handler.usageCounter.Update(x, "ingress.tx", time.Now(), uint64(msgSize))

		handler.ingressTxMsgMeter.Mark(1)
		handler.ingressTxBytesMeter.Mark(msgSize)
		handler.ingressTxMsgSizeHistogram.Update(msgSize)
	} else {
		handler.usageCounter.Update(x, "egress.tx", time.Now(), uint64(msgSize))
		handler.egressTxMsgMeter.Mark(1)
		handler.egressTxBytesMeter.Mark(msgSize)
		handler.egressTxMsgSizeHistogram.Update(msgSize)
	}
}

func (handler *xgressPeekHandler) Close(*xgress.Xgress) {
}
