package frame

import (
	"bytes"
	"encoding/hex"
	"errors"
	"github.com/XiaoXice/AirNet/common/logger"
	"github.com/XiaoXice/AirNet/common/proto"
	"github.com/XiaoXice/AirNet/common/util"
	"github.com/XiaoXice/AirNet/net/database"
	proto2 "github.com/golang/protobuf/proto"
	"net"
)

type CallBackFun func(pkg proto.Pkg)

type UdpCon struct {
	Local  *database.Node
	Logger logger.Logger
	c      *net.UDPConn
	R      RouterTable
	List   database.NodeList
}

func (c *UdpCon) Start(ip net.IP) error {
	if c.Local == nil {
		return errors.New("LocalNode info can't be nil")
	}
	stop := make(chan bool)
	go udpListener(c, ip, stop)
	var next []*Next
	c.R = RouterTable{next}
	return nil
}

func udpListener(c *UdpCon, ip net.IP, stop chan bool) {
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
	buf := make([]byte, 10240)
	for {
		l, address, err := con.ReadFromUDP(buf)
		if err != nil {
			stop <- true
			return
		}
		go getUdpFrame(c, l, address, buf)
	}
}

func getUdpFrame(c *UdpCon, l int, address *net.UDPAddr, buf []byte) {
	getFrame := &proto.Frame{}
	err := proto2.Unmarshal(buf[:l], getFrame)
	if err != nil {
		frameToSend := &proto.Frame{
			Type: proto.FrameType_RE,
			From: c.Local.HashInfo.ToBytes(),
			TTL:  1,
		}
		answer, _ := frameToSend.Marshal()
		_, err := c.c.WriteToUDP(answer, address)
		if err != nil {
			c.Logger.Handler(err.Error())
		}
	} else {
		if getFrame.TTL -= 1; getFrame.TTL > 0 {
			getPkg := &proto.Pkg{}
			if getFrame.Content != nil {
				err := proto2.Unmarshal(getFrame.Content, getPkg)
				if err == nil {
					// TODO: 发送报文请求重发
				}
			}
			if bytes.Equal(getFrame.To, c.Local.HashInfo.ToBytes()) {
				switch getFrame.Type {
				case proto.FrameType_ACK:
					go func() {
						c.R.ACK(
							database.NewHashInfo(getFrame.PassWay[len(getFrame.PassWay)-1]),
							database.NewHashInfo(getFrame.From))
						getFrame.TTL -= 1
						t := util.Find(util.ToSlice(getFrame.PassWay), c.Local.HashInfo.ToBytes())
						if t > 0 {
							send, _ := getFrame.Marshal()
							for i := 1; t >= i ; i += 1 {
								next := database.NewHashInfo(getFrame.PassWay[t-i])
								if n, ok := c.List.List[next]; ok {
									_, err = c.c.WriteToUDP(send, n.ToUDPAddr())
									if err == nil {break}
								}
								// TODO: 把ACK发给更多的主机
							}
						}
					}()
					break
				}
				// TODO: 发给监听的部分进行处理
			} else {
				l := c.R.Next(database.NewHashInfo(getPkg.To), database.NewHashInfoList(getFrame.PassWay))
				next := c.R.Table[util.RandomInList(l,nil)[0]]
				if getFrame.Type != proto.FrameType_ACK {
					getFrame.PassWay = append(getFrame.PassWay, c.Local.HashInfo.ToBytes())
				}
				send, _ := getFrame.Marshal()
				c.R.ToSend(database.NewHashInfo(getPkg.To),next.Node.HashInfo,func() {
					c.Logger.Handler("[Warn] TimeOut ->" + hex.EncodeToString(getPkg.To))
				})
				_, err = c.c.WriteToUDP(send, next.ToUDPAddr())
				if err != nil {
					// TODO: 直接干掉
				}
			}
		}
	}
}
