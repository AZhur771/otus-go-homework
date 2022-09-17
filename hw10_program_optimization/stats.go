package hw10programoptimization

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/mailru/easyjson"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	result := make(DomainStat)

	re, err := regexp.Compile("\\." + domain)
	if err != nil {
		return nil, err
	}

	bytesScanner := bufio.NewScanner(r)
	bytesScanner.Split(ScanLines)
	for bytesScanner.Scan() {
		var user User

		if err := easyjson.Unmarshal(bytesScanner.Bytes(), &user); err != nil {
			return nil, fmt.Errorf("get users error: %w", err)
		}

		matched := re.Match([]byte(user.Email))

		if matched {
			num := result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]
			num++
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])] = num
		}
	}

	return result, nil
}
