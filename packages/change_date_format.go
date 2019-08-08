package millgo

import (
	"fmt"
	"reflect"
	"time"
)

type ChangeDateFormatRule struct {
	FieldName      string
	NewDateFormat  string
	OrigDateFormat string
}

// TODO
//	- Use reflection to lookup field
func (t ChangeDateFormatRule) Process(v *AuditLog) error {
	// Use Reflect to get struct field dynamically
	AuditLogValue := reflect.ValueOf(v).Elem()
	AuditLogFieldValue := AuditLogValue.FieldByName(t.FieldName)

	if !AuditLogFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in AuditLog", t.FieldName)
	}

	// time.Parse() returns meta information about what's left over from string
	//	could be useful for debugging
	timestamp, _ := time.Parse(t.OrigDateFormat, AuditLogFieldValue.String())

	AuditLogFieldValue.SetString(timestamp.Format(t.NewDateFormat))

	return nil
}
