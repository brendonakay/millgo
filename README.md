# millgo
Go implementation of Mill

### TODO
 - [ ] Benchmark against Python
 - [ ] Arg parser
 - [ ] Finish YAML parsing logic.
   - Field Ops parsing
 - [ ] Test on bad data
 - [ ] Write an errored rows pipeline stage
 - [ ] Split up project packages per domain! Reorganize.
   - This includes models, utils, all the things
   - Put domain specific files in respective directories
 - [ ] Change print statements to loggers
 - [ ] Implement exclude list
 - [x] Use reflection for dynamic struct field lookup from string
 - [x] Scanner that reads CSV
 - [x] YAML parser
 - [x] Write a successful rows pipeline stage
 - [x] Write a transofmration pipeline stage
 - [x] One proof of concept field operation
 - [x] Models for proof of concept file type
 - [x] CSV to Struct

### Dependencies
 - YAML
   - go get gopkg.in/yaml.v2

### Notes
 - stageOneChan
   - (Extract) Unmarshalls data from file to appropriate struct (e.g.; AuditLog, PatientInfo, Evidence, Etc...)
 - stageTwoChan
   - (Transform) Perform field operations on line. Divert bad lines to errorChan.
 - stageThreeChan
   - (Load) Append lines to CSV for uploader to consume. Or, upload via HTTP to tool in channel.

Field operations are methods of the operation type that implement the Rule interface. For example...
```Go
type ChangeDateFormatRule struct {
	FieldName string
	NewDateFormat  string
}
func (t ChangeDateFormatRule) Process(v *AuditLog) {
	// Transformation logic goes here
}
```
