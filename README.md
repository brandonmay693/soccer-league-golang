# soccer-league-golang
A small golang application, that reads in a set of game scores from a file and outputs a ranked leaderboard table
in order of descending score, and ascending team name if the score is level.


## Setup:
- Ensure you have golang 1.16+ and make (optional) installed on your machine

##Run
There are 2 ways to run this application:
Method 1, using standard go commands:
- Run `go install`
- Run `go run main.go` with the path to a properly formatted
file (like testInput.txt in the root of the project) as the first option to the script


Method 2, using make and environment variables:
- Set `SCORE_FILE_PATH` environment variable to a path of a properly formatted file (like testInput.txt in the root of the project) in the project root 
- run `make install` followed by `make start`
