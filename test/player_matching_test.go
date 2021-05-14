package test

import (
	"happyball-matcher/internal/event/request"
	"testing"
)

func getPlayerMatchingReqMMessageBytes() []byte {
	req := request.NewPlayerMatchingRequest(123)
	return req.ToGMessageBytes()
}

func TestName(t *testing.T) {
	data := getPlayerMatchingReqMMessageBytes()
	SendDataToMatcher(data)
}
