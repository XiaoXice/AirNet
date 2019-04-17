package database

import (
	n "github.com/XiaoXice/AirNet/net"
)

type NodeList struct {
	list map[n.HashInfo] n.Node
}

func (d *DataBase) LoadNodeList() *NodeList {

}