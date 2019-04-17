package net

import (
	"crypto/ecdsa"
	"github.com/XiaoXice/AirNet/common/proto"
	"net"
	"time"
)
type HashInfo [64]byte

type Node struct {
	HashInfo HashInfo
	Address net.IPAddr
	Port int32
	NeedHole bool
	CanConnection bool
	Description string
	Type proto.NodeType
	PublicKey *ecdsa.PublicKey
	LastCheckTime time.Time
}


