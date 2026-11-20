package grep

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Search(pattern string, flags, files []string) []string {
	var lineNumbers, filesOnly, ignoreCase, invert, wholeLine bool
	for _, f := range flags {
		switch f {
		case "-n":
			lineNumbers = true
		case "-l":
			filesOnly = true
		case "-i":
			ignoreCase = true
		case "-v":
			invert = true
		case "-x":
			wholeLine = true
		}
	}

	if ignoreCase {
		pattern = strings.ToLower(pattern)
	}
	matches := func(line string) bool {
		if ignoreCase {
			line = strings.ToLower(line)
		}

		var ok bool
		if wholeLine {
			ok = line == pattern
		} else {
			ok = strings.Contains(line, pattern)
		}
		return ok != invert
	}

	var out []string
	for _, name := range files {
		file, err := os.Open(name)
		if err != nil {
			continue
		}

		scanner := bufio.NewScanner(file)
		for n := 1; scanner.Scan(); n++ {
			line := scanner.Text()
			if !matches(line) {
				continue
			}
			if filesOnly {
				out = append(out, name)
				break
			}
			prefix := ""
			if len(files) > 1 {
				prefix = name + ":"
			}
			if lineNumbers {
				prefix += strconv.Itoa(n) + ":"
			}
			out = append(out, prefix+line)
		}
		file.Close()
	}
	return out
}
