package millgo

/*
AUDIT_LOG = '"EVIDENCE"|"AUDIT LOG"|"{0[timestamp]}"|"{0[patient_id]}"|"{0[employee_id]}"
  |"{0[access_action]}"|"{0[employee_role]}"|"{0[machine_id]}"|"{0[data_field]}"
  |"{0[source]}"|"{0[doc_access_type]}"'
AUDIT_LOG_V2 = '"EVIDENCE"|"AUDIT LOG"|"{0[timestamp]}"|"{0[patient_id]}"|"{0[employee_id]}"
  |"{0[access_action]}"|"{0[employee_role]}"|"{0[machine_id]}"|"{0[data_field]}"
  |"{0[source]}"|"{0[doc_access_type]}"|"{0[encounter_details]}"|"{0[additional_data]}"'
PATIENT_INFO = '"PATIENT INFO"|"{field}"|"{timestamp}"|"{patient_id}"|"{value}"'
EMPLOYEE_INFO = '"EMPLOYEE INFO"|"{field}"|"{timestamp}"|"{employee_id}"|"{value}"'
*/

/* TODO
- turn this into JSON
*/
type AuditLog struct {
	Evidence           string
	AuditLog           string
	Timestamp          string
	PatientId          string
	EmployeeId         string
	AccessAction       string
	EmployeeRole       string
	MachineId          string
	DataField          string
	Source             string
	DocumentAccessType string
}

type AuditLogV2 struct {
	Evidence           string
	AuditLog           string
	Timestamp          string
	PatientId          string
	EmployeeId         string
	AccessAction       string
	EmployeeRole       string
	MachineId          string
	DataField          string
	Source             string
	DocumentAccessType string
	EncounterDetails   string
	AdditionalData     string
}

type PatientInfo struct {
	PatientInfo string
	Field       string
	TimeStamp   string
	PatientId   string
	Value       string
}

type EmployeeInfo struct {
	EmployeeInfo string
	Field        string
	TimeStamp    string
	EmployeeId   string
	Value        string
}
