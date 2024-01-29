package capture

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	// "github.com/google/gopacket/layers"
)

func StartCapture(interfaceNome string) error {
	handle, err := pcap.OpenLive(interfaceNome, 1600, true, pcap.BlockForever)
	if err != nil {
		return fmt.Errorf("erro ao abrir a interface: %w", err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for {
		select {
		case packet := <-packetSource.Packets():
			InterpretPackage(packet)
			// fmt.Println("-----------")
		}
	}
}

func InterpretPackage(packet gopacket.Packet) {
	// access the Ethernet layer
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
		fmt.Printf("Endereço MAC de origem: %s\n", ethernetPacket.SrcMAC)
		fmt.Printf("Endereço MAC de destino: %s\n", ethernetPacket.DstMAC)
	}

	// access the IP layer
	networkLayer := packet.Layer(layers.LayerTypeIPv4)
	if networkLayer != nil {
		ipPacket, _ := networkLayer.(*layers.IPv4)
		fmt.Printf("Endereço IP de origem: %s\n", ipPacket.SrcIP)
		fmt.Printf("Endereço IP de destino: %s\n", ipPacket.DstIP)
	}

	// access the TCP layer
	transportLayer := packet.Layer(layers.LayerTypeTCP)
	if transportLayer != nil {
		tcpPacket, _ := transportLayer.(*layers.TCP)
		fmt.Printf("Porta de origem: %d\n", tcpPacket.SrcPort)
		fmt.Printf("Porta de destino: %d\n", tcpPacket.DstPort)
	}

	// Access application layer (HTTP)
	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		payload := applicationLayer.Payload()
		fmt.Println("Dados da aplicação (HTTP):")
		fmt.Printf("%s\n", hex.Dump(payload))

		// Verificar se a palavra "TESTE" está presente
		fmt.Printf("Dados da aplicação (HTTP): %s\n", payload)
		if strings.Contains(string(payload), "TESTE") {
			fmt.Println("Dados da aplicação (HTTP):")
			fmt.Printf("%s\n", hex.Dump(payload))
			fmt.Println("Palavra TESTE encontrada nos dados da aplicação!")
		}
	}

	// fmt.Println("-----------")
}