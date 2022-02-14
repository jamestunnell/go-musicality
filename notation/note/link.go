package note

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/pitch"
)

const (
	LinkTie         = "tie"
	LinkSlur        = "slur"
	LinkGlide       = "glide"
	LinkStep        = "step"
	LinkStepSlurred = "stepSlurred"
)

type Link struct {
	Source, Target *pitch.Pitch
	Type           string
}

type Links []*Link

type linkLite struct {
	Target *pitch.Pitch `json:"target"`
	Type   string       `json:"type"`
}

type linkLiteMap map[string]*linkLite

func (link *Link) Equal(other *Link) bool {
	return link.Source.Equal(other.Source) && link.Type == other.Type && link.Target.Equal(other.Target)
}

func (links Links) ToLinkLiteMap() linkLiteMap {
	m := linkLiteMap{}

	for _, l := range links {
		m[l.Source.String()] = &linkLite{
			Target: l.Target,
			Type:   l.Type,
		}
	}

	return m
}

func (links Links) FindBySource(p *pitch.Pitch) (*Link, bool) {
	for _, l := range links {
		if l.Source.Equal(p) {
			return l, true
		}
	}

	return nil, false
}

func (links Links) Equal(other Links) bool {
	if len(links) != len(other) {
		return false
	}

	for _, link := range links {
		link2, found := other.FindBySource(link.Source)
		if !found {
			return false
		}

		if !link.Equal(link2) {
			return false
		}
	}

	return true
}

func (m linkLiteMap) ToLinks() (Links, error) {
	links := Links{}

	for srcStr, ll := range m {
		src, err := pitch.Parse(srcStr)
		if err != nil {
			err = fmt.Errorf("failed to parse source pitch string '%s': %w", srcStr, err)

			return Links{}, err
		}

		l := &Link{
			Source: src,
			Target: ll.Target,
			Type:   ll.Type,
		}

		links = append(links, l)
	}

	return links, nil
}
