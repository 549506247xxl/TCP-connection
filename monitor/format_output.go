package monitor

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)


func outputPacket(packet gopacket.Packet) {

	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		ipPacket, _ := ipLayer.(*layers.IPv4)
		fmt.Printf("IP %s to %s\n", ipPacket.SrcIP, ipPacket.DstIP)
		//if string(ipPacket.SrcIP) != "127.0.0.1" || string(ipPacket.DstIP) != "127.0.0.1" {
		//	os.Exit(0)
		//}
	}

	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		tcpPacket, _ := tcpLayer.(*layers.TCP)
		fmt.Printf("Port %s to %s\n", tcpPacket.SrcPort, tcpPacket.DstPort)
		if tcpPacket.SYN == true && tcpPacket.ACK == false {
			fmt.Println("[first]")
			//fmt.Println("SYN=1 and ACK=0, a connection request packet")
			fmt.Printf("seq: %d\n", tcpPacket.Seq)
		} else if tcpPacket.SYN == true && tcpPacket.ACK == true {
			fmt.Println("[second]:")
			//fmt.Println("SYN=1 and ACK=1, a connection receive packet")
			fmt.Printf("seq: %d, ack: %d\n", tcpPacket.Seq, tcpPacket.Ack)
		} else {
			fmt.Printf("seq: %d, ack: %d\n",tcpPacket.Seq, tcpPacket.Ack)
		}
	}

	app := packet.ApplicationLayer()
	if app != nil {
		payload := string(app.Payload())
		fmt.Println("payload is ", len(payload))
		fmt.Println(payload)
	}
	fmt.Println()
}


func printPacketInfo(packet gopacket.Packet, format string) {
	switch format {
	case "eth":
		outputEth(packet)
	case "tcp":
		outputTCP(packet)
	case "ip":
		outputIP(packet)
	default:
		fmt.Println("protocol error")
	}
}

func outputEth(packet gopacket.Packet) {
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
		fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
		fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
		fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
	} else {
		fmt.Println("ethernet not detected")
	}
}

func outputTCP(packet gopacket.Packet) {
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		tcpPacket, _ := tcpLayer.(*layers.TCP)
		fmt.Printf("From port %d to %d\n", tcpPacket.SrcPort, tcpPacket.DstPort)
		fmt.Println("Sequence number: ", tcpPacket.Seq)
		fmt.Println("ack: ", tcpPacket.Ack)
	} else {
		fmt.Println("tcp not detected")
	}
}

func outputIP(packet gopacket.Packet) {
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		ipPacket, _ := ipLayer.(*layers.IPv4)
		fmt.Printf("From %s to %s\n", ipPacket.SrcIP, ipPacket.DstIP)
		fmt.Println("Protocol: ", ipPacket.Version)
	}
}

func outputPayload(packet gopacket.Packet) {
	app := packet.ApplicationLayer()
	if app != nil {
		fmt.Println("application layer information")
		fmt.Println(string(app.Payload()))
	}
}