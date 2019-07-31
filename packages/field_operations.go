package millgo

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

func MyFieldOps() map[string]interface{} {
	m := map[string]interface{}{
		"useConstant": useConstant,
		"changeDateFormat": changeDateFormat,
		"useMappedValue": useMappedValue,
		"removeFromValue": removeFromValue,
		"substituteInValue": substituteInValue,
		"truncateValue": truncateValue,
	}
	return m
}
