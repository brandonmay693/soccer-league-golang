package main

import (
	"fmt"
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

	rankingTable := league.GenerateRankingTable(gameResults)

	for i := 0; i < rankingTable.Len(); i++ {
		fmt.Println(fmt.Sprintf("%d. %s, %d pts", i+1, rankingTable.Get(i).Team, rankingTable.Get(i).Score))
	}
}
