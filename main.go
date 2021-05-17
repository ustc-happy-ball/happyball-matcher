package main

import (
	"flag"
	"happyball-matcher/configs"
	"happyball-matcher/dgs"
	"happyball-matcher/internal/matching"
	"log"
)

func initAddress() {
	var (
		dgsPort string
		dgsHost string
	)

	flag.StringVar(&dgsHost, "DgsHost", "", "Host addr of dbproxy")
	flag.StringVar(&dgsPort, "DgsPort", "", " Port of dbproxy")
	// -DgsHost localhost -DgsPort 9000
	flag.Parse()
	configs.DgsAddr = dgsHost + ":" + dgsPort
}

func main() {
	initAddress()
	if configs.DgsAddr == ":" {
		log.Fatalln("dgs addr is nil")
	}
	log.Println("[dgs] dgs地址初始化为", configs.DgsAddr)
	go dgs.InitConnection(configs.DgsAddr)
	matcher := matching.NewMatcher()
	matcher.Serv()
}
