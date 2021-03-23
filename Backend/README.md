# Installation
If you wanna run this program there are a few different things you need.
* MSSQL database, with the correct structure scaffolded, you can run the Initial.sql file from this repo to get the right structure

* Mongo database
* Okta account
* Machine learning endpoints running, look for the machineLearning project in this repo

When you have all of those you next need to set the different env variables up, here is a list of the once you need to change

```
MSSQL_URI=
MONGO_URI=

ML_DIAGNOSE_ENDPOINT=
ML_IMAGE_ENDPOINT=

OKTA_SDK_TOKEN=
OKTA_URL=
OKTA_CLIENT_ID=
OKTA_AUTH_ENDPOINT=

IS_DEV=TRUE
```

For the rest you can just use the default

```
HTTP_SERVICE_ADDR=0.0.0.0:8080
AUTHENTICATION_SERVICE_ADDR=0.0.0.0:9000
JOURNAL_SERVICE_ADDR=0.0.0.0:9001
BOOKING_SERVICE_ADDR=0.0.0.0:9002
PATIENT_SERVICE_ADDR=0.0.0.0:9003
USERADMIN_SERVICE_ADDR=0.0.0.0:9004
IOT_SERVICE_ADDR=0.0.0.0:9005

AUTHENTICATION_REMOTE_ADDR=127.0.0.1:9000
JOURNAL_REMOTE_ADDR=127.0.0.1:9001
BOOKING_REMOTE_ADDR=127.0.0.1:9002
PATIENT_REMOTE_ADDR=127.0.0.1:9003
USERADMIN_REMOTE_ADDR=127.0.0.1:9004
IOT_REMOTE_ADDR=127.0.0.1:9005

IS_DEV=TRUE
```

When all of that is done you need to make sure you have the go tooling installed and properly setup on the machine you are doing this on.

```
go mod download
```

Once thats all done its a matter of going into the services and running
```
go run
```

If all done correctly you can access the HTTP/JSON endpoint via the address you used from the env, default is localhost:8080

So access http://localhost:8080/health and you should see this response
```
🚀 Server is up and running!!!!
```


# Structure
## HTTP/JSON Service
All of the handlers for the different endpoints are all located in the services/httpService/server package
They start with the root name of their URL path and then Handler. so For example JournalByPatient would be /journal/byPatient/{usernamer}
therefore it would be in the journalHandler.go

the authentication and logging middleware is in the middleware.go

If you need to add a new api route, it is done in the routes.go

## services
All the gRPC services protocol is defined in the protocol.proto file, and then the individual clients are created in their own services package
with a main. that handles the starting and registering of the service client, and then in a subfolder the service with the same name,
So again the JournalService would be in the services/journalService package with a main.go and a journal package which has the journal.go client implementation

So adding a new service would first require you to change the protocol.proto file and then generate the needed go code.
```
protoc --go_out=./packages/protocol --go-grpc_out=packages/protocol protocol.proto
```

This creates the two required protocol files for the client and server interfaces of the different services.

## packages
The packages folder contains everything thats not its own standalone program, so here we have he mssql package for the MSSQL data communication.
We have the mongo for the MongoDB datta communication.
We have some models
And last but no least we have the auto generated gRPC helper code defined by the protocol.proto file

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