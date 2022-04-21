package game

import (
	"github.com/brandonmay693/soccer-league-golang/models"
	"reflect"
	"testing"
)

func TestBasicLeague_RankGameResult(t *testing.T) {
	type fields struct {
		TiePts    int
		WinnerPts int
		LoserPts  int
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := &BasicLeague{
				TiePts:    tt.fields.TiePts,
				WinnerPts: tt.fields.WinnerPts,
				LoserPts:  tt.fields.LoserPts,
			}
			if got := ba.RankGameResult(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RankGameResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
