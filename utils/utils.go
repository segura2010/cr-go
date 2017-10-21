package utils

import (
	"encoding/binary"
	"bytes"
	"io"
)

// it converts int24 to int32
func Int24ToInt32(bs []byte) uint32 {
    return uint32(bs[2]) | uint32(bs[1])<<8 | uint32(bs[0])<<16
}

// it returns a RRSINT32 (supercell's special int32) from an int32
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
        value = int32(uint32(value) >> 7)
    }

    return buf.Bytes()
}

// it reads a RRSINT32 (supercell's special int32) from a packet and returns it as int32
func ReadRrsInt32(w io.Reader, order binary.ByteOrder, data *int32){
	var c int32 = 0
	var value int32 = int32(uint32(0) >> 0) // uint32(x)>>y = x>>>y
	var seventh byte
	var b byte
	var bb int32
	var msb byte

	for{

		binary.Read(w, order, &b)
		bb = int32(b)

		if c == 0{
			seventh = (b & 0x40) >> 6 // save 7th bit
			msb = (b & 0x80) >> 7 // save msb
			bb = int32(int32(b) << 1) // rotate to the left
			bb = int32(bb & int32((-1)^(0x181))) // clear 8th and 1st bit and 9th if any
			bb = bb | int32(msb << 7) | int32(seventh) // insert msb and 6th back in
		}

		value |= int32(uint32(bb & 0x7f) << uint32(7 * c))

		c += 1

		if ((b & 0x80) == 0){
			break;
		}
    }

    value = ( int32(uint32(value) >> uint32(1)) ^ -(value & 1) ) | 0

    *data = int32(value)
}

// it writes a byte array in a buffer (first it writes the length of the []byte and then the array itself)
func WriteBytes(w io.Writer, order binary.ByteOrder, data []byte){
	var fieldLen int32
	fieldLen = int32(len(data))
	binary.Write(w, order, fieldLen)
	binary.Write(w, order, data)
}

// it reads a string from a buffer (packet received from the server)
func ReadString(w io.Reader, order binary.ByteOrder, data *string){
	var fieldLen int32
	var field []byte

	binary.Read(w, order, &fieldLen)
	if fieldLen <= 0{
		*data = ""
		return
	}
	field = make([]byte, fieldLen)
	binary.Read(w, order, &field)
	*data = string(field)
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

// it transforms the user tag to a long int represented as two int32 (the player ID)
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

// it transforms the player ID (two int32) to a tag
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
