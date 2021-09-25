package lidar

import "encoding/binary"

type LidarPacket struct {
	Value []byte
}

func (l LidarPacket) Timestamp() uint32 {
	return binary.LittleEndian.Uint32(l.Value[1242:1246])
}

func (l LidarPacket) ReturnMode() byte {
	return l.Value[1246]
}

func (l LidarPacket) ProductId() byte {
	return l.Value[1247]
}

func (l LidarPacket) ProductModel() string {
	switch l.ProductId() {
	case 0x21:
		return "HDL-32E"
	case 0x22:
		return "VLP-16"
	case 0x24:
		return "Puck Hi-Res"
	case 0x28:
		return "VLP-32C"
	case 0x31:
		return "Velarray"
	case 0xA1:
		return "VLS-128"

	default:
		return ""
	}
}

func (l LidarPacket) Blocks() []LidarPacketBlock {
	blocks := make([]LidarPacketBlock, LB_COUNT12)

	start := LB_INDEX42
	for i := 0; i < len(blocks); i++ {
		blocks[i] = LidarPacketBlock{l.Value[start : start+LB_SIZE100]}
		start += LB_SIZE100
	}

	return blocks
}

func (l LidarPacket) Azimuths() []uint16 {
	azimuths := make([]uint16, LB_COUNT12)
	for i, lpb := range l.Blocks() {
		azimuths[i] = lpb.AzimuthI()
	}
	return azimuths
}
