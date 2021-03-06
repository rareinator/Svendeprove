package authentication

import (
	context "context"
	"fmt"
)

type AuthenticationServer struct {
	UnimplementedAuthenticationServiceServer
	ListenAddress string
}

func (a *AuthenticationServer) GetHealth(ctx context.Context, e *AEmpty) (*AHealth, error) {
	return &AHealth{Message: fmt.Sprintf("Authentication service is up and running on: %v 🚀", a.ListenAddress)}, nil
}
