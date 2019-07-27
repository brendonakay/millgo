package millgo

const (
	AUDIT_LOG     = "AUDIT_LOG"
	EMPLOYEE_INFO = "EMPLOYEE INFO"
	EVIDENCE      = "EVIDENCE"
	PATIENT_INFO  = "PATIENT INFO"
)

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

type YamlConfig struct {
	AuditLog struct {
		Timestamp    int `yaml:"timestamp"`
		PatientId    int `yaml:"patient_id"`
		EmployeeId   int `yaml:"employee_id"`
		AccessAction int `yaml:"access_action"`
	} `yaml:"audit_log"`

	PatientInfo struct {
		PatientId int
		Field     struct {
			Dob      int
			FullName int
		}
	}

	EmployeeInfo struct {
		EmployeetId int
		Fields      struct {
			Department int
			FullName   int
		}
	}
}
