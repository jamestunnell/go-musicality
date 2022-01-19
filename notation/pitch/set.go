package pitch

import "sort"

type Set struct {
	pitches Pitches
}

func NewSet(pitches ...*Pitch) *Set {
	ps := &Set{
		pitches: make(Pitches, 0, len(pitches)),
	}

	for _, p := range pitches {
		ps.Add(p)
	}

	return ps
}

func (ps *Set) Pitches() Pitches {
	return ps.pitches
}

func (ps *Set) Len() int {
	return len(ps.pitches)
}

func (ps *Set) Union(other *Set) *Set {
	union := make(Pitches, len(ps.pitches))

	for i, p := range ps.pitches {
		union[i] = p
	}

	for _, p := range other.pitches {
		if indexOf(union, p) == -1 {
			union = append(union, p)
		}
	}

	return &Set{pitches: union}
}

func (ps *Set) Intersect(other *Set) *Set {
	intersect := Pitches{}

	for _, p := range ps.pitches {
		if other.Contains(p) {
			intersect = append(intersect, p)
		}
	}

	return &Set{pitches: intersect}
}

func (ps *Set) Diff(other *Set) *Set {
	diff := Pitches{}

	for _, p := range ps.pitches {
		if !other.Contains(p) {
			diff = append(diff, p)
		}
	}

	return &Set{
		pitches: diff,
	}
}

func (ps *Set) Add(p *Pitch) {
	if ps.indexOf(p) == -1 {
		ps.pitches = append(ps.pitches, p)
	}
}

func (ps *Set) Remove(p *Pitch) bool {
	idx := ps.indexOf(p)
	if idx == -1 {
		return false
	}

	ps.pitches = append(ps.pitches[:idx], ps.pitches[idx+1:]...)

	return true
}

func (ps *Set) Sorted() Pitches {
	sort.Sort(ps.pitches)

	return ps.pitches
}

func (ps *Set) Contains(tgt *Pitch) bool {
	return ps.indexOf(tgt) != -1
}

func (ps *Set) indexOf(tgt *Pitch) int {
	return indexOf(ps.pitches, tgt)
}

func indexOf(pitches Pitches, tgt *Pitch) int {
	for i, p := range pitches {
		if p.Equal(tgt) {
			return i
		}
	}

	return -1
}
