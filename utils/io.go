package utils

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ReadListOfInts(year int, day int) (*[]int, error) {
	reader, err := GetInputReader(year, day)
	if err != nil {
		return nil, fmt.Errorf("unable to get input as list of int: %w", err)
	}
	scanner := bufio.NewScanner(reader)
	r := make([]int, 0, 10)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("unable to get input as list of int: %w", err)
		}
		r = append(r, i)
	}
	return &r, nil
}

func ReadListOfStrings(year int, day int) (*[]string, error) {
	reader, err := GetInputReader(year, day)
	if err != nil {
		return nil, fmt.Errorf("unable to get input as list of int: %w", err)
	}
	scanner := bufio.NewScanner(reader)
	r := make([]string, 0, 10)
	for scanner.Scan() {
		r = append(r, scanner.Text())
	}
	return &r, nil
}

func ReadGridOfBytes(year int, day int) (*[][]byte, error) {
	reader, err := GetInputReader(year, day)
	if err != nil {
		return nil, fmt.Errorf("unable to get input as list of int: %w", err)
	}
	scanner := bufio.NewScanner(reader)
	r := make([][]byte, 0, 10)
	for scanner.Scan() {
		r = append(r, []byte(scanner.Text()))
	}
	return &r, nil
}

func GetInputReader(year int, day int) (io.Reader, error) {
	cookie, err := getSessionCookie()
	if err != nil {
		return nil, fmt.Errorf("unable to load aoc input: %w", err)
	}
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to load aoc input: %w", err)
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: cookie})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to load aoc input: %w", err)
	}
	return resp.Body, nil
}

func getSessionCookie() (string, error) {
	cookie, err := ioutil.ReadFile("../aoc_session_cookie")
	if err != nil {
		return "", fmt.Errorf("unable to load aoc session cookie from file: %w", err)
	}
	return string(cookie), nil
}
