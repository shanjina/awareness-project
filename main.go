package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		clearScreen()
		banner()
		fmt.Println()
		fmt.Println(colorize(dim, "Test your knowledge of phishing, passwords, malware, social"))
		fmt.Println(colorize(dim, "engineering, and everyday online safety."))
		fmt.Println()
		fmt.Println("  " + colorize(cyan+bold, "1") + ") Start Quiz (10 questions)")
		fmt.Println("  " + colorize(cyan+bold, "2") + ") Choose a Category")
		fmt.Println("  " + colorize(cyan+bold, "3") + ") View High Scores")
		fmt.Println("  " + colorize(cyan+bold, "4") + ") How to Play")
		fmt.Println("  " + colorize(cyan+bold, "5") + ") Exit")
		fmt.Println()
		fmt.Print(colorize(yellow, "Choose an option: "))

		choice := strings.TrimSpace(readLine())
		switch choice {
		case "1":
			startQuizFlow(pickQuestions(QuestionBank, 10))
		case "2":
			categoryMenu()
		case "3":
			viewHighScores()
		case "4":
			howToPlay()
		case "5":
			clearScreen()
			fmt.Println(colorize(green+bold, "Stay safe out there. Goodbye! 👋"))
			return
		default:
			fmt.Println(colorize(red, "Invalid option."))
			pressEnterToContinue(readLine)
		}
	}
}

func categoryMenu() {
	cats := allCategories()
	clearScreen()
	banner()
	fmt.Println()
	boxTitle("Choose a Category")
	fmt.Println()
	for i, c := range cats {
		count := len(questionsInCategory(c))
		fmt.Printf("  %s) %s %s\n", colorize(cyan+bold, strconv.Itoa(i+1)), c, colorize(dim, fmt.Sprintf("(%d questions)", count)))
	}
	fmt.Printf("  %s) Back to main menu\n", colorize(cyan+bold, "0"))
	fmt.Println()
	fmt.Print(colorize(yellow, "Choose a category: "))

	choice := strings.TrimSpace(readLine())
	idx, err := strconv.Atoi(choice)
	if err != nil || idx < 0 || idx > len(cats) {
		fmt.Println(colorize(red, "Invalid option."))
		pressEnterToContinue(readLine)
		return
	}
	if idx == 0 {
		return
	}
	cat := cats[idx-1]
	pool := questionsInCategory(cat)
	startQuizFlow(pickQuestions(pool, len(pool)))
}

func startQuizFlow(questions []Question) {
	if len(questions) == 0 {
		fmt.Println(colorize(red, "No questions available."))
		pressEnterToContinue(readLine)
		return
	}
	result := runQuiz(questions)
	showSummary(result)

	fmt.Println()
	fmt.Print(colorize(yellow, "Save this score to the leaderboard? (y/n): "))
	if strings.ToLower(strings.TrimSpace(readLine())) == "y" {
		fmt.Print(colorize(yellow, "Enter your name: "))
		name := strings.TrimSpace(readLine())
		if name == "" {
			name = "Anonymous"
		}
		if err := saveScore(newScoreEntry(name, result.Score, result.Total)); err != nil {
			fmt.Println(colorize(red, "Could not save score: "+err.Error()))
		} else {
			fmt.Println(colorize(green, "Score saved!"))
		}
	}
	pressEnterToContinue(readLine)
}

func viewHighScores() {
	clearScreen()
	banner()
	fmt.Println()
	boxTitle("🏆 Top Scores")
	fmt.Println()

	scores := loadScores()
	if len(scores) == 0 {
		fmt.Println(colorize(dim, "No scores yet. Play a quiz to set the first record!"))
	} else {
		for i, s := range scores {
			pct := 0.0
			if s.Total > 0 {
				pct = float64(s.Score) / float64(s.Total) * 100
			}
			fmt.Printf("  %s  %-15s %d/%d  (%.0f%%)  %s\n",
				colorize(cyan+bold, fmt.Sprintf("%2d.", i+1)),
				s.Name, s.Score, s.Total, pct, colorize(dim, s.Date))
		}
	}
	fmt.Println()
	pressEnterToContinue(readLine)
}

func howToPlay() {
	clearScreen()
	banner()
	fmt.Println()
	boxTitle("How to Play")
	fmt.Println()
	fmt.Println("• Each round asks a multiple-choice question about a real-world")
	fmt.Println("  cybersecurity scenario (phishing, passwords, malware, and more).")
	fmt.Println("• Type A, B, C, or D and press Enter to answer.")
	fmt.Println("• Type Q at any time during a quiz to quit early.")
	fmt.Println("• After each answer you'll see whether you were right and why —")
	fmt.Println("  read the explanations, that's where the real learning happens!")
	fmt.Println("• At the end, you can save your score to the local leaderboard.")
	fmt.Println()
	pressEnterToContinue(readLine)
}
