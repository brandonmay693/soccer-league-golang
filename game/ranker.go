package game

import (
	"fmt"
	"github.com/brandonmay693/soccer-league-golang/models"
)

type League interface {
	RankGameResult(a models.Score, b models.Score) models.RankedGameResult
}

type BasicLeague struct {
	TiePts    int
	WinnerPts int
	LoserPts  int
}

func NewBasicLeague(tiePoints int, winnerPts int, loserPts int) BasicLeague {
	return BasicLeague{
		TiePts:    tiePoints,
		WinnerPts: winnerPts,
		LoserPts:  loserPts,
	}
}

//RankGameResult accepts 2 models.Score(s) and returns a models.RankedGameResult
//where the score inside the RankedGameResult for each team is the score based on the
func (l *BasicLeague) RankGameResult(a models.Score, b models.Score) models.RankedGameResult {

	if a.Score == b.Score {
		fmt.Println(fmt.Sprintf("we have a tie! awarding both teams %d points", l.TiePts))
		return models.RankedGameResult{
			TeamA: models.Score{
				Team:  a.Team,
				Score: l.TiePts,
			},
			TeamB: models.Score{
				Team:  a.Team,
				Score: l.TiePts,
			},
		}
	}

	if a.Score > b.Score {
		fmt.Println(fmt.Sprintf("%s beat %s, awarding %d and %d respectively", a.Team, b.Team, l.WinnerPts, l.LoserPts))
		return models.RankedGameResult{
			TeamA: models.Score{
				Team:  a.Team,
				Score: l.WinnerPts,
			},
			TeamB: models.Score{
				Team:  b.Team,
				Score: l.LoserPts,
			},
		}
	} else {
		fmt.Println(fmt.Sprintf("%s beat %s, awarding %d and %d respectively", b.Team, a.Team, l.WinnerPts, l.LoserPts))
		return models.RankedGameResult{
			TeamA: models.Score{
				Team:  a.Team,
				Score: l.LoserPts,
			},
			TeamB: models.Score{
				Team:  b.Team,
				Score: l.WinnerPts,
			},
		}
	}
}
