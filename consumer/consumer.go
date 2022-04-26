package consumer

import (
	"bufio"
	"github.com/brandonmay693/soccer-league-golang/models"
	"log"
	"os"
	"strconv"
	"strings"
)

type Consumer interface {
	Read() error
}

type FileConsumer struct {
	FileLocation string
}

func NewFileConsumer(fileLocation string) FileConsumer {
	return FileConsumer{
		FileLocation: fileLocation,
	}
}

func (f FileConsumer) Read() ([]models.MatchResult, error) {
	file, err := os.Open(f.FileLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var returnVal []models.MatchResult
	for scanner.Scan() {
		teams := strings.Split(scanner.Text(), ",")

		teamA := strings.Split(teams[0], " ")
		teamB := strings.Split(teams[1], " ")

		teamAScore, err2 := parseScore(teamA)
		if err2 != nil {
			return returnVal, err2
		}
		teamBScore, err3 := parseScore(teamB)
		if err3 != nil {
			return returnVal, err3
		}
		returnVal = append(returnVal, models.MatchResult{
			TeamA: teamAScore,
			TeamB: teamBScore,
		})
	}
	return returnVal, nil
}

func parseScore(teamA []string) (models.Score, error) {
	s, err := strconv.Atoi(teamA[len(teamA)-1])

	if err != nil {
		return models.Score{}, err
	}

	return models.Score{
		Team:  strings.TrimSpace(strings.Join(teamA[:len(teamA)-1], " ")),
		Score: s,
	}, nil
}

//StdInConsumer is a Noop for demonstration of extensibility purposes
type StdInConsumer struct {
}

func (StdInConsumer) Read() error {
	return nil
}
