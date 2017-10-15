package utils

func Int24ToInt32(bs []byte) uint32 {
    return uint32(bs[2]) | uint32(bs[1])<<8 | uint32(bs[0])<<16
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
