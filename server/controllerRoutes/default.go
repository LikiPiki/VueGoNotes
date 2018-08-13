package controllerRoutes

import "io"

type GetController interface {
	GetById(id string) (interface{}, error)
}

type PostController interface {
	Create() error
}

type PutController interface {
	Update(id string) error
	Decode(io.ReadCloser) (PutController, error)
}

type DeleteController interface {
	Delete(id string) error
}
