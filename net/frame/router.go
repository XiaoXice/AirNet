package frame

import (
	"github.com/XiaoXice/AirNet/net/database"
	"time"
)

type RouterTable struct {
	Table []*Next
}

type Next struct {
	Node   *database.Node
	Weight map[database.HashInfo]float32
	timer  map[database.HashInfo][]chan bool
}

func (n *Next) TimerListener(needToAdd database.HashInfo, CB *func()) {
	stopChan := make(chan bool)
	if l, ok := n.timer[needToAdd]; ok {
		l = append(l, stopChan)
	}
	n.timer[needToAdd] = []chan bool{stopChan}
	var NotTimeOut = false
	go func() {
		<-stopChan
		NotTimeOut = true
		return
	}()
	go func() {
		time.Sleep(5 * time.Second)
		if NotTimeOut == false {
			stopChan <- true
			(*CB)()
		}
		return
	}()
}
func (r *RouterTable) Next(remote database.HashInfo) (list []float32) {
	var total float32
	for index, next := range r.Table {
		if weight, ok := next.Weight[remote]; ok {
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

func (r *RouterTable) g(n database.HashInfo) *Next {
	for _, v := range r.Table {
		if n == v.Node.HashInfo {
			return v
		}
	}
	return nil
}

func (r *RouterTable) ToSend(ro database.HashInfo, ne database.HashInfo) {
	if nx := r.g(ne); nx != nil {

	}
}
