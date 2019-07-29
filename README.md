# millgo
Go implementation of Mill

### TODO
 - [ ] Arg parser
 - [x] Scanner that reads CSV
 - [x] YAML parser
 - [ ] One proof of concept field operation
 - [x] Models for proof of concept file type
 - [x] CSV to Struct
   - [ ] Test on bad data
 - [x] Write a transofmration pipeline stage
 - [ ] Write an errored rows pipeline stage
 - [ ] Write a successful rows pipeline stage
 - [ ] Split up project packages per domain
   - This includes models, utils, all the things
 - Change print statements to loggers

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
