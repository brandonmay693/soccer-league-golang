package main

import (
	"github.com/Netflix/go-env"
	"github.com/brandonmay693/soccer-league-golang/consumer"
	l "github.com/brandonmay693/soccer-league-golang/league"
	"github.com/brandonmay693/soccer-league-golang/models"
)

func main() {
	var config models.Environment
	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		panic(err)
	}

	c := consumer.NewFileConsumer(config.ScoreInputFilePath)

	gameResults, err := c.Read()
	if err != nil {
		return
	}

	league := l.NewBasicLeague(1, 3, 0)

	league.GenerateRankingTable(gameResults)
}
