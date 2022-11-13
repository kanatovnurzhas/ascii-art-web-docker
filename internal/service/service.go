package service

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"strings"
)

func GetResult(banner, text string) (string, error) {
	content, err := ioutil.ReadFile("static/banners/" + banner + ".txt")
	if err != nil {
		return "", ErrorNoBanner
	}

	if !checkTheBanner(content, banner) {
		return "", ErrorBanner
	}

	text = strings.ReplaceAll(text, "\r\n", "\n")

	for _, r := range text {
		if r != 10 && (r < ' ' || r > '~') {
			return "", ErrorCharacter
		}
	}

	m := setSymbols(content)
	art := toString(getAsciiArt(m, text))

	return art, nil
}

func setSymbols(content []byte) map[rune][]string {
	arr := splitLines(strings.Split(string(content), "\n"))
	m := make(map[rune][]string)
	for i, w := range arr {
		m[rune(i+32)] = w
	}
	return m
}

func getAsciiArt(m map[rune][]string, s string) [][8]string {
	result := [][8]string{}
	lines := [8]string{}

	for i := 0; i < len(s); i++ {
		if s[i] != 10 {
			for j := range lines {
				lines[j] += m[rune(s[i])][j]
			}
		}
		if s[i] == 10 || i == len(s)-1 {
			result = append(result, lines)
			lines = [8]string{}
		}
	}

	return result
}

func splitLines(lines []string) [][]string {
	symbol := []string{}
	symbols := [][]string{}
	for i, line := range lines {
		if line != "" {
			symbol = append(symbol, line)
		}
		if (line == "" || i == len(lines)-1) && len(symbol) > 0 {
			symbols = append(symbols, symbol)
			symbol = []string{}
		}
	}
	return symbols
}

func toString(arr [][8]string) string {
	s := ""
	for _, w := range arr {
		for _, q := range w {
			if len(q) == 0 {
				s += "\n"
				break
			}
			s += q + "\n"
		}
	}
	return s
}

func checkTheBanner(data []byte, banner string) bool {
	banners := map[string]string{
		"shadow":     "a49d5fcb0d5c59b2e77674aa3ab8bbb1",
		"standard":   "ac85e83127e49ec42487f272d9b9db8b",
		"thinkertoy": "bf1d925662e40f5278b26a0531bfdb63",
	}
	return fmt.Sprintf("%x", md5.Sum(data)) == banners[banner]
}
