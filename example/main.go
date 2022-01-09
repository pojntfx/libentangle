package main

import (
	"bufio"
	"log"
	"os"

	"github.com/alphahorizonio/libentangle/pkg/networking"
	"github.com/pion/webrtc/v3"
)

func main() {
	networking.Connect("test", func(msg webrtc.DataChannelMessage) { log.Printf("Message: %s", msg.Data) }, func() {})

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		networking.Write([]byte(text))
	}
}
