package millgo

import (
	"fmt"
)

/* TODO
- Inspiration coming from https://stackoverflow.com/questions/6769020/go-map-of-functions
FIELD_OPS = {
    'constant': _use_constant,
    'date_fmt': _change_date_format,
    'map': _use_mapped_value,
    'remove': _remove_from_value,
    'substitute': _substitute_in_value,
    'truncate': _truncate_value
}
*/

type Rule interface {
	Process(*AuditLog)// alternatively use a pointer (and remove the return) if you want to change it in-place
}

type UseConstantRule struct {
	FieldName string
	Constant  string
}

//func getField(v *AuditLog, field string) string {
//    r := reflect.ValueOf(v)
//    //f := reflect.Indirect(r)
//    return string(r.String())
//}

func (t UseConstantRule) Process(v *AuditLog) {
	switch t.FieldName {
	case "AccessAction":
		v.AccessAction = t.Constant
	default: fmt.Printf("DEFAULT HIT IN SWITCH")
	}
}

// TODO
//	- Put all of these in their own files and use same pattern ase UseConstsnt
func changeDateFormat() {
	// TODO: implement
}

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

