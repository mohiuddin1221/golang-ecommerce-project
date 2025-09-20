package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func (mngr *Manager) Use(middlewares ...Middleware) *Manager {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
	return mngr
}

func (mngr *Manager) With(next http.Handler) http.Handler {
	for i := len(mngr.globalMiddlewares) - 1; i >= 0; i-- {
		next = mngr.globalMiddlewares[i](next)
	}
	return next
}
