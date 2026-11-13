package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Record struct {
	Team   string
	Played int
	Won    int
	Drawn  int
	Lost   int
	Score  int
}

func (r *Record) Win() {
	r.Won++
	r.Played++
	r.Score += 3 // A won has 3 points
}

func (r *Record) Draw() {
	r.Drawn++
	r.Played++
	r.Score++ // A draw has 1 point
}

func (r *Record) Loss() {
	r.Lost++
	r.Played++
}

type RecordBoard map[string]*Record

func (r RecordBoard) getOrCreateRecord(name string) *Record {
	if r[name] == nil {
		r[name] = &Record{Team: name}
	}
	return r[name]
}

func (r RecordBoard) sort() RecordList {
	result := make(RecordList, 0)
	for _, v := range r {
		result = append(result, v)
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Score == result[j].Score {
			return result[i].Team < result[j].Team
		}
		return result[i].Score > result[j].Score
	})
	return result
}

type RecordList []*Record

func (p RecordList) Len() int      { return len(p) }
func (p RecordList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p RecordList) Less(i, j int) bool {
	if p[i].Score == p[j].Score {
		return strings.Compare(p[i].Team, p[j].Team) == -1
	}
	return p[i].Score < p[j].Score
}

func Tally(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	tracker := make(RecordBoard)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		text := strings.Split(line, ";")
		if len(text) != 3 {
			return fmt.Errorf("invalid line")
		}

		team1, team2, result := text[0], text[1], text[2]
		switch result {
		case "win":
			tracker.getOrCreateRecord(team1).Win()
			tracker.getOrCreateRecord(team2).Loss()
		case "loss":
			tracker.getOrCreateRecord(team1).Loss()
			tracker.getOrCreateRecord(team2).Win()
		case "draw":
			tracker.getOrCreateRecord(team1).Draw()
			tracker.getOrCreateRecord(team2).Draw()
		default:
			return fmt.Errorf("invalid line")
		}

	}

	line := fmt.Sprintf("%-31s|%3s |%3s |%3s |%3s |%3s\n",
		"Team", "MP", "W", "D", "L", "P")
	writer.Write([]byte(line))
	for _, p := range tracker.sort() {
		line := fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n",
			p.Team, p.Played, p.Won, p.Drawn, p.Lost, p.Score)
		writer.Write([]byte(line))
	}
	return nil
}
