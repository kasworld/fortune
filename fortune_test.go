package fortune

import (
	"testing"
)

func TestLoadFortuneFile(t *testing.T) {
	f := LoadFile("/usr/share/games/fortunes/computers")
	println(len(f))
	// for _, v := range f {
	// 	println(v)
	// }
}

func TestLoadDir(t *testing.T) {
	f := LoadDir("/usr/share/games/fortunes")
	println(len(f))
	// for _, v := range f {
	// 	println(v)
	// }
}
