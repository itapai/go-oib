package oib

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
)

func Generate() string {
	random := fmt.Sprintf("%010d", rand.Int63n(1e10))
	checksum := checksum(random)

	return fmt.Sprintf("%s%s", random, checksum)
}

func IsValid(code string) error {
	pattern := regexp.MustCompile("^[0-9]{11}$")
	if !pattern.MatchString(code) {
		return fmt.Errorf("invalid oib: bad format")
	}

	checksum := checksum(code)
	match := string(code[10])

	if checksum != match {
		return fmt.Errorf("invalid oib: bad checksum")
	}

	return nil
}

func checksum(code string) string {
	current := 10

	for i := 0; i < 10; i++ {
		n, _ := strconv.Atoi(string(code[i]))

		current += n
		current %= 10
		if current == 0 {
			current = 10
		}

		current *= 2
		current %= 11
	}

	current = 11 - current
	if current == 10 {
		current = 0
	}

	return fmt.Sprintf("%d", current)
}
