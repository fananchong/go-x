package gotcp

import (
	"bytes"
	"compress/zlib"
	"io"

	"github.com/golang/protobuf/proto"
)

const (
	DefaultMaxCompressSize = 1024
	DefaultMaxCmdSize      = 2
	CmdSizeLimit           = 8
)

func EncodeCmd(cmd uint64, msg proto.Message) ([]byte, byte, error) {
	return EncodeCmdEx(cmd, msg, DefaultMaxCompressSize, DefaultMaxCmdSize)
}

func EncodeCmdEx(cmd uint64, msg proto.Message, maxCompressSize, maxCmdSize int) ([]byte, byte, error) {
	data, err := proto.Marshal(msg)
	if err != nil {
		xlog.Errorln("[协议] 编码错误 ", err)
		return nil, 0, err
	}
	datalen := len(data)
	flag := 0

	if datalen >= maxCompressSize {
		mbuff, err := zlibCompress(data)
		if mbuff != nil {
			mbufflen := len(mbuff)
			if mbufflen < datalen {
				data = mbuff
				datalen = mbufflen
				flag = 1
			} else {
				xlog.Errorln("[协议] zlib压缩，大小更大! cmd = ", cmd)
			}
		}
		if err != nil {
			xlog.Errorln("[协议] zlib压缩，error = ", err)
		}
	}

	p := make([]byte, maxCmdSize+datalen)
	for i := 0; i < maxCmdSize && i < CmdSizeLimit; i++ {
		p[i] = byte(cmd >> uint(8*i))
	}
	copy(p[maxCmdSize:], data)
	return p, byte(flag), nil
}

func GetCmd(buf []byte) uint64 {
	return GetCmdEx(buf, DefaultMaxCmdSize)
}

func GetCmdEx(buf []byte, maxCmdSize int) uint64 {
	if len(buf) < maxCmdSize || len(buf) == 0 {
		return 0
	}
	var v uint64 = 0
	for i := 0; i < maxCmdSize && i < CmdSizeLimit; i++ {
		v = v | uint64(buf[i])<<uint(8*i)
	}
	return v
}

func DecodeCmd(buf []byte, flag byte, msg proto.Message) proto.Message {
	return DecodeCmdEx(buf, flag, msg, DefaultMaxCmdSize)
}

func DecodeCmdEx(buf []byte, flag byte, msg proto.Message, maxCmdSize int) proto.Message {
	if len(buf) < maxCmdSize || len(buf) == 0 {
		xlog.Errorln("[协议] 数据错误 ", buf)
		return nil
	}
	var mbuff []byte
	if flag == 1 {
		mbuff, err := zlibUnCompress(buf[maxCmdSize:])
		if mbuff == nil {
			xlog.Errorln("[协议] 解压错误 ", err)
			return nil
		}
	} else {
		mbuff = buf[maxCmdSize:]
	}
	err := proto.Unmarshal(mbuff, msg)
	if err != nil {
		xlog.Errorln("[协议] 解码错误 ", err)
		return nil
	}
	return msg
}

func zlibCompress(src []byte) ([]byte, error) {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	_, err := w.Write(src)
	if err != nil {
		w.Close()
		return nil, err
	}
	w.Close()
	return in.Bytes(), nil
}

func zlibUnCompress(src []byte) ([]byte, error) {
	b := bytes.NewReader(src)
	var out bytes.Buffer
	r, err := zlib.NewReader(b)
	if err != nil {
		r.Close()
		return nil, err
	}
	_, err = io.Copy(&out, r)
	if err != nil {
		r.Close()
		return nil, err
	}
	r.Close()
	return out.Bytes(), nil
}
