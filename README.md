# soccer-league-golang
A small golang application, that reads in a set of game scores from a file and outputs a ranked leaderboard table
in order of descending score, and ascending team name if the score is level.


## Setup:
- Ensure you have golang 1.16+ and make installed on your machine
- Set the `SCORE_FILE_PATH` environment variable to a properly formatted file (like testInput.txt in the root of the project)
- in the project root run `make install` followed by `make start`
