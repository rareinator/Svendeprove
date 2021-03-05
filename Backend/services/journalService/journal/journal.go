package journal

import (
	"context"
	"log"
)

type JournalServer struct {
	UnimplementedJournalServiceServer
}

func (j *JournalServer) GetJournal(ctx context.Context, journal *Journal) (*Journal, error) {
	log.Printf("Received journal from client: %s", journal.Intro)
	return &Journal{Intro: "hello from the server!"}, nil
}
