package main

import (
	"fmt"	
	"go2p"
)

func initPeer() {
	localAddr := "localhost:7077"
	net := go2p.NewNetworkConnectionTCP(*localAddr, &map[string]func(peer *go2p.Peer, msg *go2p.Message){
		"msg": func(peer *go2p.Peer, msg *go2p.Message) {
			fmt.Printf("%s > %s\n", peer.RemoteAddress(), msg.PayloadGetString())
		},
    })
    
    net.OnPeer(func(p *go2p.Peer) {
		fmt.Printf("new peer: %s\n", p.RemoteAddress())
    })
    
    err := net.Start()
	if err != nil {
		panic(err)
    }

  defer net.Stop()
}

func main() {
  fmt.Println("Hello World")
}
