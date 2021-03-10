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
* [x] CreatePatient /patient POST
* [x] ReadPatient /patient/{id:[0-9]+} GET
* [X] UpdatePatient /patient/{id:[0-9]+} POST
* [x] DeletePatient /patient/{id:[0-9]+} DELETE
#### Patient symptoms
* [ ] GetPatientSymptoms /patient/{patientID:[0-9]+}/symptom GET
* [ ] GetPatientSymptom /patient/{patientID:[0-9]+}/symptom/{id:[0-9]+} GET
* [ ] CreatePatientSymptoms /patient/{patientID:[0-9]+}/symptom POST
* [ ] UpdatePatientSymptoms /patient/{patientID:[0-9]+}/symptom/{id:[0-9]+} POST

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