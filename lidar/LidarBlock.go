package lidar

import "encoding/binary"

type LidarPacketBlock struct {
	value []byte
}

func (b LidarPacketBlock) AzimuthF() float32 {
	return float32(b.AzimuthI()) / 100
}

func (b LidarPacketBlock) AzimuthI() uint16 {
	return binary.LittleEndian.Uint16(b.value[2:4])
}

func (b LidarPacketBlock) Channels() []LidarChannel {
	channels := make([]LidarChannel, LC_COUNT16)

	start := LC_INDEX4
	for i := 0; i < len(channels); i++ {
		channels[i] = LidarChannel{
			b.value[start : start+LC_SIZE3],
		}
		start += LC_SIZE3
	}

	return channels
}
