package main

import (
	"encoding/json"
	"os"
	"sort"
	"time"
)

const scoreFile = "highscores.json"

// ScoreEntry represents one completed quiz attempt.
type ScoreEntry struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
	Total int    `json:"total"`
	Date  string `json:"date"`
}

func loadScores() []ScoreEntry {
	data, err := os.ReadFile(scoreFile)
	if err != nil {
		return []ScoreEntry{}
	}
	var scores []ScoreEntry
	if err := json.Unmarshal(data, &scores); err != nil {
		return []ScoreEntry{}
	}
	return scores
}

func saveScore(entry ScoreEntry) error {
	scores := loadScores()
	scores = append(scores, entry)

	sort.Slice(scores, func(i, j int) bool {
		pi := float64(scores[i].Score) / float64(max1(scores[i].Total))
		pj := float64(scores[j].Score) / float64(max1(scores[j].Total))
		return pi > pj
	})

	if len(scores) > 10 {
		scores = scores[:10]
	}

	data, err := json.MarshalIndent(scores, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(scoreFile, data, 0644)
}

func max1(n int) int {
	if n == 0 {
		return 1
	}
	return n
}

func newScoreEntry(name string, score, total int) ScoreEntry {
	return ScoreEntry{
		Name:  name,
		Score: score,
		Total: total,
		Date:  time.Now().Format("2006-01-02 15:04"),
	}
}
