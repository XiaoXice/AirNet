package config

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"github.com/XiaoXice/AirNet/common/proto"
	"github.com/XiaoXice/AirNet/net"
	"io/ioutil"
	"os"
	path2 "path"
)

type Configs struct {
	HashInfo     []byte
	PublicKey    []byte
	PrivateKey   []byte
	Description  string
	Type         int
	DataBaseType string
	Path         string
	UDPPort      int
}

func (c *Configs) ToConfig() (*Config, error) {
	k, e := x509.ParseECPrivateKey(c.PrivateKey)
	if e != nil {
		return nil, e
	}
	config := Config{
		HashInfo:     net.NewHashInfo(c.HashInfo),
		PublicKey:    &k.PublicKey,
		PrivateKey:   k,
		Description:  c.Description,
		Type:         proto.NodeType(c.Type),
		DataBaseType: c.DataBaseType,
		Path:         c.Path,
		UDPPort:      int32(c.UDPPort),
	}
	return &config, nil
}

type Config struct {
	HashInfo     net.HashInfo
	PublicKey    *ecdsa.PublicKey
	PrivateKey   *ecdsa.PrivateKey
	Description  string
	Type         proto.NodeType
	DataBaseType string
	Path         string
	UDPPort      int32
}

func LoadJsonConfig(path string) (c *Config, err error) {
	if contents, err := ioutil.ReadFile(path); err == nil {
		_, e := os.Stat(path)
		if e != nil {
			return initConfig(path2.Join(path2.Dir(path), "data.db"))
		}
		cj := Configs{}
		err := json.Unmarshal(contents, &cj)
		if err != nil {
			return nil, err
		}
		return cj.ToConfig()
	}
	return nil, err
}

func (c *Config) toJson() ([]byte, error) {
	pr, err := x509.MarshalECPrivateKey(c.PrivateKey)
	if err != nil {
		return nil, err
	}
	pu, err := x509.MarshalPKIXPublicKey(c.PublicKey)
	if err != nil {
		return nil, err
	}
	cj := Configs{
		HashInfo:     c.HashInfo.ToBytes(),
		PublicKey:    pu,
		PrivateKey:   pr,
		Description:  c.Description,
		Type:         int(c.Type),
		DataBaseType: c.DataBaseType,
		Path:         c.Path,
		UDPPort:      int(c.UDPPort),
	}
	b, err := json.Marshal(cj)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (c *Config) SaveJsonConfig(path string) error {
	b, err := c.toJson()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, b, 0600)
	if err != nil {
		return err
	}
	return nil
}

func initConfig(dbPath string) (*Config, error) {
	hash := make([]byte, 1024)
	sha := sha256.New()
	_, _ = rand.Reader.Read(hash)
	sha.Write(hash)
	hashInfo := net.NewHashInfo(sha.Sum(nil))
	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		return nil, err
	}
	config := &Config{
		HashInfo:     hashInfo,
		PublicKey:    &privateKey.PublicKey,
		PrivateKey:   privateKey,
		Description:  "我只是一个萌新节点, 请多关照.",
		Type:         proto.NodeType_Common,
		DataBaseType: "sqlite3",
		Path:         dbPath,
		UDPPort:      0,
	}
	return config, nil
}
