package millgo

import (
	"fmt"
	"time"
)

type ChangeDateFormatRule struct {
	FieldName string
	NewDateFormat  string
}

// TODO
//	- Use reflection to lookup field
func (t ChangeDateFormatRule) Process(v *AuditLog) {
	switch t.FieldName {
	case "Timestamp":
		// time.Parse returns meta information about line being parsed
		// could be useful for debugging field op
		timestamp, _ := time.Parse(t.NewDateFormat, v.Timestamp)
		v.Timestamp = timestamp.Format(t.NewDateFormat)
	default:
		fmt.Printf("No field match for rule")
	}
}


