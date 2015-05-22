// Copyright 2015 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
