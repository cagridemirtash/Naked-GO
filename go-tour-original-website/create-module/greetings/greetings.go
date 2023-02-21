package greetings

import (
	e "errors"
	f "fmt"
)

func Hello(name string) (string, error) {
	//When name is blank string, return an error which is empty name
	if name == "" {
		return "", e.New("empty name")
	}
	message := f.Sprintf("Hi, %v. Welcome!", name)
	//When the assign is ok, return message and nil(for handle errors).
	return message, nil
}
