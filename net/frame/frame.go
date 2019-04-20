package frame

import (
	"errors"
	"github.com/XiaoXice/AirNet/common/logger"
	"github.com/XiaoXice/AirNet/common/proto"
	"github.com/XiaoXice/AirNet/net/database"
	"net"
)

type CallBackFun func(pkg proto.Pkg)

type UdpCon struct {
	Local  *database.Node
	Logger logger.Logger
	c      *net.UDPConn
	R      RouterTable
	List	database.NodeList
}

func (c *UdpCon) Start(ip net.IP) error {
	if c.Local == nil {
		return errors.New("LocalNode info can't be nil")
	}
	stop := make(chan bool)
	go udpListener(c,ip,stop)
	var next []*Next
	c.R = RouterTable{next}
}

func udpListener(c *UdpCon,ip net.IP, stop chan bool) {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   ip,
		Port: int(c.Local.Port),
	})
	if err != nil {
		c.Logger.Handler(err.Error())
		stop <- true
		return
	}
	c.c = con
	buf := make([]byte,10240)
	for  {
		l, address, err := con.ReadFromUDP(buf)
		if err != nil {
			stop <- true
			return
		}
		if
	}
}
