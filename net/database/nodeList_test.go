package database

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"net"
	"testing"
	"time"
)

func TestDataBase_LoadAndSaveNodeList(t *testing.T) {
	if !DB.CanUse() {
		TestDataBase_Connect(t)
	}
	nl, err := DB.LoadNodeList()
	if err != nil {
		t.Error(err)
	}
	r := randomNode()
	nl.List[r.HashInfo] = r
	err = DB.SaveNodeList(nl)
	if err != nil {
		t.Error(err)
	}
}

func randomNode() Node {
	hash := make([]byte, 1024)
	sha := sha256.New()
	_, _ = rand.Reader.Read(hash)
	sha.Write(hash)
	hashInfo := NewHashInfo(sha.Sum(nil))
	privateKey, _ := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	return Node{
		HashInfo:      hashInfo,
		Address:       net.IPAddr{IP: net.IPv4(0, 0, 0, 0)},
		Port:          0,
		NeedHole:      false,
		CanConnection: false,
		Description:   "",
		Type:          0,
		PublicKey:     &privateKey.PublicKey,
		LastCheckTime: time.Now(),
	}
}
