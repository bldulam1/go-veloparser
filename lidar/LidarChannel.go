package lidar

import "encoding/binary"

type LidarChannel struct {
	value []byte
}

func (c LidarChannel) Distance() uint16 {
	return binary.LittleEndian.Uint16(c.value[:2])
}

func (c LidarChannel) Reflectivity() uint8 {
	return uint8(c.value[2])
}
