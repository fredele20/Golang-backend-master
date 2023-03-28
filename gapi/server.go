package gapi

import (
	"fmt"

	db "github.com/fredele20/Golang-backend-master/db/sqlc"
	"github.com/fredele20/Golang-backend-master/pb"
	"github.com/fredele20/Golang-backend-master/token"
	"github.com/fredele20/Golang-backend-master/util"
)

// Server servs gRPC requests for out banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot crate token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
