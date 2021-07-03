package main

import (
	"errors"
	"strings"
)

/*
	==========================
	SERVICE DEFINITIONS
	==========================
*/
var ErrEmptyString = errors.New("empty string")

// StringService provides operations on strings
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// stringService is a concrete implementation of StringService
type stringService struct{}

/* the next two methods satisfy the internal StringService interface */
func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmptyString
	}

	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}
