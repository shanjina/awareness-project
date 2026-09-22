package main

import (
	"fmt"
	"strings"
)

// ANSI color / style codes
const (
	reset     = "\033[0m"
	bold      = "\033[1m"
	dim       = "\033[2m"
	red       = "\033[31m"
	green     = "\033[32m"
	yellow    = "\033[33m"
	blue      = "\033[34m"
	magenta   = "\033[35m"
	cyan      = "\033[36m"
	white     = "\033[37m"
	bgBlue    = "\033[44m"
	brightRed = "\033[91m"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func colorize(color, text string) string {
	return color + text + reset
}

func hr(char string, width int) string {
	return strings.Repeat(char, width)
}

func banner() {
	fmt.Println(colorize(cyan+bold, hr("=", 60)))
	fmt.Println(colorize(cyan+bold, "   🛡️   CYBERSECURITY AWARENESS QUIZ   🛡️"))
	fmt.Println(colorize(cyan+bold, hr("=", 60)))
}

func boxTitle(title string) {
	line := hr("─", len(title)+4)
	fmt.Println(colorize(blue, "┌"+line+"┐"))
	fmt.Println(colorize(blue, "│  ") + colorize(bold+white, title) + colorize(blue, "  │"))
	fmt.Println(colorize(blue, "└"+line+"┘"))
}

// progressBar renders something like: [██████░░░░] 6/10
func progressBar(current, total, width int) string {
	if total == 0 {
		total = 1
	}
	filled := (current * width) / total
	if filled > width {
		filled = width
	}
	bar := strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
	return fmt.Sprintf("[%s] %d/%d", colorize(green, bar), current, total)
}

func categoryTag(category string) string {
	return colorize(magenta+bold, "["+strings.ToUpper(category)+"]")
}

func pressEnterToContinue(reader func() string) {
	fmt.Print(colorize(dim, "\nPress Enter to continue..."))
	reader()
}
