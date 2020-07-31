package server

type problem struct {
	Description string
	Solutions   []string
}

func (srv *server) Problem() problem {
	p := srv.ctg.Problem()
	q := problem{
		Description: p.Description,
	}
	for _, solution := range p.Solutions {
		q.Solutions = append(q.Solutions, solution.Description)
	}
	return q
}
