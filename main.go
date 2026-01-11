package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

var font = map[rune][]string{
	'H': {
		"1   1",
		"1   1",
		"1   1",
		"11111",
		"1   1",
		"1   1",
		"1   1",
	},
	'I': {
		"11111",
		"  1  ",
		"  1  ",
		"  1  ",
		"  1  ",
		"  1  ",
		"11111",
	},
	'R': {
		"11110",
		"1   1",
		"1   1",
		"11110",
		"1 1  ",
		"1  1 ",
		"1   1",
	},
	'E': {
		"11111",
		"1    ",
		"1    ",
		"11110",
		"1    ",
		"1    ",
		"11111",
	},
	'M': {
		"1   1",
		"11 11",
		"1 1 1",
		"1   1",
		"1   1",
		"1   1",
		"1   1",
	},
	' ': {
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
	},
}

func main() {
	message := "HIRE ME"
	minCommits := 9

	now := time.Now()
	startDate := now.AddDate(0, 0, -365)
	for startDate.Weekday() != time.Sunday {
		startDate = startDate.AddDate(0, 0, 1)
	}

	currentDate := startDate

	fmt.Println("Generating git history for:", message)

	for _, char := range message {
		pattern, exists := font[char]
		if !exists {
			currentDate = currentDate.AddDate(0, 0, 7)
			continue
		}

		width := len(pattern[0])

		for col := range width {
			for row := range 7 {
				if col < len(pattern[row]) && pattern[row][col] == '1' {
					commitDate := currentDate.AddDate(0, 0, row)
					numCommits := rand.Intn(2) + minCommits
					for range numCommits {
						makeCommit(commitDate)
					}
				}
			}
			currentDate = currentDate.AddDate(0, 0, 7)
		}
		currentDate = currentDate.AddDate(0, 0, 7)
	}

	fmt.Println("Done! Push the repo to see changes.")
}

func makeCommit(date time.Time) {
	dateStr := date.Format("2006-01-02 12:00:00")
	cmd := exec.Command("git", "commit", "--allow-empty", "-m", "contribution", "--date", dateStr)
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("GIT_AUTHOR_DATE=%s", dateStr),
		fmt.Sprintf("GIT_COMMITTER_DATE=%s", dateStr),
	)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error committing for date %s: %v\n", dateStr, err)
	}
}
