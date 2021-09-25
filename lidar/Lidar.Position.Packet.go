package lidar

import "encoding/binary"

func NewPositionPacket(data []byte) LidarPositionPacket {
	return LidarPositionPacket{
		Message:   string(data[0xf8 : 0x14d+1]),
		TimeStamp: binary.LittleEndian.Uint32(data[0xf0 : 0xf0+4]),
		PPS:       uint8(data[0xf4]),
	}
}

type LidarPositionPacket struct {
	TimeStamp uint32
	Message   string
	PPS       uint8
}
