package support

import (
	"github.com/ArtisanCloud/go-libs/str"
	"strings"
)

func QuickRandom(length int) string {
	pool := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// repeat the string
	strRepeat := strings.Repeat(pool, length)

	// shuffle the string
	strShuffle := str.Shuffle(strRepeat)

	return strShuffle[0:length]
}
