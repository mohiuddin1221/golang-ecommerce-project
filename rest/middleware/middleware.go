package middleware

import "ecommerce/config"

type Middlewares struct {
	cnf *config.Config
}

func Newmiddleare(cnf *config.Config) *Middlewares {
	return &Middlewares{
		cnf: cnf,
	}
}
