package main

import (
	"fmt"

	"github.com/bldulam1/veloparser/lidarpacket"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	handle, err := pcap.OpenOffline("./.cache/sample.pcap")
	if err != nil {
		panic(err)
	}

	packets := gopacket.NewPacketSource(handle, handle.LinkType()).Packets()

	for packet := range packets {
		if len(packet.Data()) == 1248 {
			data := pkt.LidarPacket{
				packet.Data(),
			}

			fmt.Println(data.Timestamp(), data.ReturnMode(), data.ProductId(), data.ProductModel())

			// for _, b := range data.blocks() {
			// 	for ci, c := range b.channels() {
			// 		fmt.Println(ci, b.azimuthF(), c.distance(), c.reflectivity(), data.timestamp())
			// 	}
			// }
		}
	}
}
