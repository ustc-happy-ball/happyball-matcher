package main

import (
	"happyball-matcher/configs"
	"happyball-matcher/dgs"
	"happyball-matcher/internal/matching"
	"log"
)

func init() {
	dgs.GlobalDgsInfo = dgs.NewDgsAddress("default","dgs-srv")
	dgs.GlobalDgsInfo.PrintAddress()
}

func initAddress() {
	//var (
	//	dgsPort string
	//	dgsHost string
	//)

	//flag.StringVar(&dgsHost, "DgsHost", "", "Host addr of dgs")
	//flag.StringVar(&dgsPort, "DgsPort", "", " Port of dgs")
	//// -DgsHost localhost -DgsPort 9000
	//flag.Parse()

	//configs.DgsAddr = dgs.GlobalDgsInfo.Address[0].InternalIP + ":" + dgs.GlobalDgsInfo.Address[0].InternalPort
	configs.DgsAddr = dgs.GlobalDgsInfo.Address[0].InternalIP + ":" + dgs.GlobalDgsInfo.Address[0].InternalPort
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
