package lidar

import "encoding/binary"

const (
	LB_SIZE100 = 100
	LB_INDEX42 = 42
	LB_COUNT12 = 12
	LC_INDEX4  = 4
	LC_SIZE3   = 3
	LC_COUNT16 = 16
)

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


