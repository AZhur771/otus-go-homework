package hw10programoptimization

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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

	bytesScanner := bufio.NewScanner(r)
	bytesScanner.Split(ScanLines)
	for bytesScanner.Scan() {
		var user User

		if err := easyjson.Unmarshal(bytesScanner.Bytes(), &user); err != nil {
			return nil, fmt.Errorf("get users error: %w", err)
		}

		matched := strings.HasSuffix(user.Email, "."+domain)

		if matched {
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])] = result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])] + 1
		}
	}

	return result, nil
}
