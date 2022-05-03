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
- Using extra variables for copying arguments could do with just referencing them
- Whatever happens after starting quiz SIGTERM or time up or some error, it must show score in the end and then exit
- `csv` package doesn't trim spaces itself, have to take care of that
- tick is \u2705 
- wrong is \u274c

## Second Attempt
- `Timer` sends a message to a channel after time expires
- A datatype for problems
- A goroutine to get answers which pushes the answer in a channel, so that answers are non-blocking and not keep the program waiting after time expires
- similar treatment to listen for SIGTERM, SIGKILL . Registered a channel to be notified when any such signal is received and then it will call 






