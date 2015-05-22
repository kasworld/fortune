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
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"github.com/kasworld/log"
)

func LoadFile(filename string) []string {
	fd, err := os.Open(filename)
	if err != nil {
		log.Error("err in open %v\n", err)
		return nil
	}
	defer fd.Close()
	rtn := make([]string, 0)
	scanner := bufio.NewScanner(fd)
	scanner.Split(scanFortune)
	for scanner.Scan() {
		s := scanner.Text()
		rtn = append(rtn, s)
	}
	if err := scanner.Err(); err != nil {
		log.Error("reading %v %v", filename, err)
	}
	return rtn
}

func LoadDir(dir string) []string {
	// dir := "/usr/share/games/fortunes"
	rtn := make([]string, 0)
	for _, v := range listFiles(dir) {
		l := LoadFile(v)
		rtn = append(rtn, l...)
	}
	return rtn
}

func listFiles(basedir string) []string {
	rcsLists := make([]string, 0)
	rcsWalk := func(ppath string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() && !strings.Contains(info.Name(), ".") {
			rcsLists = append(rcsLists, ppath)
		}
		return nil //errors.New(text)
	}
	filepath.Walk(basedir, rcsWalk)
	return rcsLists
}

////

var delim = []byte("%\n")

// scanFortune is a bufio.SplitFunc to use on an io.Reader from a fortune file
func scanFortune(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	idx := bytes.Index(data, delim)
	switch {
	case idx == -1 && !atEOF:
		break
	case idx == -1 && atEOF:
		advance = len(data)
		token = data
	case idx != -1:
		token = data[:idx]
		advance = idx + 2
	}
	return
}
