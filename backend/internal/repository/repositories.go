package repository

import "github.com/Adedunmol/glimpse/internal/server"

type Repositories struct{}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}
