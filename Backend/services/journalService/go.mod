module github.com/rareinator/Svendeprove/Backend/services/journalService

go 1.15

require (
	github.com/denisenkom/go-mssqldb v0.9.0 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/rareinator/Svendeprove/Backend/services/journalService/journal v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.36.0
	gorm.io/gorm v1.21.2 // indirect
)

replace github.com/rareinator/Svendeprove/Backend/services/journalService => ./

replace github.com/rareinator/Svendeprove/Backend/services/journalService/journal => ./journal
