# Backend

## API routes

* [x] Health /health GET

### Authentication
* [x] Health /authentication/health GET
* [x] PatientLogin /authentication/patient/login POST
* [x] EmployeeLogin /authentication/patient/login POSt

### Journal
* [x] Health /journal/health GET
* [x] CreateJournal /journal POST
* [x] ReadJournal /journal/{id:[0-9]+} GET
* [x] UpdateJournal /journal/{id:[0-9]+} POST
* [x] DeleteJournal /journal/{id:[0-9]+} DELETE
* [x] JournalByPatient /journal/byPatient/{id:[0-9]+} GET
* [ ] UploadJournalDocumentsToML /journal/ml POST
* [ ] UploadSymptomstoML /journal/symptoms POST

#### Journal Documents
* [x] CreateJournalDocument /journal/document POST
* [x] ReadJournalDocument /journal/document/{id:[0-9]+} GET
* [x] UpdateJournalDocument /journal/document/{id:[0-9]+} POST
* [x] DeleteJournalDocument /journal/document/{id:[0-9]+} DELETE
* [x] JournalDocumentsByJournal /journal/document/byJournal/{id:[0-9]+} GET

### Patient

#### General
* [x] Health /patient/health GET
* [x] CreatePatient /patient POST
* [x] GetPatients /patient GET
* [x] GetPatient /patient/{id:[0-9]+} GET
* [X] UpdatePatient /patient/{id:[0-9]+} POST
* [x] DeletePatient /patient/{id:[0-9]+} DELETE

#### Diagnose
* [x] GetDiagnose /diagnose/{id:[0-9]+} GET
* [x] GetDiagnoses /diagnose GET

#### Symptom
* [X] GetSymptom /symptom/{id:[0-9]+} GET
* [X] GetSymptoms /symptom GET

#### PatientDiagnose
* [x] CreatePatientDiagnose /patient/{patientID:[0-9]+}/diagnose POST
* [ ] GetPatientDiagnoses /patient/{patientID:[0-9]+}/diagnose GET
* [ ] GetPatientDiagnose /patient/{patientID:[0-9]+}/diagnose/{id:[0-9]+} GET
* [ ] UpdatePatientDiagnose /patient/{patientID:[0-9]+}/diagnose/{id:[0-9]+} POST
* [ ] DeletePatientDiagnose /patient/{patientID:[0-9]+}/diagnose/{id:[0-9]+} DELETE
#### Patient symptoms
* [ ] CreatePatientSymptom /patient/{patientID:[0-9]+}/diagnose{diagnoseID:[0-9]+}/symptom POST
* [ ] GetPatientSymptoms /patient/{patientID:[0-9]+}/diagnose{diagnoseID:[0-9]+}/symptom GET
* [ ] GetPatientSymptom /patient/{patientID:[0-9]+}/diagnose{diagnoseID:[0-9]+}/symptom/{id:[0-9]+}  GET
* [ ] UpdatePatientSymptom /patient/{patientID:[0-9]+}/diagnose{diagnoseID:[0-9]+}/symptom/{id:[0-9]+} POST
* [ ] DeletePatientSymptom /patient/{patientID:[0-9]+}/diagnose{diagnoseID:[0-9]+}/symptom/{id:[0-9]+} DELETE

#### Patient diagnose
* [] GetPatientDiagnose /

### Booking
 * [ ] CreateBooking /booking POST
 * [ ] ReadBooking /booking/{id:[0-9]+} GET
 * [ ] UpdateBooking /booking/{id:[0-9]+} POST
 * [ ] DeleteBooking /booking/{id:[0-9]+} DELETE
 * [ ] GetBookingsByPatient /booking/byPatient/{id:[0-9]+} GET
 * [ ] GetBookingsByCreatingEmployee /booking/byEmployee/{id:[0-9]+]} GET
 * [ ] GetBookingsInTimeFrame /booking?TimeStart=02/01/2006 15:04:05&TimeEnd=02/01/2006 15:04:05 GET (Sidenote the time should be url encoded so 02%2F01%2F2006%2015%3A04%3A05)


### User administration
* [ ] GetEmployee /useradmin/{id:[0-9]+} GET

### IOT
* [ ] UploadData /iot POST
* [ ] ReadData /iot/{id:[0-9]+} GET
* [ ] ReadDataInTimeFrame /iot?TimeStart=02/01/2006 15:04:05&TimeEnd=02/01/2006 15:04:05 GET (Sidenote the time should be url encoded so 02%2F01%2F2006%2015%3A04%3A05) 