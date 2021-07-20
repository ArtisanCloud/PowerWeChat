package support

import (
	"github.com/ArtisanCloud/go-libs/fmt"
	"testing"
)

func Test_QuickRandom(t *testing.T) {

	response := QuickRandom(32)
	fmt.Dump(response)

	response = QuickRandom(32)
	fmt.Dump(response)

	response = QuickRandom(32)
	fmt.Dump(response)

}
