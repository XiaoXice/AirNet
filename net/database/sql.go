package database

import (
	"crypto/ecdsa"
	"crypto/x509"
	"github.com/XiaoXice/AirNet/common/proto"
	net2 "net"
	"time"
)

type PortNew struct {
	LocalPort     int
	RemotePort    int
	RemoteAddress string
}

func (d *DataBase) SelectALlPort() (reRows []PortNew, err error) {
	rows, err := d.db.Query("SELECT * FROM NodeList;")
	if err != nil {
		d.Logger.Handler("Err on SelectPort: " + err.Error())
		return nil, err
	}
	//noinspection GoUnhandledErrorResult
	defer rows.Close()
	reRows = make([]PortNew, 0)
	for rows.Next() {
		row := PortNew{}
		if err = rows.Scan(&row.LocalPort, &row.RemotePort, &row.RemoteAddress); err != nil {
			d.Logger.Handler(err.Error())
			return nil, err
		} else {
			reRows = append(reRows, row)
		}
	}
	return reRows, nil
}

type NodeListNew struct {
	HashInfo      []byte
	Address       []byte
	Port          int
	NeedHole      bool
	CanConnection bool
	Description   string
	Type          int
	PublicKey     []byte
	LastCheckTime time.Time
}

func (n *NodeListNew) ToNode() (node *Node, err error) {
	var p interface{}
	if n.PublicKey != nil {
		p, err = x509.ParsePKIXPublicKey(n.PublicKey)
		if err != nil {
			p = nil
			//return nil, err
		}
	} else {
		p = nil
	}
	var pu *ecdsa.PublicKey
	if p == nil {
		pu = nil
	} else {
		pu = p.(*ecdsa.PublicKey)
	}
	node = &Node{
		HashInfo:      NewHashInfo(n.HashInfo),
		Address:       net2.IPAddr{IP: n.Address},
		Port:          int32(n.Port),
		NeedHole:      n.NeedHole,
		CanConnection: n.CanConnection,
		Description:   n.Description,
		Type:          proto.NodeType(n.Type),
		PublicKey:     pu,
		LastCheckTime: n.LastCheckTime,
	}
	return node, nil
}
func (n *NodeListNew) ToList() []interface{} {
	return nil
}
func (d *DataBase) SelectALlNodeList() (reRows []NodeListNew, err error) {
	rows, err := d.db.Query("SELECT * FROM NodeList;")
	if err != nil {
		d.Logger.Handler("Err on SelectPort: " + err.Error())
		return nil, err
	}
	//noinspection GoUnhandledErrorResult
	defer rows.Close()
	reRows = make([]NodeListNew, 0)
	for rows.Next() {
		row := NodeListNew{}
		err = rows.Scan(&row.HashInfo,
			&row.Address,
			&row.Port,
			&row.NeedHole,
			&row.CanConnection,
			&row.Description,
			&row.Type,
			&row.PublicKey,
			&row.LastCheckTime)
		if err != nil {
			d.Logger.Handler(err.Error())
			return nil, err
		} else {
			reRows = append(reRows, row)
		}
	}
	return reRows, nil
}

type RouterNew struct {
	Remote []byte
	Next   []byte
	Weight float32
}

func (d *DataBase) SelectALlRouter() (reRows []RouterNew, err error) {
	rows, err := d.db.Query("SELECT * FROM Router;")
	if err != nil {
		d.Logger.Handler("Err on SelectPort: " + err.Error())
		return nil, err
	}
	//noinspection GoUnhandledErrorResult
	defer rows.Close()
	reRows = make([]RouterNew, 0)
	for rows.Next() {
		row := RouterNew{}
		err = rows.Scan(
			&row.Remote,
			&row.Next,
			&row.Weight)
		if err != nil {
			d.Logger.Handler(err.Error())
			return nil, err
		} else {
			reRows = append(reRows, row)
		}
	}
	return reRows, nil
}
