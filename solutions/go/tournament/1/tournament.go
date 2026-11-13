package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value [5]int
}

type PairList []Pair

func (p PairList) Len() int      { return len(p) }
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool {
	if p[i].Value[4] == p[j].Value[4] {
		return strings.Compare(p[i].Key, p[j].Key) == 1
	}
	return p[i].Value[4] < p[j].Value[4]
}

func Tally(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	tracker := map[string][5]int{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		text := strings.Split(line, ";")
		if len(text) != 3 {
			return fmt.Errorf("invalid line")
		}
		team1, team2 := [5]int{}, [5]int{}
		if val, ok := tracker[text[0]]; ok {
			team1 = val
		}
		if val, ok := tracker[text[1]]; ok {
			team2 = val
		}

		if text[2] == "win" {
			team1[1]++
			team1[4] += 3
			team2[3]++
		} else if text[2] == "draw" {
			team1[2]++
			team2[2]++
			team1[4]++
			team2[4]++
		} else if text[2] == "loss" {
			team1[3]++
			team2[1]++
			team2[4] += 3
		} else {
			return fmt.Errorf("invalid line")
		}

		team1[0]++
		team2[0]++
		tracker[text[0]] = team1
		tracker[text[1]] = team2
	}

	result := make(PairList, len(tracker))
	i := 0
	for k, v := range tracker {
		result[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(result))

	line := fmt.Sprintf("%-31s|%3s |%3s |%3s |%3s |%3s\n",
		"Team", "MP", "W", "D", "L", "P")
	writer.Write([]byte(line))
	for _, p := range result {
		line := fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n",
			p.Key, p.Value[0], p.Value[1], p.Value[2], p.Value[3], p.Value[4])
		writer.Write([]byte(line))
	}
	return nil
}
