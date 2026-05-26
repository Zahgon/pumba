package docker

import (
	"context"

	ctr "github.com/alexei-led/pumba/pkg/container"
)

// NetemContainer injects sidecar netem container into the given container network namespace
func (client dockerClient) NetemContainer(ctx context.Context, req *ctr.NetemRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// StopNetemContainer stops the netem container injected into the given container network namespace
func (client dockerClient) StopNetemContainer(ctx context.Context, req *ctr.NetemRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (client dockerClient) startNetemContainer(ctx context.Context, req *ctr.NetemRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// use dockerclient ExecStart to run Traffic Control:
// 'tc qdisc add dev eth0 root netem delay 100ms'
// http://www.linuxfoundation.org/collaborate/workgroups/networking/netem

// stop disruption command
// netemStopCommand := "tc qdisc del dev eth0 root netem"

func (client dockerClient) stopNetemContainer(ctx context.Context, req *ctr.NetemRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// delete qdisc 'parent 1:1 handle 10:'
// http://www.linuxfoundation.org/collaborate/workgroups/networking/netem

// delete qdisc 'parent 1:2 handle 20:'
// http://www.linuxfoundation.org/collaborate/workgroups/networking/netem

// delete qdisc 'parent 1:3 handle 30:'
// http://www.linuxfoundation.org/collaborate/workgroups/networking/netem

// delete qdisc 'root handle 1: prio'
// http://www.linuxfoundation.org/collaborate/workgroups/networking/netem

// stop netem command
// http://www.linuxfoundation.org/collaborate/workgroups/networking/netem

func (client dockerClient) startNetemContainerIPFilter(ctx context.Context, req *ctr.NetemRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// use dockerclient ExecStart to run Traffic Control
// to filter network, needs to create a priority scheduling, add a low priority
// queue, apply netem command on that queue only, then route IP traffic to the low priority queue
// See more: http://www.linuxfoundation.org/collaborate/workgroups/networking/netem

//            1:   root qdisc
//           / | \
//          /  |  \
//         /   |   \
//       1:1  1:2  1:3    classes
//        |    |    |
//       10:  20:  30:    qdiscs
//      sfq  sfq  netem
// band  0    1     2

// Create a priority-based queue. This *instantly* creates classes 1:1, 1:2, 1:3
// 'tc qdisc add dev <netInterface> root handle 1: prio'
// See more: http://man7.org/linux/man-pages/man8/tc-netem.8.html

// Create Stochastic Fairness Queueing (sfq) queueing discipline for 1:1 class.
// 'tc qdisc add dev <netInterface> parent 1:1 handle 10: sfq'
// See more: https://linux.die.net/man/8/tc-sfq

// Create Stochastic Fairness Queueing (sfq) queueing discipline for 1:2 class
// 'tc qdisc add dev <netInterface> parent 1:2 handle 20: sfq'
// See more: https://linux.die.net/man/8/tc-sfq

// Add queueing discipline for 1:3 class. No traffic is going through 1:3 yet
// 'tc qdisc add dev <netInterface> parent 1:3 handle 30: netem <netemCmd>'
// See more: http://man7.org/linux/man-pages/man8/tc-netem.8.html

// # redirect traffic to specific IP through band 3
// 'tc filter add dev <netInterface> protocol ip parent 1:0 prio 1 u32 match ip dst <targetIP> flowid 1:3'
// See more: http://man7.org/linux/man-pages/man8/tc-netem.8.html

// # redirect traffic to specific sport through band 3
// 'tc filter add dev <netInterface> protocol ip parent 1:0 prio 1 u32 match ip <s/d>port <targetPort> 0xffff flowid 1:3'
// See more: http://man7.org/linux/man-pages/man8/tc-netem.8.html

// # redirect traffic to specific dport through band 3
// 'tc filter add dev <netInterface> protocol ip parent 1:0 prio 1 u32 match ip <s/d>port <targetPort> 0xffff flowid 1:3'
// See more: http://man7.org/linux/man-pages/man8/tc-netem.8.html

func (client dockerClient) tcCommands(ctx context.Context, c *ctr.Container, argsList [][]string, tcimg string, pull bool) error {
	_ = "STUB: not implemented"
	return nil
}
