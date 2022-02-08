package note_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/note"
)

func TestUnmarshalLink(t *testing.T) {
	var l note.Link

	str := `{"type":"tie","target":"C5"}`
	err := json.Unmarshal([]byte(str), &l)

	assert.NoError(t, err)
}
