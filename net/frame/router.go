package frame

import (
	"github.com/XiaoXice/AirNet/common/util"
	"github.com/XiaoXice/AirNet/net/database"
	"net"
	"time"
)

type RouterTable struct {
	Table []*Next
}

type Next struct {
	Node   *database.Node
	Weight map[database.HashInfo]float32
	timer  map[database.HashInfo][]*chan bool
}

func (n *Next) TimerListener(needToAdd database.HashInfo, CB func()) {
	stopChan := make(chan bool)
	if l, ok := n.timer[needToAdd]; ok {
		l = append(l, &stopChan)
	}
	n.timer[needToAdd] = []*chan bool{&stopChan}
	var NotTimeOut = false
	go func() {
		<-stopChan
		NotTimeOut = true
		t := util.Find(util.ToSlice(n.timer[needToAdd]), stopChan)
		if t > 0 && t < len(n.timer[needToAdd])- 1 {
			n.timer[needToAdd] = append(n.timer[needToAdd][:t], n.timer[needToAdd][t+1:]...)
		} else if t == 0 {
			n.timer[needToAdd] = append(n.timer[needToAdd][1:])
		} else if t == len(n.timer[needToAdd]) - 1 {
			n.timer[needToAdd] = append(n.timer[needToAdd][:t])
		}
		return
	}()
	go func() {
		time.Sleep(5 * time.Second)
		if NotTimeOut == false {
			stopChan <- true
			CB()
		}
		return
	}()
}

func (n *Next)ToUDPAddr() *net.UDPAddr {
	return n.Node.ToUDPAddr()
}

func (r *RouterTable) Next(remote database.HashInfo, passWay []database.HashInfo) (list []float32) {
	var total float32
	l := util.ToSlice(passWay)
	for index, next := range r.Table {
		if util.Find(l, next.Node.HashInfo) > -1 {
			list[index] = 0
		} else if weight, ok := next.Weight[remote]; ok {
			list[index] = weight
		} else {
			next.Weight[remote] = 0.5
			list[index] = 0.5
		}
		total += list[index]
	}
	return list
}

func (r *RouterTable) AddNext(node *database.Node) error {
	N := &Next{
		Node:   node,
		Weight: make(map[database.HashInfo]float32),
	}
	r.Table = append(r.Table, N)
	return nil
}

func (r *RouterTable) G(n database.HashInfo) *Next {
	for _, v := range r.Table {
		if n == v.Node.HashInfo {
			return v
		}
	}
	return nil
}

func (r *RouterTable) ToSend(ro database.HashInfo, ne database.HashInfo, timeOut func()) *database.Node {
	if nx := r.G(ne); nx != nil {
		nx.TimerListener(ro, timeOut)
		return nx.Node
	}
	return nil
}

func (r *RouterTable) ACK(ro database.HashInfo, ne database.HashInfo) {

}