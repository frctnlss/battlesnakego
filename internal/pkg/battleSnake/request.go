package battleSnake

import (
	"reflect"
)

type Request struct {
	Game  game        `json:"game"`
	Turn  uint16      `json:"turn"`
	Board board       `json:"board"`
	You   battleSnake `json:"you"`
}

type game struct {
	Id      string  `json:"id"`
	Ruleset ruleset `json:"ruleset"`
	Map     string  `json:"map"`
	Timeout uint16  `json:"timeout"`
	Source  string  `json:"source"`
}

type ruleset struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type board struct {
	Height  uint8         `json:"height"`
	Width   uint8         `json:"width"`
	Food    []grid        `json:"food"`
	Hazards []grid        `json:"hazards"`
	Snakes  []battleSnake `json:"snakes"`
}

func (b *board) EdgeMovementClockwise(position *grid) string {
	bottom := uint8(0)
	left := uint8(0)
	top := b.Height - uint8(1)
	right := b.Width - uint8(1)
	bottomLeft := &grid{X: left, Y: bottom}
	bottomRight := &grid{X: right, Y: bottom}
	topLeft := &grid{X: left, Y: top}
	topRight := &grid{X: right, Y: top}

	if reflect.DeepEqual(position, bottomLeft) {
		return "up"
	} else if reflect.DeepEqual(position, bottomRight) {
		return "left"
	} else if reflect.DeepEqual(position, topLeft) {
		return "right"
	} else if reflect.DeepEqual(position, topRight) {
		return "down"
	}

	if position.X == left {
		return "up"
	} else if position.Y == bottom {
		return "left"
	} else if position.X == right {
		return "down"
	} else if position.Y == top {
		return "right"
	}

	return ""
}

type grid struct {
	X uint8 `json:"x"`
	Y uint8 `json:"y"`
}

type battleSnake struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	Health         uint8          `json:"health"`
	Body           []grid         `json:"body"`
	Latency        string         `json:"latency"`
	Head           grid           `json:"head"`
	Length         uint16         `json:"length"`
	Shout          string         `json:"shout"`
	Squad          string         `json:"squad"`
	Customizations customizations `json:"customizations"`
}

type customizations struct {
	Color string `json:"color"`
	Head  string `json:"head"`
	Tail  string `json:"tail"`
}
