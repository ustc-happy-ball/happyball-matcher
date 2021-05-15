package component

type MatchRule interface {
	IsMatch(players []*MatchPlayer) bool
}
