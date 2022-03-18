package block

import "strings"

type Addr struct {
	Block, Port string
}

const AddrSep = "."

func NewAddr(block, port string) *Addr {
	return &Addr{Block: block, Port: port}
}

func (a *Addr) Equal(other *Addr) bool {
	return a.Block == other.Block && a.Port == other.Port
}

// func (pa *Addr) Parse(s string) bool {
// 	results := strings.Split(s, AddrSep)
// 	if len(results) != 2 {
// 		return false
// 	}

// 	pa.Block = results[0]
// 	pa.Port = results[1]

// 	return true
// }

func (a *Addr) String() string {
	return strings.Join([]string{a.Block, a.Port}, AddrSep)
}
