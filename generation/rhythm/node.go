package rhythm

import (
	"fmt"
	"strings"

	"github.com/jamestunnell/go-musicality/notation/rat"
)

type Node struct {
	Dur  rat.Rat
	Str  string
	Rest bool
	Subs []*Node
}

type SubdivideCallback func(i uint64, sub *Node)

func NewNode(dur rat.Rat) *Node {
	str := dur.String()

	return &Node{Dur: dur, Str: str, Rest: false, Subs: []*Node{}}
}

func (e *Node) Subdivide(n uint64, f SubdivideCallback) {
	subDur := e.Dur.Div(rat.FromUint64(n))
	subs := make([]*Node, n)

	for i := uint64(0); i < n; i++ {
		sub := NewNode(subDur)
		subs[i] = sub

		f(i, sub)
	}

	e.Subs = subs
}

func SubdivideDoNothing(i uint64, sub *Node) {}

func (e *Node) Visit(f func(*Node)) {
	f(e)

	for _, n := range e.Subs {
		f(n)
	}
}

func (e *Node) TerminalNodes() []*Node {
	if len(e.Subs) == 0 {
		return []*Node{e}
	}

	nodes := []*Node{}
	for _, n := range e.Subs {
		nodes = append(nodes, n.TerminalNodes()...)
	}

	return nodes
}

func (e *Node) Print() {
	// for the sub elems
	builders := map[int]*strings.Builder{}

	level := 0
	_ = e.print(level, true, builders)

	builder, found := builders[level]
	for found {
		fmt.Println(builder.String() + "|")

		level++

		builder, found = builders[level]
	}
}

func (e *Node) print(level int, firstSub bool, builders map[int]*strings.Builder) int {
	builder, found := builders[level]
	if !found {
		builder = &strings.Builder{}

		builders[level] = builder
	}

	nSubs := len(e.Subs)
	if nSubs == 0 {
		count := 1 + len(e.Str)

		if firstSub {
			builder.WriteString("| ")

			count += 2
		}

		builder.WriteString(e.Str)
		builder.WriteRune(' ')

		return count
	}

	count := 0
	for i, sub := range e.Subs {
		count += sub.print(level+1, i == 0, builders)
	}

	nSpacesLeft := count

	if firstSub {
		builder.WriteRune('|')

		nSpacesLeft -= 1
	}

	leftSpaces := strings.Repeat(" ", nSpacesLeft/2)

	builder.WriteString(leftSpaces)

	nSpacesLeft -= len(leftSpaces)

	builder.WriteString(e.Str)

	nSpacesLeft -= len(e.Str)

	builder.WriteString(strings.Repeat(" ", nSpacesLeft))

	return count
}
