package parsinglogfiles

import "regexp"

var validLineRegex = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
var splitLogRegex = regexp.MustCompile(`<[~*=-]*>`)
var quotedPasswordRegex = regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
var endOfLineRegex = regexp.MustCompile(`end-of-line\d+`)
var userRegex = regexp.MustCompile(`User\s+(\S+)`)

func IsValidLine(text string) bool {
	return validLineRegex.MatchString(text)
}

func SplitLogLine(text string) []string {
	return splitLogRegex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		if quotedPasswordRegex.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return endOfLineRegex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	result := make([]string, 0, len(lines))

	for _, line := range lines {
		matches := userRegex.FindStringSubmatch(line)
		if matches == nil {
			result = append(result, line)
			continue
		}

		username := matches[1]
		tagged := "[USR] " + username + " " + line
		result = append(result, tagged)
	}

	return result
}
