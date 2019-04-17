package database

import (
	n "github.com/XiaoXice/AirNet/net"
)

type NodeList struct {
	list map[n.HashInfo]n.Node
}

func (d *DataBase) LoadNodeList() (*NodeList, error) {
	rows, err := d.SelectALlNodeList()
	if err != nil {
		return nil, err
	}
	list := make(map[n.HashInfo]n.Node)
	for _,row := range rows {
		node,err := row.ToNode()
		if err != nil {
			return nil, err
		}
		list[node.HashInfo] = *node
	}
	return &NodeList{list}, nil
}
