package millgo

const (
	AUDIT_LOG     = "AUDIT_LOG"
	EMPLOYEE_INFO = "EMPLOYEE INFO"
	EVIDENCE      = "EVIDENCE"
	PATIENT_INFO  = "PATIENT INFO"
	EOF           = "END OF FILE"
)

type Rule interface {
	Process(*AuditLog) error
}

/* TODO
- turn this into JSON
*/
type AuditLog struct {
	EvidenceConstant   string `json:"Evidence_constant"`
	AuditLogConstant   string `json:"Audit_log_constant"`
	Timestamp          string `json:"Timestamp"`
	PatientId          string `json:"Patient_id"`
	EmployeeId         string `json:"Employee_id"`
	AccessAction       string `json:"Access_action"`
	EmployeeRole       string `json:"Employee_role"`
	MachineId          string `json:"Machein_id"`
	DataField          string `json:"Data_field"`
	Source             string `json:"Source"`
	DocumentAccessType string `json:"Document_access_type"`
}

type AuditLogV2 struct {
	EvidenceConstant   string `json:"evidence_constant"`
	AuditLogConstant   string `json:"audit_log_constant"`
	Timestamp          string `json:"timestamp"`
	PatientId          string `json:"patient_id"`
	EmployeeId         string `json:"employee_id"`
	AccessAction       string `json:"access_action"`
	EmployeeRole       string `json:"employee_role"`
	MachineId          string `json:"machein_id"`
	DataField          string `json:"data_field"`
	Source             string `json:"source"`
	DocumentAccessType string `json:"document_access_type"`
	EncounterDetails   string `json:"encounter_details"`
	AdditionalData     string `json:"additional_data"`
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

type YamlConfig struct {
	AuditLog struct {
		Timestamp    int `yaml:"Timestamp"`
		PatientId    int `yaml:"Patient_id"`
		EmployeeId   int `yaml:"Employee_id"`
		AccessAction int `yaml:"Access_action"`
	} `yaml:"audit_log"`

	FieldOps map[string]map[string]map[string]string `yaml:"field_operations"`

	//	PatientInfo struct {
	//		PatientId int
	//		Field     struct {
	//			Dob      int
	//			FullName int
	//		}
	//	}
	//
	//	EmployeeInfo struct {
	//		EmployeetId int
	//		Fields      struct {
	//			Department int
	//			FullName   int
	//		}
	//	}
}
