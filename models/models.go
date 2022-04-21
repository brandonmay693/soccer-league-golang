package models

type Score struct {
	Team  string
	Score int
}

type RankedGameResult struct {
	TeamA Score
	TeamB Score
}
