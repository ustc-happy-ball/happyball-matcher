package main

import (
	"happyball-matcher/dgs"
	"happyball-matcher/internal/matching"
)

func init() {
	dgs.GlobalDgsInfo = dgs.NewDgsAddress("default", "dgs-srv")
	dgs.GlobalDgsInfo.PrintAddress()
}

func main() {
	matcher := matching.NewMatcher()
	matcher.Serv()
}
