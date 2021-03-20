# Backend

## API routes
* [x] CreateJournal /journal POST doctor
* [x] JournalByPatient /journal/byPatient/{username} GET doctor, nurse, patient
* [x] UploadJournalDocumentsToML /journal/ml POST doctor, nurse
* [x] UploadSymptomstoML /journal/symptoms POST doctor, nurse, office

* [x] CreateJournalDocument /journal/document POST  doctor, nurse
* [x] ReadJournalDocument /journal/document/{id:[0-9]+} GET doctor, nurse, patient
* [x] UpdateJournalDocument /journal/document/{id:[0-9]+} POST doctor, nurse
* [x] DeleteJournalDocument /journal/document/{id:[0-9]+} DELETE doctor, nurse
* [x] JournalDocumentsByJournal /journal/document/byJournal/{id:[0-9]+} GET doctor, nurse, patient

* [x] GetPatients /patient GET doctor, nurse, office
* [x] CreatePatientDiagnose /patient/{patientID:[0-9]+}/diagnose POST doctor
* [x] GetPatientDiagnoses /patient/{patientID:[0-9]+}/diagnose GET doctor nurse, patient
* [x] DeletePatientDiagnose /patient/{patientID:[0-9]+}/diagnose/{id:[0-9]+} DELETE doctor
* [x] CreatePatientSymptom /patient/{patientID:[0-9]+}/diagnose/{diagnoseID:[0-9]+}/symptom POST doctor

* [x] CreateBooking /booking POST all
* [x] DeleteBooking /booking/{id:[0-9]+} DELETE doctor, nurse, office
* [x] GetBookingsByPatient /booking/byPatient/{id:[0-9]+} GET doctor, nurse, office, patient
* [x] GetBookingsByEmployee /booking/byEmployee/{id:[0-9]+]} GET all
* [x] GetAvailableTimesForDoctor /booking/availableTimesForDoctor POST all

* [x] GetHospitals /admin/hospitals GET all
* [x] GetAvailableBeds /admin/availableBeds POST all
* [x] GetDoctorsInHospital /admin/doctors/inHospital/{hospitalId:[0-9]+} GET all

* [x] GetDiagnoses /diagnose GET all
* [x] GetSymptoms /symptom GET all

* [x] UploadData /iot/uploadData?Key=sdfgasfa&sensorId=1&Date=22 POST
* [x] ReadData /iot/{deviceID:[0-9]+} GET
* [x] ReadDataInTimeFrame /iot/readDataInTimeframe POST