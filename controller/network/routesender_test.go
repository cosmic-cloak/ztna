package network

import (
	"github.com/cosmic-cloak/ztna/controller/model"
	"testing"

	"github.com/michaelquigley/pfxlog"
	"github.com/cosmic-cloak/ztna/common/ctrl_msg"
	"github.com/cosmic-cloak/ztna/controller/xt"
	"github.com/cosmic-cloak/ztna/controller/xt_smartrouting"
)

func TestRouteSender_DestroysTerminatorWhenInvalidOnHandleRouteSendAndWeControl(t *testing.T) {
	ctx := model.NewTestContext(t)
	defer ctx.Cleanup()

	config := newTestConfig(ctx)
	defer close(config.closeNotify)

	network, err := NewNetwork(config, ctx)
	ctx.NoError(err)

	entityHelper := newTestEntityHelper(ctx, network)
	logger := pfxlog.ChannelLogger("test")

	router1 := entityHelper.addTestRouter()
	router2 := entityHelper.addTestRouter()
	path := &model.Path{
		Nodes: []*model.Router{router1, router2},
	}

	svc := entityHelper.addTestService("svc")

	instanceId := "instanceId"

	term := entityHelper.addTestTerminator(svc.Id, router1.Id, instanceId, true)
	term.Binding = "edge"

	errCode := byte(ctrl_msg.ErrorTypeInvalidTerminator)

	rs := routeSender{
		serviceCounters: network,
		terminators:     network.Terminator,
		attendance:      make(map[string]bool),
	}

	status := &RouteStatus{
		Router:    router1,
		ErrorCode: &errCode,
		Success:   false,
		Attempt:   1,
		Err:       "THIS IS A TEST",
	}

	peerData, cleanup, err := rs.handleRouteSend(1, path, xt_smartrouting.NewFactory().NewStrategy(), status, term, logger)
	ctx.Error(err)
	ctx.ErrorContains(err, status.Err)
	ctx.Nil(peerData)
	ctx.Empty(cleanup)

	newTerm, err := network.Terminator.Read(term.Id)
	ctx.Error(err)
	ctx.Nil(newTerm)
}

func TestRouteSender_SetPrecidenceToNilTerminatorWhenInvalidOnHandleRouteSendAndWeDontControl(t *testing.T) {
	ctx := model.NewTestContext(t)
	defer ctx.Cleanup()

	config := newTestConfig(ctx)
	defer close(config.closeNotify)

	network, err := NewNetwork(config, ctx)
	ctx.NoError(err)

	entityHelper := newTestEntityHelper(ctx, network)
	logger := pfxlog.ChannelLogger("test")

	router1 := entityHelper.addTestRouter()
	router2 := entityHelper.addTestRouter()
	path := &model.Path{
		Nodes: []*model.Router{router1, router2},
	}

	svc := entityHelper.addTestService("svc")

	identity := "identity"

	term := entityHelper.addTestTerminator(svc.Id, router1.Id, identity, true)
	term.Binding = "DNE"

	errCode := byte(ctrl_msg.ErrorTypeInvalidTerminator)

	rs := routeSender{
		serviceCounters: network,
		terminators:     network.Terminator,
		attendance:      make(map[string]bool),
	}

	status := &RouteStatus{
		Router:    router1,
		ErrorCode: &errCode,
		Success:   false,
		Attempt:   1,
		Err:       "THIS IS A TEST",
	}

	peerData, cleanup, err := rs.handleRouteSend(1, path, xt_smartrouting.NewFactory().NewStrategy(), status, term, logger)
	ctx.Error(err)
	ctx.ErrorContains(err, status.Err)
	ctx.Nil(peerData)
	ctx.Empty(cleanup)

	newTerm, err := network.Terminator.Read(term.Id)
	ctx.NoError(err)
	ctx.NotNil(newTerm)

	ctx.Equal(xt.Precedences.Failed, newTerm.GetPrecedence())
}
