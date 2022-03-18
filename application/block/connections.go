package block

import (
	"hash/fnv"
	"math/rand"
	"time"

	"github.com/dcadenas/pagerank"
)

type Connections struct {
	pairs []*AddrPair
}

type AddrPair struct {
	From *Addr
	To   *Addr
}

func HashString(s string) uint32 {
	h := fnv.New32()

	h.Write([]byte(s))

	return h.Sum32()
}

func NewConnections() *Connections {
	return &Connections{
		pairs: []*AddrPair{},
	}
}

func (conns *Connections) Len() int {
	return len(conns.pairs)
}

func (conns *Connections) Connect(output, input *Addr) bool {
	for _, pair := range conns.pairs {
		if pair.From.Equal(output) || pair.To.Equal(input) {
			return false
		}
	}

	pair := &AddrPair{From: output, To: input}

	conns.pairs = append(conns.pairs, pair)

	return true
}

func (conns *Connections) EachConnection(f func(output, input *Addr)) {
	for _, pair := range conns.pairs {
		f(pair.From, pair.To)
	}
}

func (conns *Connections) ConnectedInput(output *Addr) (*Addr, bool) {
	for _, pair := range conns.pairs {
		if pair.From.Equal(output) {
			return pair.To, true
		}
	}

	return nil, false
}

func (conns *Connections) ConnectedOutput(input *Addr) (*Addr, bool) {
	for _, pair := range conns.pairs {
		if pair.To.Equal(input) {
			return pair.From, true
		}
	}

	return nil, false
}

func (conns *Connections) RankBlocks() map[string]float64 {
	const (
		followProb = 0.95
		tolerance  = 0.01
	)

	// use PageRank algorithm to determine Ordinals
	graph := pagerank.New()
	blockLabels := map[int]string{}

	// hook up the graph
	for _, pair := range conns.pairs {
		outLabel := int(HashString(pair.From.Block))
		inLabel := int(HashString(pair.To.Block))

		blockLabels[outLabel] = pair.From.Block
		blockLabels[inLabel] = pair.To.Block

		graph.Link(outLabel, inLabel)
	}

	ranks := map[string]float64{}

	rand.Seed(time.Now().Unix())

	// run the PageRank algorithm
	graph.Rank(followProb, tolerance, func(label int, rank float64) {
		ranks[blockLabels[label]] = rank
	})

	return ranks
}
