package main

import (
	"fmt"

	"github.com/bldulam1/go-veloparser/lidar"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	for frame := range generateFrame() {
		fmt.Println(frame)
		break
	}
}

func generateBlocks() <-chan lidar.LidarPacketBlock {
	ch := make(chan lidar.LidarPacketBlock)

	go func() {
		defer close(ch)

		handle, err := pcap.OpenOffline("./.cache/sample.pcap")
		if err != nil {
			panic(err)
		}

		for packet := range gopacket.NewPacketSource(handle, handle.LinkType()).Packets() {
			data := packet.Data()
			if len(data) == lidar.LP_BYTE_LEN {
				blocks := lidar.LidarPacket{data}.Blocks()
				for _, block := range blocks {
					ch <- block
				}
			}
		}
	}()

	return ch
}

func generateFrame() <-chan []lidar.LidarPacketBlock {
	ch := make(chan []lidar.LidarPacketBlock)

	go func() {
		defer close(ch)

		var firstAz uint16
		var currAz uint16
		var prevAz uint16

		blocks := make([]lidar.LidarPacketBlock, 0)

		for block := range generateBlocks() {
			blocks = append(blocks, block)
			currAz = block.AzimuthI()

			if prevAz <= firstAz && firstAz <= currAz && len(blocks) > 2 {
				ch <- blocks
				blocks = make([]lidar.LidarPacketBlock, 0)
			}

			// Initialize Az0
			if firstAz == 0 {
				firstAz = currAz
			}

			// Save currAzimuth to prevAzimuth
			prevAz = currAz
		}
	}()

	return ch
}
