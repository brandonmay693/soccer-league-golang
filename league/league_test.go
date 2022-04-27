package league

import (
	"github.com/brandonmay693/soccer-league-golang/models"
	"reflect"
	"sort"
	"testing"
)

type fields struct {
	TiePts    int
	WinnerPts int
	LoserPts  int
}

func TestBasicLeague_RankGameResult(t *testing.T) {
	defaultLeagueFields := fields{
		TiePts:    1,
		WinnerPts: 3,
		LoserPts:  0,
	}

	type args struct {
		a models.Score
		b models.Score
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   models.RankedGameResult
	}{
		{
			name:   "should award the winner of a game 3 points and positive goal difference",
			fields: defaultLeagueFields,
			args: args{
				a: models.Score{
					Team:  "Bafana Bafana",
					Score: 2,
				},
				b: models.Score{
					Team:  "Man United",
					Score: 1,
				},
			},
			want: models.RankedGameResult{
				TeamA: models.LeagueScore{
					Team:           "Bafana Bafana",
					Score:          3,
					GoalDifference: 1,
				},
				TeamB: models.LeagueScore{
					Team:           "Man United",
					Score:          0,
					GoalDifference: -1,
				},
			},
		},
		{
			name:   "should award the loser of a game 0 points and negative goal difference",
			fields: defaultLeagueFields,
			args: args{
				a: models.Score{
					Team:  "Man United",
					Score: 1,
				},
				b: models.Score{
					Team:  "Chelsea",
					Score: 2,
				},
			},
			want: models.RankedGameResult{
				TeamA: models.LeagueScore{
					Team:           "Man United",
					Score:          0,
					GoalDifference: -1,
				},
				TeamB: models.LeagueScore{
					Team:           "Chelsea",
					Score:          3,
					GoalDifference: 1,
				},
			},
		},
		{
			name:   "should award both teams of a game 1 point in a tie with zero goal difference",
			fields: defaultLeagueFields,
			args: args{
				a: models.Score{
					Team:  "Bafana Bafana",
					Score: 2,
				},
				b: models.Score{
					Team:  "Man United",
					Score: 2,
				},
			},
			want: models.RankedGameResult{
				TeamA: models.LeagueScore{
					Team:           "Bafana Bafana",
					Score:          1,
					GoalDifference: 0,
				},
				TeamB: models.LeagueScore{
					Team:           "Man United",
					Score:          1,
					GoalDifference: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := NewBasicLeague(tt.fields.TiePts, tt.fields.WinnerPts, tt.fields.LoserPts)

			if got := ba.RankGameResult(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RankGameResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasicLeague_GenerateRankingTable(t *testing.T) {
	type args struct {
		gameResults []models.MatchResult
	}
	defaultLeagueFields := fields{
		TiePts:    1,
		WinnerPts: 3,
		LoserPts:  0,
	}

	g := []models.MatchResult{
		{
			TeamA: models.Score{
				Team:  "Lions",
				Score: 3,
			},
			TeamB: models.Score{
				Team:  "Snakes",
				Score: 3,
			},
		},
		{
			TeamA: models.Score{
				Team:  "Tarantulas",
				Score: 1,
			},
			TeamB: models.Score{
				Team:  "FC Awesome",
				Score: 0,
			},
		},
		{
			TeamA: models.Score{
				Team:  "Lions",
				Score: 1,
			},
			TeamB: models.Score{
				Team:  "FC Awesome",
				Score: 1,
			},
		},
		{
			TeamA: models.Score{
				Team:  "Tarantulas",
				Score: 3,
			},
			TeamB: models.Score{
				Team:  "Snakes",
				Score: 1,
			},
		},
		{
			TeamA: models.Score{
				Team:  "Lions",
				Score: 4,
			},
			TeamB: models.Score{
				Team:  "Grouches",
				Score: 0,
			},
		},
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   sort.Interface
	}{
		{
			name:   "should generate a league table, in descending order of score, and ascending order of team name when score is level",
			fields: defaultLeagueFields,
			args: args{
				gameResults: g,
			},
			want: BasicLeagueTable{
				{Team: "Tarantulas",
					Score:          6,
					GoalDifference: 3},
				{Team: "Lions", Score: 5, GoalDifference: 4},
				{Team: "FC Awesome", Score: 1, GoalDifference: -1},
				{Team: "Snakes", Score: 1, GoalDifference: -2},
				{Team: "Grouches", Score: 0, GoalDifference: -4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &BasicLeague{
				TiePts:    tt.fields.TiePts,
				WinnerPts: tt.fields.WinnerPts,
				LoserPts:  tt.fields.LoserPts,
			}
			if got := l.GenerateRankingTable(tt.args.gameResults); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateRankingTable() = %+v, want %v", got, tt.want)
			}
		})
	}
}
