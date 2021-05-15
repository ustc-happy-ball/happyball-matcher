package rule

import (
	"happyball-matcher/internal/matching/component"
)

type PlayerCountMatchRule struct {
	CountRequirement int32
}

func NewPlayerCountMatchRule(countRequirement int32) *PlayerCountMatchRule {
	return &PlayerCountMatchRule{CountRequirement: countRequirement}
}

func (p *PlayerCountMatchRule) IsMatch(players []*component.MatchPlayer) bool {
	if int32(len(players)) >= p.CountRequirement {
		return true
	} else {
		return false
	}
}
