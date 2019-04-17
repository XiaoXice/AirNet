package config

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSON(t *testing.T){
	type User struct{
		Name []byte `json:"name"`
		Age int     `json:"ageeeeeeeeeeeeeeeeee"`
	}
	user := User{
		Name:[]byte("tom"),
		Age:3,
	}
	b, _ := json.Marshal(user)
	fmt.Println(string(b))
	user = User{}
	err := json.Unmarshal(b, &user)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s\n", user.Name)
}

func TestRandReader(t *testing.T) {
	hash := make([]byte, 128)
	md := sha256.New()
	_,_ = rand.Reader.Read(hash)
	md.Write(hash)
	fmt.Printf("%s\n", hex.EncodeToString(md.Sum(nil)))
}