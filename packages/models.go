package millgo

const (
	AUDIT_LOG     = "AUDIT_LOG"
	EMPLOYEE_INFO = "EMPLOYEE INFO"
	EVIDENCE      = "EVIDENCE"
	PATIENT_INFO  = "PATIENT INFO"
	EOF           = "END OF FILE"
)

type Rule interface {
	Process(*AuditLog)
}

/* TODO
- turn this into JSON
*/
type AuditLog struct {
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

type YamlConfig struct {
	AuditLog struct {
		Timestamp    int `yaml:"timestamp"`
		PatientId    int `yaml:"patient_id"`
		EmployeeId   int `yaml:"employee_id"`
		AccessAction int `yaml:"access_action"`
	} `yaml:"audit_log"`

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
