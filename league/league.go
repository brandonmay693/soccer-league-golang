package league

import (
	"fmt"
	"github.com/brandonmay693/soccer-league-golang/models"
	"sort"
	"unicode/utf8"
)

type League interface {
	RankGameResult(a models.Score, b models.Score) models.RankedGameResult
	GenerateRankingTable([]models.MatchResult)
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
		fmt.Println(fmt.Sprintf("we have a tie! awarding both %s and %s %d points", a.Team, b.Team, l.TiePts))
		return models.RankedGameResult{
			TeamA: models.LeagueScore{
				Team:           a.Team,
				Score:          l.TiePts,
				GoalDifference: 0,
			},
			TeamB: models.LeagueScore{
				Team:           b.Team,
				Score:          l.TiePts,
				GoalDifference: 0,
			},
		}
	}

	if a.Score > b.Score {
		fmt.Println(fmt.Sprintf("%s beat %s, awarding %d and %d respectively", a.Team, b.Team, l.WinnerPts, l.LoserPts))
		return models.RankedGameResult{
			TeamA: models.LeagueScore{
				Team:           a.Team,
				Score:          l.WinnerPts,
				GoalDifference: a.Score - b.Score,
			},
			TeamB: models.LeagueScore{
				Team:           b.Team,
				Score:          l.LoserPts,
				GoalDifference: b.Score - a.Score,
			},
		}
	} else {
		fmt.Println(fmt.Sprintf("%s beat %s, awarding %d and %d respectively", b.Team, a.Team, l.WinnerPts, l.LoserPts))
		return models.RankedGameResult{
			TeamA: models.LeagueScore{
				Team:           a.Team,
				Score:          l.LoserPts,
				GoalDifference: a.Score - b.Score,
			},
			TeamB: models.LeagueScore{
				Team:           b.Team,
				Score:          l.WinnerPts,
				GoalDifference: b.Score - a.Score,
			},
		}
	}
}

type basicLeagueTable []models.LeagueScore

func (a basicLeagueTable) Len() int {
	return len(a)
}

func (a basicLeagueTable) Less(i, j int) bool {

	if a[i].Score == a[j].Score {
		iRune, _ := utf8.DecodeRuneInString(a[i].Team)
		jRune, _ := utf8.DecodeRuneInString(a[j].Team)
		//this line ensures that if scores are equal, we compare the team names in ascending order
		return iRune < jRune
	}

	//this line ensures that we show scores from descending
	return a[i].Score > a[j].Score
}

func (a basicLeagueTable) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (l *BasicLeague) GenerateRankingTable(gameResults []models.MatchResult) {
	var scoreMap = map[string]models.LeagueScore{}
	for _, r := range gameResults {
		rgr := l.RankGameResult(r.TeamA, r.TeamB)

		handleAddScoreToMap(scoreMap, rgr.TeamA)
		handleAddScoreToMap(scoreMap, rgr.TeamB)
	}

	var leagueTable basicLeagueTable
	for _, element := range scoreMap {
		leagueTable = append(leagueTable, element)
	}

	sort.Sort(leagueTable)
	for i := 0; i < len(leagueTable); i++ {
		fmt.Println(fmt.Sprintf("%d. %s, %d pts", i+1, leagueTable[i].Team, leagueTable[i].Score))
	}
}

func handleAddScoreToMap(scoreMap map[string]models.LeagueScore, ls models.LeagueScore) {
	if val, ok := scoreMap[ls.Team]; ok {
		scoreMap[ls.Team] = models.LeagueScore{
			Team:           val.Team,
			Score:          val.Score + ls.Score,
			GoalDifference: val.GoalDifference + ls.GoalDifference,
		}
	} else {
		scoreMap[ls.Team] = ls
	}
}
