package bowling

import "errors"

type Frame [2]int

// Define the Game type here.
type Game struct {
	roll      int
	frames    [10]Frame
	tenth     []int // balls rolled in the 10th frame
	completed bool
}

func NewGame() *Game {
	return &Game{roll: 0}
}

func (g *Game) Roll(pins int) error {
	if pins < 0 || pins > 10 || g.completed {
		return errors.New("Pin count exceeds pins on the lane")
	}
	frame := min(g.roll/2, 9) // current frame, 0–9
	for i := frame - 2; i < frame; i++ {
		if i >= 0 && g.frames[i][0] > 0 {
			g.frames[i][1] += pins
			g.frames[i][0]--
		}
	}
	if g.roll >= 18 {
		t := g.tenth
		if (len(t) == 1 && t[0] < 10 && t[0]+pins > 10) || // 2nd ball must fit what's left on the lane
			(len(t) == 2 && t[0] == 10 && t[1] < 10 && t[1]+pins > 10) { // strike, then non-strike
			return errors.New("Pin count exceeds pins on the lane")
		}
		g.tenth = append(g.tenth, pins)
		g.frames[9][1] += pins
		g.roll++
		if len(g.tenth) == 3 || (len(g.tenth) == 2 && g.tenth[0]+g.tenth[1] < 10) {
			g.completed = true
		}
		return nil
	}
	if pins == 10 {
		if g.roll%2 == 0 {
			g.frames[int(g.roll/2)] = Frame{2, pins}
			g.roll += 2
		} else {
			g.frames[int(g.roll/2)] = Frame{1, pins}
			g.roll++
		}
	} else {
		total := g.frames[int(g.roll/2)][1] + pins
		if total > 10 {
			return errors.New("Pin count exceeds pins on the lane")
		}
		g.frames[int(g.roll/2)][1] = total
		if total == 10 {
			g.frames[int(g.roll/2)][0] = 1
		}
		g.roll++
	}

	return nil
}

func (g *Game) Score() (int, error) {
	if !g.completed {
		return 0, errors.New("Game is not completed.")
	}

	total := 0
	for _, f := range g.frames {
		total += f[1]
	}
	return total, nil
}
