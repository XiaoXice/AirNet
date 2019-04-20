package database

import (
	"crypto/ecdsa"
	"crypto/x509"
	"github.com/XiaoXice/AirNet/common/proto"
	"net"
	"time"
)

type HashInfo [32]byte

func NewHashInfo(bs []byte) (h HashInfo) {
	h = HashInfo{}
	for index, b := range bs {
		if index>>5 > 0 {
			return h
		}
		h[index] = b
	}
	return h
}
func (h *HashInfo) ToBytes() []byte {
	b := make([]byte, 32)
	for index := range b {
		b[index] = h[index]
	}
	return b
}

type Node struct {
	HashInfo      HashInfo
	Address       net.IPAddr
	Port          int32
	NeedHole      bool
	CanConnection bool
	Description   string
	Type          proto.NodeType
	PublicKey     *ecdsa.PublicKey
	LastCheckTime time.Time
}

func (n *Node) ToNodeListNew() (*NodeListNew, error) {
	var puk []byte
	if n.PublicKey == nil {
		puk = nil
	} else {
		var err error
		puk, err = x509.MarshalPKIXPublicKey(n.PublicKey)
		if err != nil {
			return nil, err
		}
	}
	return &NodeListNew{
		HashInfo:      n.HashInfo.ToBytes(),
		Address:       n.Address.IP,
		Port:          int(n.Port),
		NeedHole:      n.NeedHole,
		CanConnection: n.CanConnection,
		Description:   n.Description,
		Type:          int(n.Type),
		PublicKey:     puk,
		LastCheckTime: n.LastCheckTime,
	}, nil
}
