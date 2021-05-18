package test

import (
	"happyball-matcher/internal/event/request"
	"testing"
	"time"
)

func getPlayerMatchingReqMMessageBytes(playerID int) []byte {
	req := request.NewPlayerMatchingRequest(int32(playerID))
	return req.ToGMessageBytes()
}

func TestName(t *testing.T) {
	for i := 0; i < 10; i++ {
		data := getPlayerMatchingReqMMessageBytes(i)
		SendDataToMatcher(data)
		time.Sleep(200 * time.Millisecond)
	}
}
