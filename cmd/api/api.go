package api

type APIServer struct {
	addr string
	db	 *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	&APISever{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error {
	
}
