# Backend

## API routes

* [x] Health /health GET

### Journal
* [x] Health /journal/health GET
* [x] CreateJournal /journal POST
* [x] ReadJournal /journal/{id:[0-9]+} GET
* [x] JournalByPatient /journal/byPatient/{username} GET
* [x] UploadJournalDocumentsToML /journal/ml POST
* [x] UploadSymptomstoML /journal/symptoms POST

#### Journal Documents
* [x] CreateJournalDocument /journal/document POST
* [x] ReadJournalDocument /journal/document/{id:[0-9]+} GET
* [x] UpdateJournalDocument /journal/document/{id:[0-9]+} POST
* [x] DeleteJournalDocument /journal/document/{id:[0-9]+} DELETE
* [x] JournalDocumentsByJournal /journal/document/byJournal/{id:[0-9]+} GET

### Patient

#### General
* [x] Health /patient/health GET
* [x] GetPatients /patient GET
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
* [x] GetPatientDiagnoses /patient/{patientID:[0-9]+}/diagnose GET
* [x] GetPatientDiagnose /patient/{patientID:[0-9]+}/diagnose/{id:[0-9]+} GET
* [x] UpdatePatientDiagnose /patient/{patientID:[0-9]+}/diagnose/{id:[0-9]+} POST
* [x] DeletePatientDiagnose /patient/{patientID:[0-9]+}/diagnose/{id:[0-9]+} DELETE
#### Patient symptoms
* [x] CreatePatientSymptom /patient/{patientID:[0-9]+}/diagnose/{diagnoseID:[0-9]+}/symptom POST
* [x] GetPatientSymptoms /patient/{patientID:[0-9]+}/diagnose/{diagnoseID:[0-9]+}/symptom GET
* [x] UpdatePatientSymptom /patient/{patientID:[0-9]+}/diagnose/{diagnoseID:[0-9]+}/symptom/{id:[0-9]+} POST
* [x] DeletePatientSymptom /patient/{patientID:[0-9]+}/diagnose/{diagnoseID:[0-9]+}/symptom/{id:[0-9]+} DELETE

### Booking
* [x] GetHealth /booking/health GET
* [x] CreateBooking /booking POST
* [x] ReadBooking /booking/{id:[0-9]+} GET
* [x] UpdateBooking /booking/{id:[0-9]+} POST
* [x] DeleteBooking /booking/{id:[0-9]+} DELETE
* [x] GetBookingsByPatient /booking/byPatient/{id:[0-9]+} GET
* [x] GetBookingsByEmployee /booking/byEmployee/{id:[0-9]+]} GET
* [x] GetAvailableTimesForDoctor /booking/availableTimesForDoctor POST

### Administration
* [x] GetUser /admin/{id:[0-9]+} GET
* [x] GetHospitals /admin/hospitals GET
* [x] GetDepartments /admin/departments GET
* [x] GetBeds /admin/beds GET
* [x] GetAvailableBeds /admin/availableBeds POST
* [x] GetDoctorsInHospital /admin/doctors/inHospital/{hospitalId:[0-9]+} GET


### IOT
* [x] UploadData /iot/uploadData?Key=sdfgasfa&Date=22 POST
* [x] ReadData /iot/{deviceID:[0-9]+} GET
* [x] ReadDataInTimeFrame /iot/readDataInTimeframe POST