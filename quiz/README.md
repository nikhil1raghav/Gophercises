# Quiz App


## Usage
- `-limit` time limit in seconds for the quiz
- `-csv` path to the csv file that holds the quiz data

## To learn

- How to parse args in golang
- Use channels to simulate timer
- Read a csv file

## First attempt

- For timer, spinned up a goroutine which will sleep for the defined time limit and then trigger termination of program, problem is if it ends prematurely then there is no way to print score, should use a channel to end it gracefully
- Reading complete csv file at once, not good






