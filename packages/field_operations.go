package millgo

import (
	"reflect"
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
	Process(AuditLog) AuditLog // alternatively use a pointer (and remove the return) if you want to change it in-place
}

type UseConstantRule struct {
	FieldName string
	Constant  string
}

func getField(v *AuditLog, field string) string {
    r := reflect.ValueOf(v)
    //f := reflect.Indirect(r)
    return string(r.String())
}

func (t UseConstantRule) Process(v AuditLog) AuditLog {
	// process here, sadly you probably need reflection to edit the field name provided string dynamically
	//var m = map[string]interface{}{
	//	"Evidence":            o.Evidence,
	//	"AuditLog":            o.AuditLog,
	//	"Timestamp":           o.Timestamp,
	//	"PatientId":           o.PatientId,
	//	"EmployeeId":          o.EmployeeId,
	//	"AccessAction":        o.AccessAction,
	//	"EmployeeRole":        o.EmployeeRole,
	//	"MachineId":           o.MachineId,
	//	"DataField":           o.DataField,
	//	"Source":              o.Source,
	//	"DodcumentAccessType": o.DocumentAccessType,
	//}

	fmt.Printf("%v", getField(&v, t.FieldName))
	switch getField(&v, t.FieldName) {
	case "AccessAction":
		fmt.Println("HIT!")
		v.AuditLog = t.Constant
	}

	return v
}

func useConstant(c string) string {
	return c
}

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

// TODO implement this as ClientRules
func MyFieldOps() map[string]interface{} {
	m := map[string]interface{}{
		"useConstant":       useConstant,
		"changeDateFormat":  changeDateFormat,
		"useMappedValue":    useMappedValue,
		"removeFromValue":   removeFromValue,
		"substituteInValue": substituteInValue,
		"truncateValue":     truncateValue,
	}
	return m
}
