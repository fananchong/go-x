package service

import (
	"testing"
)

func Test_IpConfig(t *testing.T) {
	err := LoadIpConfig("../assets/ip.toml")
	if err == nil {
		t.Log(gIP)
	} else {
		t.Error(err)
	}
}

