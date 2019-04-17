package frame

import (
	"errors"
	"github.com/XiaoXice/AirNet/common/logger"
	"github.com/XiaoXice/AirNet/common/proto"
	n "github.com/XiaoXice/AirNet/net"
	"net"
)

type CallBackChannelFun func(pkg proto.Pkg)

type UdpCon struct {
	Local  *n.Node
	Logger logger.Logger
	c      *net.UDPConn
	R      RouterTable
	List
	socket map[n.HashInfo]map[int32]CallBackChannelFun
}

func (c *UdpCon) Start(ip net.IP) error {
	if c.Local == nil {
		return errors.New("LocalNode info can't be nil")
	}
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   ip,
		Port: int(c.Local.Port),
	})
	if err != nil {
		return err
	}
	c.c = con
	var next []*Next
	c.R = RouterTable{next}
	go udpListener(c)
}

func udpListener(c *UdpCon) {

}
