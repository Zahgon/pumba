package iptables

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/container"
)

// `iptables loss` command
type lossCommand struct {
	client      iptablesClient
	gp          *chaos.GlobalParams
	req         *container.IPTablesRequest
	iface       string
	protocol    string
	limit       int
	mode        string
	probability float64
	every       int
	packet      int
}

const (
	ModeRandom = "random"
	ModeNTH    = "nth"
)

// NewLossCommand create new iptables loss command
func NewLossCommand(client iptablesClient,
	gp *chaos.GlobalParams,
	base *RequestBase,
	mode string, // loss mode
	probability float64, // loss probability
	every int, // drop every nth
	packet int, // start budget for every nth
) (chaos.Command, error) {
	_ = "STUB: not implemented"
	// get mode
	return *new(chaos.Command), nil
}

// get loss probability

// get every

// get packet

// Run iptables loss command
func (n *lossCommand) Run(ctx context.Context, random bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *lossCommand) buildIPTablesCmd() (addCmdPrefix, delCmdPrefix, cmdSuffix []string) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// mode == nth
