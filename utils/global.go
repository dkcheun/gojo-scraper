package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func CleanDuplicates(input []string) []string {
	unique := make(map[string]struct{})
	result := make([]string, 0)

	for _, item := range input {
		if _, found := unique[item]; !found {
			unique[item] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}

func CleanStringArray(inputStrings []string) []string {
	seenStrings := make(map[string]bool)
	outputStrings := []string{}

	for _, inputString := range inputStrings {
		inputString = strings.ToLower(inputString)

		if seenStrings[inputString] {
			continue
		}

		seenStrings[inputString] = true

		if strings.Contains(inputString, ",") {
			splitStrings := strings.Split(inputString, ",")
			outputStrings = append(outputStrings, splitStrings...)
		} else {
			outputStrings = append(outputStrings, inputString)
		}
	}

	return outputStrings
}

func CleanTitle(input string) string {
	input = CleanUnicode(input)
	input = strings.ToLower(input)

	chars := []string{
		"\\",
		"\"",
		"\t",
		"\n",
		"\f",
		"\r",
		"\a",
		"\v",
		"\b",
		">",
		"<",
		"~",
		".",
		",",
		"`",
		"'",
		":",
		";",
		"|",
		"}",
		"{",
		"_",
		"*",
		"]",
		"[",
		"(",
		")",
		"-",
		"+",
		"/",
		"「",
		"」",
		"!",
		"?",
		"@",
		"#",
		"$",
		"%",
		"^",
		"&",
	}
	for _, c := range chars {
		input = strings.ReplaceAll(input, c, " ")
	}
	input = strings.Join(strings.Fields(input), " ")
	input = strings.ReplaceAll(input, " ", "-")
	input = strings.ReplaceAll(input, "--", "-")

	pattern := regexp.MustCompile(`\n`)
	cleanedText := pattern.ReplaceAllString(input, "")

	return cleanedText
}

func CleanUnicode(input string) string {
	unicodes := []string{
		"\u200b",
		"\u200d",
		"\u200e",
		"\u200f",
		"\u00ad",
		"\u200c",
		"\u180e",
		"\u202a",
		"\u202b",
		"\u202d",
		"\u202e",
	}

	for _, u := range unicodes {
		input = strings.ReplaceAll(input, u, " ")
	}
	input = strings.Join(strings.Fields(input), " ")

	return input
}

func CleanResText(input string) string {
	text := input

	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, "  ", " ")
	input = strings.ReplaceAll(input, " ", "")
	if len(input) <= 4 {
		input = strings.ReplaceAll(input, "null", "")
		input = strings.ReplaceAll(input, "nil", "")
		text = input
	}

	return text
}

func CleanOverview(input string) string {
	pattern := regexp.MustCompile(`\n\n\[\]|\[[^\]]*]`)

	input = CleanUnicode(input)
	cleanedText := pattern.ReplaceAllString(input, "")

	return cleanedText
}

func CleanTag(input string) string {
	input = strings.ToLower(input)

	tags := []string{
		"maintenance",
		"to episode",
		"moved to",
		"tag",
		"element",
		"setting",
		"themes",
		"deleted",
		"-- ",
	}
	for _, t := range tags {
		if strings.Contains(input, t) {
			return ""
		}
	}

	return CleanUnicode(input)
}

func CleanRuntime(input string) string {
	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, "hours", "h ")
	input = strings.ReplaceAll(input, "hour", "h ")
	input = strings.ReplaceAll(input, "hr", "h ")
	input = strings.ReplaceAll(input, "minutes", "m ")
	input = strings.ReplaceAll(input, "minute", "m ")
	input = strings.ReplaceAll(input, "min", "m ")
	input = strings.ReplaceAll(input, "seconds", "s ")
	input = strings.ReplaceAll(input, "second", "s ")
	input = strings.ReplaceAll(input, "sec", "s ")
	input = strings.Join(strings.Fields(input), " ")
	return input
}

func CleanRating(input string) (string, error) {
	bracketPattern := regexp.MustCompile(`\([^)]*\)`)

	cleanedText := bracketPattern.ReplaceAllString(input, "")

	cleanedText = strings.TrimSpace(cleanedText)
	cleanedText = strings.ReplaceAll(cleanedText, " ", "")
	return cleanedText, nil
}

func ExtractYear(input string) (int, error) {
	re := regexp.MustCompile(`(\d{4})`)
	matches := re.FindStringSubmatch(input)

	if len(matches) < 2 {
		return 0, fmt.Errorf("no four digits found")
	}
	year, err := strconv.ParseInt(matches[1], 0, 0)
	if err != nil {
		return 0, fmt.Errorf("cannot make four digits string a int")
	}
	return int(year), nil
}