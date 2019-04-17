package frame

import (
	n "github.com/XiaoXice/AirNet/net"
	"math/rand"
)

type RouterTable struct {
	Table []*Next
}

type Next struct {
	Node *n.Node
	Weight map[n.HashInfo]float32
	stop chan bool
}

func (r *RouterTable)Next(remote n.HashInfo) (list []float32) {
	var total float32
	for index,next := range r.Table{
		if weight,ok := next.Weight[remote];ok {
			list[index] = weight
		}else {
			next.Weight[remote] = 0.5
			list[index] = 0.5
		}
		total += list[index]
	}
	return list
}

func (r *RouterTable)AddNext(node *n.Node) error {
	N := &Next{
		Node: node,
		Weight: make(map[n.HashInfo]float32),
	}
}
