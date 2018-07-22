package main

// Model default interface for db models
type Model interface {
	Update() bool
	GetAll() (interface{}, error)
	Create() bool
}
