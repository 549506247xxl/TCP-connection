package monitor

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"time"
)

var (
	snapShotLen int32 = 1024000
	err error
	timeout time.Duration = 30 * time.Second
	handle *pcap.Handle
)

// capture one specified device and output
func Capture(device string) {
	// open device
	handle, err = pcap.OpenLive(device, snapShotLen, false, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	var filter string = "tcp and port 8087"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	// use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		outputPacket(packet)
	}
}