package gapi

import (
	"fmt"

	db "github.com/fredele20/Golang-backend-master/db/sqlc"
	"github.com/fredele20/Golang-backend-master/pb"
	"github.com/fredele20/Golang-backend-master/token"
	"github.com/fredele20/Golang-backend-master/util"
	"github.com/fredele20/Golang-backend-master/worker"
)

// Server servs gRPC requests for out banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot crate token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
