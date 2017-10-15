package utils

func Int24ToInt32(bs []byte) uint32 {
    return uint32(bs[2]) | uint32(bs[1])<<8 | uint32(bs[0])<<16
}