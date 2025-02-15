package transport

import (
	"context"
	"net/http"
)

type Transport interface {
	Server(
		endpoint Endpoint,
		decode func(ctx context.Context, r *http.Request) (interface{}, error),
		encode func(ctx context.Context, w http.ResponseWriter, resp interface{}) error,
		encondeError func(ctx context.Context, err error, w http.ResponseWriter),

	)
}

type Endpoint func(ctx context.Context, request interface{}) (interface{}, error)

type transport struct {
	w   http.ResponseWriter
	r   *http.Request
	ctx context.Context
}

func New(w http.ResponseWriter, r *http.Request, ctx context.Context) Transport {
	return &transport{
		w: w,
		r: r,
		ctx:ctx,
	}
}

func (t *transport) Server(
	endpoint Endpoint,
	decode func(ctx context.Context, r *http.Request) (interface{}, error),
	encode func(ctx context.Context, w http.ResponseWriter, resp interface{}) error,
	encondeError func(ctx context.Context, err error, w http.ResponseWriter),

) {
	data, err := decode(t.ctx, t.r)
	if err != nil {
		encondeError(t.ctx, err, t.w)
		return
	}

	res, err := endpoint(t.ctx, data)
	if err != nil {
		encondeError(t.ctx, err, t.w)
		return
	}

	if err := encode(t.ctx, t.w, res); err != nil {
		encondeError(t.ctx, err, t.w)
		return
	}

}
