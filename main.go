package main

import (
	"github.com/brandonmay693/soccer-league-golang/consumer"
	l "github.com/brandonmay693/soccer-league-golang/league"
)

func main() {
	c := consumer.NewFileConsumer("/Users/brandonmay/Documents/softwareProjects/playground/soccer-league-golang/testInput.txt")

	gameResults, err := c.Read()
	if err != nil {
		return
	}

	league := l.NewBasicLeague(1, 3, 0)

	league.GenerateRankingTable(gameResults)

}
