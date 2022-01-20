package sequence

import "github.com/jamestunnell/go-musicality/notation/note"

const (
	AttackNone    = 0.0
	AttackNormal  = 0.25
	AttackTenuto  = 0.5
	AttackAccent  = 0.75
	AttackMarcato = 1.0
)

func Attack(articulation string) float64 {
	attack := AttackNormal

	switch articulation {
	case note.Normal, note.Portato:
		attack = AttackNormal
	case note.Tenuto, note.Staccato, note.Staccatissimo:
		attack = AttackTenuto
	case note.Accent, note.Marcato:
		attack = AttackAccent
	}

	return attack
}
