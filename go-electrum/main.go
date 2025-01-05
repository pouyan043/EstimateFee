package main

import (
	"fmt"
	"log"

	"github.com/d4l3k/go-electrum/electrum"
)

func main() {
	node := electrum.NewNode() /// Create a New Electrum Object.

	err := node.ConnectTCP("electrum1.cipig.net:50001") // Connecting to the Public ElectrumServer
	if err != nil {
		log.Fatalf("Error connecting to Electrum server: %v\n", err)
	}
	//// Estimate the Transaction Fee for The Next 6 Blocks
	fee, err := node.BlockchainEstimateFee(6)
	if err != nil {
		log.Fatalf("Error estimating fee: %v\n", err)
	} ///Print the Estimated fee in satoshi per byte
	fmt.Printf("Estimated transaction fee for the next 6 blocks: %f satoshis/byte\n", fee)

}
