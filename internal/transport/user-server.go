// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"

	"github.com/sah4ez/dapr-example/interfaces"
	"github.com/sah4ez/dapr-example/interfaces/types"
)

type serverUser struct {
	svc         interfaces.User
	getNameByID UserGetNameByID
}

type MiddlewareSetUser interface {
	Wrap(m MiddlewareUser)
	WrapGetNameByID(m MiddlewareUserGetNameByID)

	WithMetrics()
	WithLog()
}

func newServerUser(svc interfaces.User) *serverUser {
	return &serverUser{
		getNameByID: svc.GetNameByID,
		svc:         svc,
	}
}

func (srv *serverUser) Wrap(m MiddlewareUser) {
	srv.svc = m(srv.svc)
	srv.getNameByID = srv.svc.GetNameByID
}

func (srv *serverUser) GetNameByID(ctx context.Context, id types.ID) (user types.User, err error) {
	return srv.getNameByID(ctx, id)
}

func (srv *serverUser) WrapGetNameByID(m MiddlewareUserGetNameByID) {
	srv.getNameByID = m(srv.getNameByID)
}

func (srv *serverUser) WithMetrics() {
	srv.Wrap(metricsMiddlewareUser)
}

func (srv *serverUser) WithLog() {
	srv.Wrap(loggerMiddlewareUser())
}
