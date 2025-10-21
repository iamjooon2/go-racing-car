# Mission - Racing Car

## 🔍 How to Proceed

- The mission consists of Functional Requirements, Programming Requirements, and Submission Requirements.
- You must strive to meet all three. Especially, before implementing, create a feature list and commit by feature unit.
- If something is not specified in the requirements, make your own reasonable judgment.

## 📮 Submission Method

- After completing the mission, submit it via GitHub.
- Fork the go-racing-car repository to your personal repository.
- Create a branch with your GitHub ID in your personal repository and implement the mission there.
- Once completed, create a Pull Request to the branch corresponding to your GitHub ID in the iamjooon2/go-racing-car repository.

### Test Execution Guide

- Run go version in the terminal to confirm that the version is 1.25.3.
- Run the command go test in the terminal and make sure all tests pass as shown below.

```
=== RUN   TestHello
--- PASS: TestHello (0.00s)
PASS
```

🚀 Functional Requirements

Implement a simple racing car game.
- For a given number of attempts, n cars can either move forward or stop.
- Each car must have a name, and the name must be displayed when it moves forward.
- Car names are separated by commas (,), and each name must not exceed 8 characters.
- The user must be able to input how many rounds the race will run.
- A car moves forward when a random number between 0 and 9 is greater than or equal to 4.
- After the race, the program should announce the winner(s). There can be one or multiple winners.
- If there are multiple winners, they should be separated by commas (,).
- If invalid input is entered, the program should panic and terminate. (To be revised later.)

### Input/Output Requirements

##### Input

- Names of cars participating in the race (separated by commas)
```
joon, moritz, ali
```

-  Number of attempts
```
5
```
#### Output
- Result after each round

```
joon : --
moritz : ----
ali : ---
```

- Message for a single winner
```
Final winner: moritz
```

- Message for multiple winners
```
Final winners: moritz, ali
```

#### Example Execution
```
Enter the names of the cars to race. (Names are separated by commas)
joon,moritz,ali
How many attempts?
5

Race result
joon : -
moritz :
ali : -

joon : --
moritz : -
ali : --

joon : ---
moritz : --
ali : ---

joon : ----
moritz : ---
ali : ----

joon : -----
moritz : ----
ali : -----

Final winners: joon, ali
```
--- 
## 🎯 Programming Requirements

- The program must run on Go version 1.25.3.
- The entry point of the program is the main() function in application.go.
- Follow the [Official Go Coding Guide](https://go.dev/doc/effective_go) and [Banksalad Go Coding Convention](https://blog.banksalad.com/tech/go-best-practice-in-banksalad/)
- Before committing, always use `gofmt` to format your code.
- Once the implementation is complete, all tests must pass using the go test command.
- Unless otherwise specified, do not modify or move file or package names.

## ✏️ Submission Requirements

- Start the mission by Forking & Cloning the go-racing-car repository.
- Before implementation, list the features to be implemented in docs/README.md.
- Each commit must correspond to a single feature listed in docs/README.md.
- Follow the Commit Message Convention  guide when writing commit messages.
- 