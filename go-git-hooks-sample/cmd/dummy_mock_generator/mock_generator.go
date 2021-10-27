//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"log"
	"os"
)

/*
   1- Fill existing_mocks.txt with the names of the interfaces you want to mock under picus-digital/go directory (obeying the syntax specified in there)
   2- Make sure your working directory is picus-digital/go (you should be in picus-digital/go directory before running the script)
   3- Then you can run this script with following flags (both from terminal as in below and from IDE with providing required params and env).
	AVAILABLE COMMANDS:
		=> Mock all defined interfaces in existing_mocks.txt
			./main.go --mock-all

		=> Check all modified files between <CommitSHA1> <CommitSHA2> against existing mocks defined in existing_mocks.txt,
			autogenerate new ones if detected changes that requires existing mocks to be regenerated.

			./main.go --check-mocks-autogen 0f4d52427e373f4a38834fad98dea0867ea18838 66f4afca933c2a152bad0ebd843e2f96b8f22633

		=> Check if there exist any unstaged new mocks, if so stage them and commit them as 'Mocks Regenerated'
			./main.go --commit-unstaged-mocks
*/



var scriptFlags = map[string]func(){
	"--mock-all":              mockAll,
	"--check-mocks-autogen":   checkMocksAutoGenerate,
	"--commit-unstaged-mocks": checkAndCommitUnstagedMocks,
}

func main() {
	if len(os.Args) >= 2 {
		task, exists := scriptFlags[os.Args[1]]
		if exists {
			task()
			os.Exit(0)
		}
		logColored(fatal, "Invalid usage, usages: `./mock_generator --mock-all` or `./mock_generator.go --check-mocks-autogen <CommitSHA1> <CommitSHA2>` or ./mock_generator --commit-unstaged-mocks` aborting...")
	}
	logColored(fatal, "Invalid usage, usages: `./dummy_mock_generator --mock-all` or `./dummy_mock_generator.go --check-mocks-autogen <CommitSHA1> <CommitSHA2>` or ./dummy_mock_generator --commit-unstaged-mocks` aborting...")
}

func mockAll() {
	// Generate Mocks using mockery
	logColored(success, "Yeeey! You've successfully run script with --mock-all flag. Now I am supposed to regenerate all defined mocks again...")
}

func checkMocksAutoGenerate() {
	// Check commit signatures and determine whether a regeneration for any existing mock is needed
	logColored(success, "Yeeey! You've successfully run script with --check-mocks-autogen flag. Now I am supposed to check differences between SHAs passed and look for if any changes made on any existing mock.")
}

func checkAndCommitUnstagedMocks() {
	// Check if there exist any unstaged mocks, if so commit them
	logColored(success, "Yeeey! You've successfully run script with --commit-unstaged-mocks flag. Now I am supposed to look for any regenerated but uncommitted mocks, if I find any I am going to commit them.")
}

type logType string

const (
	debug   logType = "debug"
	info    logType = "info"
	success logType = "success"
	warn    logType = "warn"
	fatal   logType = "fatal"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func logColored(logType logType, logString string) {
	header := colorCyan + "mock_generator.go: " + colorReset
	switch logType {
	case debug:
		log.Print(header, logString)
	case info:
		log.Print(header, colorBlue, logString, colorReset)
	case success:
		log.Print(header, colorGreen, logString, colorReset)
	case warn:
		log.Print(header, colorYellow, logString, colorReset)
	case fatal:
		log.Fatal(header, colorRed, logString, colorReset)
	}
}
