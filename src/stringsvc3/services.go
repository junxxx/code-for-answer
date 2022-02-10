package main

import (
	"errors"
	"strings"
)

// StringService provides operations on strings
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// + StringService implemetation
type stringService struct{}

var ErrEmpty = errors.New("Empty string")

type ServiceMiddleware func(StringService) StringService

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}
