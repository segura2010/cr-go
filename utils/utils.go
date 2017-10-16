package utils

import (
	"encoding/binary"
	"bytes"
	"io"
)

func Int24ToInt32(bs []byte) uint32 {
    return uint32(bs[2]) | uint32(bs[1])<<8 | uint32(bs[0])<<16
}

func GetRrsInt32(value int32)([]byte){
	buf := new(bytes.Buffer)

	if value == 0{
		binary.Write(buf, binary.BigEndian, byte(0))
	}

	var b int32
	rotate := true
	value = (value << 1) ^ (value >> 31)
	value = value >> 0
	for value > 0{
		b = (value & 0x7f)
    	if value >= 0x80{
       		b |= 0x80;
    	}
    	if rotate {
        	rotate = false
        	lsb := b & 0x1
        	msb := (b & 0x80) >> 7
        	b = b >> 1 // rotate to the right
        	b = b & -(1^0xC0) // clear 7th and 6th bit
        	b = b | (msb << 7) | (lsb << 6) // insert msb and lsb back in
    	}

        binary.Write(buf, binary.BigEndian, byte(b))
        value = value >> 7;
    }

    return buf.Bytes()
}

func WriteBytes(w io.Writer, order binary.ByteOrder, data []byte){
	var fieldLen int32
	fieldLen = int32(len(data))
	binary.Write(w, order, fieldLen)
	binary.Write(w, order, data)
}

func StringIndexOf(array string, e rune) (int){
	for i, ele := range array{
		if ele == e{
			return i
		}
	}
	return -1
}

var IDCHARS = "0289PYLQGRJCUV"

func Tag2HiLo(tag string) [2]int32{
	charLen := len(IDCHARS)
	id := 0
	for _,c := range tag{
		charIdx := StringIndexOf(IDCHARS, c)
		id *= charLen
		id += charIdx
	}

	hi := id % 256
    lo := (id - hi) >> 8

    return [2]int32{int32(hi), int32(lo)}
}

func HiLo2Tag(hi, lo int32) string{
	charLen := len(IDCHARS)
	id := (int64(lo) << 8) + int64(hi)
	tag := ""
	for id > 0{
		remainder := int64(id % int64(charLen))
		tag = string(IDCHARS[remainder]) + tag
		id = id - remainder
		id = id / int64(charLen)
	}

    return tag
}
