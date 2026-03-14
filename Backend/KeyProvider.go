package main

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// RandomStringProvider is an interface for generating random strings.
// The caller can specify the desired length.
type RandomStringProvider interface {
	// RandomString returns a random string of the specified length.
	// The returned string is composed of upper- and lower-case letters.
	RandomString(length int) (string, error)
}

// defaultRandomStringProvider is the default implementation of RandomStringProvider.
// It uses crypto/rand to generate secure random characters.
type defaultRandomStringProvider struct{}

// NewRandomStringProvider returns the default RandomStringProvider implementation.
func NewRandomStringProvider() RandomStringProvider {
	return &defaultRandomStringProvider{}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandomString generates a random string of the given length.
// If length is less than 0, it returns an error.
func (p *defaultRandomStringProvider) RandomString(length int) (string, error) {
	if length < 0 {
		return "", errors.New("length must be non-negative")
	}

	if length == 0 {
		return "", nil
	}

	out := make([]rune, length)
	max := big.NewInt(int64(len(letters)))

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		out[i] = letters[n.Int64()]
	}

	return string(out), nil
}
