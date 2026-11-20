package robot

import (
	"fmt"
)

// See defs.go for other definitions

// Step 1
// Define N, E, S, W here.
const (
	N Dir = iota
	E
	S
	W
)

func Right() {
	Step1Robot.Dir = (Step1Robot.Dir + 1) % 4
}

func Left() {
	Step1Robot.Dir--
	if Step1Robot.Dir < 0 {
		Step1Robot.Dir += 4
	}
}

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case E:
		Step1Robot.X++
	case S:
		Step1Robot.Y--
	case W:
		Step1Robot.X--
	}
}

func (d Dir) String() string {
	switch d {
	case N:
		return "N"
	case E:
		return "E"
	case S:
		return "S"
	case W:
		return "W"
	default:
		return ""
	}
}

// Step 2
// Define Action type here.
type Action Command

func StartRobot(command chan Command, action chan Action) {
	for cmd := range command {
		action <- Action(cmd)
	}
	close(action)
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	for act := range action {
		switch act {
		case 'R':
			robot.Dir = (robot.Dir + 1) % 4
		case 'L':
			robot.Dir = (robot.Dir + 3) % 4
		case 'A':
			next := robot.Pos
			switch robot.Dir {
			case N:
				next.Northing++
			case E:
				next.Easting++
			case S:
				next.Northing--
			case W:
				next.Easting--
			}
			if inside(extent, next) {
				robot.Pos = next
			}
		}
	}
	report <- robot
}

func inside(r Rect, p Pos) bool {
	return p.Easting >= r.Min.Easting && p.Easting <= r.Max.Easting &&
		p.Northing >= r.Min.Northing && p.Northing <= r.Max.Northing
}

// Step 3
// Define Action3 type here.
type Action3 struct {
	Name string
	Command
	Done bool
}

func StartRobot3(name, script string, action chan Action3, log chan string) {
	for _, cmd := range []byte(script) {
		if cmd != 'R' && cmd != 'L' && cmd != 'A' {
			log <- fmt.Sprintf("robot %q: undefined command %q", name, cmd)
			break
		}
		action <- Action3{Name: name, Command: Command(cmd)}
	}
	action <- Action3{Name: name, Done: true}
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	index := make(map[string]int)
	for i, r := range robots {
		if r.Name == "" {
			log <- "robot with no name"
		}
		if _, dup := index[r.Name]; dup {
			log <- fmt.Sprintf("duplicate robot name %q", r.Name)
		}
		index[r.Name] = i
		if !inside(extent, r.Pos) {
			log <- fmt.Sprintf("robot %q placed outside the room at %v", r.Name, r.Pos)
		}
		for _, other := range robots[:i] {
			if other.Pos == r.Pos {
				log <- fmt.Sprintf("robots %q and %q placed at the same position %v", other.Name, r.Name, r.Pos)
			}
		}
	}

	for done := 0; done < len(robots); {
		act := <-action
		if act.Done {
			done++
			continue
		}
		i, ok := index[act.Name]
		if !ok {
			log <- fmt.Sprintf("action from unknown robot %q", act.Name)
			continue
		}
		robot := &robots[i].Step2Robot
		switch act.Command {
		case 'R':
			robot.Dir = (robot.Dir + 1) % 4
		case 'L':
			robot.Dir = (robot.Dir + 3) % 4
		case 'A':
			next := ahead(*robot)
			switch {
			case !inside(extent, next):
				log <- fmt.Sprintf("robot %q bumped into a wall at %v", act.Name, robot.Pos)
			case occupied(robots, next):
				log <- fmt.Sprintf("robot %q bumped into another robot at %v", act.Name, next)
			default:
				robot.Pos = next
			}
		}
	}
	rep <- robots
}

func ahead(r Step2Robot) Pos {
	next := r.Pos
	switch r.Dir {
	case N:
		next.Northing++
	case E:
		next.Easting++
	case S:
		next.Northing--
	case W:
		next.Easting--
	}
	return next
}

func occupied(robots []Step3Robot, p Pos) bool {
	for _, r := range robots {
		if r.Pos == p {
			return true
		}
	}
	return false
}
