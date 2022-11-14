package model

import "github.com/samber/lo"

type Priority int

const (
	high Priority = iota + 1
	middle
	low
)

var priorityMap = map[Priority]string{
	high:   "HIGH",
	middle: "MIDDLE",
	low:    "LOW",
}

func (p Priority) Value() string {
	v, ok := priorityMap[p]
	if !ok {
		panic("invalid priority value")
	}
	return v
}

func NewPriority(val string) Priority {
	v, ok := lo.FindKey(priorityMap, val)
	if !ok {
		panic("invalid priority value")
	}
	return v
}
