package millgo

import (
	"fmt"
)

type UseConstantRule struct {
	FieldName string
	Constant  string
}

// TODO
//	- Use reflection to lookup field
func (t UseConstantRule) Process(v *AuditLog) {
	switch t.FieldName {
	case "AccessAction":
		v.AccessAction = t.Constant
	default:
		fmt.Printf("No field match for rule")
	}
}

// TODO
//	- Put all of these in their own files and use same pattern ase UseConstsnt
func useMappedValue() {
	// TODO: implement
}

func removeFromValue() {
	// TODO: implement
}

func substituteInValue() {
	// TODO: implement
}

func truncateValue() {
	// TODO: implement
}
