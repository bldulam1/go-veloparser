package main

import (
	"fmt"

	lp "github.com/bldulam1/go-veloparser/lidar"
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
		data := packet.Data()

		if len(data) == 1248 {
			lidarp := lp.LidarPacket{data}

			fmt.Println(lidarp.Timestamp(), lidarp.ReturnMode(), lidarp.ProductId(), lidarp.ProductModel())

			// for _, b := range data.blocks() {
			// 	for ci, c := range b.channels() {
			// 		fmt.Println(ci, b.azimuthF(), c.distance(), c.reflectivity(), data.timestamp())
			// 	}
			// }
		}
	}
}
