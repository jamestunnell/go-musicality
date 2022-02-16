package note_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/pitch"
)

func TestLinkEqual(t *testing.T) {
	l1 := &note.Link{Source: pitch.C2, Target: pitch.D2, Type: note.LinkSlur}
	l2 := &note.Link{Source: pitch.E2, Target: pitch.D2, Type: note.LinkSlur}
	l3 := &note.Link{Source: pitch.C2, Target: pitch.E2, Type: note.LinkSlur}
	l4 := &note.Link{Source: pitch.C2, Target: pitch.D2, Type: note.LinkStep}

	assert.True(t, l1.Equal(l1))
	assert.False(t, l1.Equal(l2))
	assert.False(t, l1.Equal(l3))
	assert.False(t, l1.Equal(l4))
}

func TestLinksFindBySource(t *testing.T) {
	l1 := &note.Link{Source: pitch.C2, Target: pitch.D2, Type: note.LinkSlur}
	l2 := &note.Link{Source: pitch.E2, Target: pitch.D2, Type: note.LinkSlur}

	links := note.Links{l1, l2}

	l, found := links.FindBySource(pitch.Db2)

	assert.False(t, found)
	assert.Nil(t, l)

	l, found = links.FindBySource(l1.Source)

	require.True(t, found)
	require.NotNil(t, l)

	assert.True(t, l.Equal(l1))

	l, found = links.FindBySource(l2.Source)

	require.True(t, found)
	require.NotNil(t, l)

	assert.True(t, l.Equal(l2))
}

func TestLinksEqual(t *testing.T) {
	l1 := &note.Link{Source: pitch.C2, Target: pitch.D2, Type: note.LinkSlur}
	l2 := &note.Link{Source: pitch.E2, Target: pitch.D2, Type: note.LinkSlur}
	l3 := &note.Link{Source: pitch.E2, Target: pitch.F2, Type: note.LinkStep}

	links1 := note.Links{l1, l2}
	links2 := note.Links{l1, l2, l3}
	links3 := note.Links{l2, l1}
	links4 := note.Links{l2, l3}

	assert.True(t, links1.Equal(links1))
	assert.False(t, links1.Equal(links2))
	assert.False(t, links2.Equal(links1))
	assert.True(t, links1.Equal(links3))
	assert.True(t, links3.Equal(links1))
	assert.False(t, links1.Equal(links4))
	assert.False(t, links4.Equal(links1))
}
