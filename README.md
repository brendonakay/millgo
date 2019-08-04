# millgo
Go implementation of Mill

### TODO
 - [ ] Arg parser
 - [ ] Finish YAML parsing logic.
   - Field Ops parsing
 - [ ] Test on bad data
 - [ ] Write an errored rows pipeline stage
 - [ ] Split up project packages per domain! Reorganize.
   - This includes models, utils, all the things
   - Put domain specific files in respective directories
 - [ ] Change print statements to loggers
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
