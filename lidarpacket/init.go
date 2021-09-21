package lidarpacket

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
	value []byte
}

func (l LidarPacket) Timestamp() uint32 {
	return binary.LittleEndian.Uint32(l.value[1242:1246])
}

func (l LidarPacket) ReturnMode() byte {
	return l.value[1246]
}

func (l LidarPacket) ProductId() byte {
	return l.value[1247]
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

func (l LidarPacket) blocks() []LidarPacketBlock {
	blocks := make([]LidarPacketBlock, LB_COUNT12)

	start := LB_INDEX42
	for i := 0; i < len(blocks); i++ {
		blocks[i] = LidarPacketBlock{l.value[start : start+LB_SIZE100]}
		start += LB_SIZE100
	}

	return blocks
}

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
	channels := make([]LidarChannel, LB_COUNT12)

	start := LC_INDEX4
	for i := 0; i < len(channels); i++ {
		channels[i] = LidarChannel{
			b.value[start : start+LC_SIZE3],
		}
		start += LC_SIZE3
	}

	return channels
}

type LidarChannel struct {
	value []byte
}

func (c LidarChannel) Distance() uint16 {
	return binary.LittleEndian.Uint16(c.value[:2])
}

func (c LidarChannel) Reflectivity() uint8 {
	return uint8(c.value[2])
}
