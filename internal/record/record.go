package record

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingRecord = errors.New(("failed to fetch record by id"))
	ErrNotImplemented = errors.New(("not implemented"))
)

// Record - a representation of the record
// structure for our service
type Record struct {
	ID string
	DATE string
	AMOUNT int
	CATEGORY string
	DESCRIPTION string
	EXPENSE_TYPE string
}

// Store - this interface defines all of the methods
// that our service needs in order to operate
type Store interface {
	GetRecord(context.Context, string)(Record, error)
}

// Service - is the struct 
// on which all our logic will be built on top of
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new service
func NewService(store Store) *Service{
	return &Service{
		Store: store,
	}
}

func (s *Service) GetRecord(ctx context.Context, id string)(Record, error){
	fmt.Println("retrieving a record")

	record, err := s.Store.GetRecord(ctx, id)
	if err !=nil{
		fmt.Println(err)
		return Record{}, err
	}
	return record, nil
}

func (s *Service) UpdateRecord(ctx context.Context, record Record) error {
	return ErrNotImplemented
}
func (s *Service) DeleteRecord(ctx context.Context, id string) error{
	return ErrNotImplemented
}

func (s *Service) CreateRecord(ctx context.Context, record Record)(Record, error) {
	return Record{}, ErrNotImplemented
}