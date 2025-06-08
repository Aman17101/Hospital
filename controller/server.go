package controller

import (
	"fmt"

	"github.com/Aman17101/Hospital/store"
)

type Server struct {
	Postgressdb store.StoreOperation
}

func (s *Server) NewServer(pgstore store.Postgress) {
	s.Postgressdb=&pgstore
	s.Postgressdb.NewStore()
	fmt.Printf("server is %v \n", s)
}
type ServerOperations interface{
	NewServer(pgstore store.Postgress)
}