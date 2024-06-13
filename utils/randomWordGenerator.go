package utils

import (
	"github.com/IIITManjeet/golexicon"
)

func GenerateWord() string {
	return golexicon.NewLexicon().Generate()
}
