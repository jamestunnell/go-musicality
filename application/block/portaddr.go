package block

import "strings"

type PortAddr struct {
	Block, Port string
}

const PortAddrSep = "."

func NewPortAddr(block, port string) *PortAddr {
	return &PortAddr{Block: block, Port: port}
}

func (pa *PortAddr) Parse(s string) bool {
	results := strings.Split(s, PortAddrSep)
	if len(results) != 2 {
		return false
	}

	pa.Block = results[0]
	pa.Port = results[1]

	return true
}

func (pa *PortAddr) String() string {
	return strings.Join([]string{pa.Block, pa.Port}, PortAddrSep)
}
