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
	DATE_CREATED string
	AMOUNT int
	CATEGORY string
	RECORD_DESCRIPTION string
	EXPENSE_TYPE string
}

// Store - this interface defines all of the methods
// that our service needs in order to operate
type Store interface {
	GetRecord(context.Context, string)(Record, error)
	PostRecord(context.Context, Record)(Record, error)
	DeleteRecord(context.Context, string)(error)
	UpdateRecord(context.Context, string, Record)(Record, error)
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

func (s *Service) UpdateRecord(ctx context.Context,ID string, updatedRecord Record) (Record, error) {
	rcd, err := s.Store.UpdateRecord(ctx, ID, updatedRecord)
	if err != nil {
		fmt.Println("error updating record")
		return Record{}, err
	}
	return rcd, nil
}
func (s *Service) DeleteRecord(ctx context.Context, id string) error{
	return s.Store.DeleteRecord(ctx, id)
}

func (s *Service) PostRecord(ctx context.Context, record Record)(Record, error) {
	insertedRecord, err := s.Store.PostRecord(ctx, record)
	if err != nil {
		return Record{}, err
	}
	return insertedRecord, nil
}