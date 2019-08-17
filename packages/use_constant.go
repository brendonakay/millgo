package millgo

import (
	"fmt"
	"reflect"
)

type UseConstantRule struct {
	FieldName string
	Constant  string
}

func (t UseConstantRule) Process(v *AuditLog) error {
	// Use Reflect to get struct field dynamically
	AuditLogValue := reflect.ValueOf(v).Elem()
	AuditLogFieldValue := AuditLogValue.FieldByName(t.FieldName)

	if !AuditLogFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in AuditLog", t.FieldName)
	}

	AuditLogFieldValue.SetString(t.Constant)

	return nil
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
