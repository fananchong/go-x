package service

import "github.com/fananchong/multiconfig"

type IpConfig struct {
	Data [][2]string
}

var gIP map[string]string = make(map[string]string)

func LoadIpConfig(path string) error {
	cfg := &IpConfig{}
	m := &multiconfig.TOMLLoader{
		Path: path,
	}
	err := m.Load(cfg)
	if err != nil {
		return err
	}
	for _, v := range cfg.Data {
		gIP[v[0]] = v[1]
	}
	return nil
}

func GetIpList() *map[string]string {
	return &gIP
}

