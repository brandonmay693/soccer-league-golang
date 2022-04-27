package models

type Score struct {
	Team  string
	Score int
}

type MatchResult struct {
	TeamA Score
	TeamB Score
}

type LeagueScore struct {
	Team           string
	Score          int
	GoalDifference int
}

type RankedGameResult struct {
	TeamA LeagueScore
	TeamB LeagueScore
}

type Environment struct {
	ScoreInputFilePath string `env:"SCORE_FILE_PATH"`
}
