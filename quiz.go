package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var stdinReader = bufio.NewReader(os.Stdin)

func readLine() string {
	text, _ := stdinReader.ReadString('\n')
	return strings.TrimSpace(text)
}

// shuffledQuestions returns a shuffled copy of the given slice.
func shuffledQuestions(qs []Question) []Question {
	cp := make([]Question, len(qs))
	copy(cp, qs)
	rand.Shuffle(len(cp), func(i, j int) { cp[i], cp[j] = cp[j], cp[i] })
	return cp
}

// pickQuestions returns up to n questions from the given pool, shuffled.
func pickQuestions(pool []Question, n int) []Question {
	shuffled := shuffledQuestions(pool)
	if n > len(shuffled) {
		n = len(shuffled)
	}
	return shuffled[:n]
}

func allCategories() []string {
	seen := map[string]bool{}
	var cats []string
	for _, q := range QuestionBank {
		if !seen[q.Category] {
			seen[q.Category] = true
			cats = append(cats, q.Category)
		}
	}
	return cats
}

func questionsInCategory(category string) []Question {
	var out []Question
	for _, q := range QuestionBank {
		if q.Category == category {
			out = append(out, q)
		}
	}
	return out
}

type QuizResult struct {
	Score       int
	Total       int
	MissedQs    []Question
	MissedAnsws []int // the user's chosen (wrong) index, parallel to MissedQs
}

// runQuiz executes the quiz loop over the given questions and returns the result.
func runQuiz(questions []Question) QuizResult {
	result := QuizResult{Total: len(questions)}
	labels := []string{"A", "B", "C", "D"}

	for i, q := range questions {
		clearScreen()
		banner()
		fmt.Println()
		fmt.Println(progressBar(i, len(questions), 30) + "   Score: " + colorize(green+bold, strconv.Itoa(result.Score)))
		fmt.Println()
		fmt.Println(categoryTag(q.Category) + "  " + colorize(dim, "Difficulty: "+q.Difficulty))
		fmt.Println()
		fmt.Println(colorize(bold+white, q.Text))
		fmt.Println()
		for idx, opt := range q.Options {
			fmt.Printf("  %s) %s\n", colorize(cyan+bold, labels[idx]), opt)
		}
		fmt.Println()
		fmt.Print(colorize(yellow, "Your answer (A/B/C/D, or Q to quit): "))

		answer := strings.ToUpper(strings.TrimSpace(readLine()))
		if answer == "Q" {
			fmt.Println(colorize(red, "\nQuiz aborted early."))
			result.Total = i
			return result
		}

		chosenIdx := -1
		switch answer {
		case "A":
			chosenIdx = 0
		case "B":
			chosenIdx = 1
		case "C":
			chosenIdx = 2
		case "D":
			chosenIdx = 3
		}

		fmt.Println()
		if chosenIdx == q.CorrectIndex {
			result.Score++
			fmt.Println(colorize(green+bold, "✔ Correct!"))
		} else {
			result.MissedQs = append(result.MissedQs, q)
			result.MissedAnsws = append(result.MissedAnsws, chosenIdx)
			correctLabel := "?"
			if q.CorrectIndex >= 0 && q.CorrectIndex < len(labels) {
				correctLabel = labels[q.CorrectIndex]
			}
			fmt.Println(colorize(brightRed+bold, "✘ Incorrect.") + " The correct answer was " + colorize(green+bold, correctLabel) + ".")
		}
		fmt.Println(colorize(dim, "Why: ") + q.Explanation)

		pressEnterToContinue(readLine)
	}

	return result
}

func grade(score, total int) (string, string) {
	if total == 0 {
		return "N/A", white
	}
	pct := float64(score) / float64(total) * 100
	switch {
	case pct == 100:
		return "🏆 Security Champion", green + bold
	case pct >= 80:
		return "🛡️  Security Aware", green
	case pct >= 60:
		return "🔎 Getting There", yellow
	case pct >= 40:
		return "⚠️  Needs Improvement", yellow + bold
	default:
		return "🚨 High Risk — Please Review!", brightRed + bold
	}
}

func showSummary(result QuizResult) {
	clearScreen()
	banner()
	fmt.Println()
	boxTitle("Quiz Complete!")
	fmt.Println()

	pct := 0.0
	if result.Total > 0 {
		pct = float64(result.Score) / float64(result.Total) * 100
	}
	fmt.Printf("Score: %s / %d  (%.0f%%)\n", colorize(bold+green, strconv.Itoa(result.Score)), result.Total, pct)

	label, color := grade(result.Score, result.Total)
	fmt.Println("Rating: " + colorize(color, label))

	if len(result.MissedQs) > 0 {
		fmt.Println()
		fmt.Println(colorize(bold+underline(), "Questions to review:"))
		labels := []string{"A", "B", "C", "D"}
		for i, q := range result.MissedQs {
			fmt.Println()
			fmt.Println(categoryTag(q.Category) + " " + q.Text)
			yourAns := "skipped/invalid"
			if idx := result.MissedAnsws[i]; idx >= 0 && idx < len(labels) {
				yourAns = labels[idx] + ") " + q.Options[idx]
			}
			fmt.Println(colorize(red, "  Your answer: ") + yourAns)
			fmt.Println(colorize(green, "  Correct: ") + labels[q.CorrectIndex] + ") " + q.Options[q.CorrectIndex])
			fmt.Println(colorize(dim, "  "+q.Explanation))
		}
	} else if result.Total > 0 {
		fmt.Println()
		fmt.Println(colorize(green+bold, "Perfect score! Great job staying security-aware. 🎉"))
	}
}

func underline() string {
	return "\033[4m"
}
