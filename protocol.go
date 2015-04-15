package protocol

import (
	"strconv"
)

type Request struct {
	Id   int
	Cmd  string
	Host string
}

type Answer struct {
	Id     int
	Status string
	Data   ContainerTop
	Error  string
}

type ContainerTop struct {
	Name     string
	LimitMb  int
	TmpUsage int
	Procs    ByMemory
}

type Proc struct {
	Pid     string
	Memory  string
	Command string
}

type ByMemory []Proc

func (a ByMemory) Len() int {
	return len(a)
}

func (a ByMemory) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByMemory) Less(i, j int) bool {
	ai, _ := strconv.Atoi(a[i].Memory)
	aj, _ := strconv.Atoi(a[j].Memory)
	return ai < aj
}
