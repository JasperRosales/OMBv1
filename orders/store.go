package main

import "context"

type store struct{
	// database here
}

func NewStore() *store{	
	return &store{}
}

func (s *store) CreateOrder(context.Context) error{
	return nil
}