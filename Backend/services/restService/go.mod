module github.com/rareinator/Svendeprove/Backend/services/restService

go 1.15

require (
	github.com/gorilla/mux v1.8.0
	github.com/rareinator/Svendeprove/Backend/services/journalService/journal v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.36.0
)

replace github.com/rareinator/Svendeprove/Backend/services/journalService => ../journalService

replace github.com/rareinator/Svendeprove/Backend/services/journalService/journal => ../journalService/journal
