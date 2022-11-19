package oib

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
)

func Generate() string {
	s := fmt.Sprintf("%010d", rand.Int63n(1e10))

	return fmt.Sprintf("%s%s", s, checksum(s))
}

func IsValid(code string) error {
	pattern := regexp.MustCompile("^[0-9]{11}$")
	if !pattern.MatchString(code) {
		return fmt.Errorf("invalid oib: bad format")
	}

	checksum := checksum(code)
	lastDigit := string(code[10])

	if checksum != lastDigit {
		return fmt.Errorf("invalid oib: bad checksum")
	}

	return nil
}

func checksum(code string) string {
	checksum := 10

	for i := 0; i < 10; i++ {
		n, _ := strconv.Atoi(string(code[i]))

		checksum += n
		checksum %= 10
		if checksum == 0 {
			checksum = 10
		}

		checksum *= 2
		checksum %= 11
	}

	checksum = 11 - checksum
	if checksum == 10 {
		checksum = 0
	}

	return fmt.Sprintf("%d", checksum)
}
