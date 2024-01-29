package capture

import (
	"fmt"
	"log"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func StartCapture(interfaceNome string) error {
	handle, err := pcap.OpenLive(interfaceNome, 1600, true, pcap.BlockForever)
	if err != nil {
		return fmt.Errorf("error when opening the interface: %w", err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for {
		select {
		case packet := <-packetSource.Packets():
			dataPackage := packet.Data()

			fmt.Printf("Package capture with timestamp: %s\n", packet.Metadata().Timestamp)
			fmt.Printf("Package length: %d bytes\n", packet.Metadata().Length)
			fmt.Printf("Data package: %v\n", dataPackage)
			fmt.Println("-----------")
		}
	}
}