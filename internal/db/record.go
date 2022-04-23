package db

import (
	"context"
	"database/sql"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/yuchida-tamu/money-tracker-server/internal/record"
)

type RecordRow struct {
	ID string
	DATE_CREATED sql.NullString
	AMOUNT sql.NullInt32
	CATEGORY sql.NullString
	RECORD_DESCRIPTION sql.NullString
	EXPENSE_TYPE sql.NullString
}

func convertRecordRowToRecord(r RecordRow) record.Record{
	return record.Record{
		ID: r.ID,
		DATE_CREATED: r.DATE_CREATED.String,
		AMOUNT: int(r.AMOUNT.Int32),
		CATEGORY: r.CATEGORY.String,
		RECORD_DESCRIPTION: r.RECORD_DESCRIPTION.String,
		EXPENSE_TYPE: r.RECORD_DESCRIPTION.String,
	}
}

func (d *Database) GetRecord(ctx context.Context, uuid string)(record.Record, error){
	var recordRow RecordRow
	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, date_created, amount, category, record_description, expense_type
		FROM records
		WHERE id = $1`,
		uuid,
	)
	err := row.Scan(&recordRow.ID, &recordRow.DATE_CREATED, &recordRow.AMOUNT, &recordRow.CATEGORY, &recordRow.RECORD_DESCRIPTION, &recordRow.EXPENSE_TYPE)
	if err != nil {
		return record.Record{}, fmt.Errorf("error fetching the record by uuid: %w", err)
	}

	return convertRecordRowToRecord(recordRow), nil
}

func (d *Database) PostRecord(ctx context.Context, rcd record.Record)(record.Record, error){
	rcd.ID = uuid.NewV4().String()
	postRow := RecordRow{
		ID: rcd.ID,
		DATE_CREATED: sql.NullString{String: rcd.DATE_CREATED, Valid: true},
		AMOUNT: sql.NullInt32{Int32: int32(rcd.AMOUNT), Valid: true},
		CATEGORY: sql.NullString{String: rcd.CATEGORY, Valid: true},
		RECORD_DESCRIPTION: sql.NullString{String: rcd.RECORD_DESCRIPTION,Valid: true},
		EXPENSE_TYPE: sql.NullString{String: rcd.EXPENSE_TYPE, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO records
		(id, date_created, amount, category, record_description, expense_type)
		VALUES
		(:id, :date_created, :amount, :category, :record_description, :expense_type)`,
		postRow,
	)

	if err != nil {
		return record.Record{}, fmt.Errorf("failed to insert record: %w", err)
	}

	if err := rows.Close(); err != nil {
		return record.Record{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return rcd, nil
}

func (d *Database) DeleteRecord(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM records where id = $1`,
	)
	if err != nil {
		return fmt.Errorf("failed to delete record from database: %w", err)
	}
	return nil
}

func (d *Database) UpdateRecord (ctx context.Context, id string, rcd record.Record) (record.Record, error){
	rcdRow := RecordRow{
		ID: id,
		DATE_CREATED: sql.NullString{String: rcd.DATE_CREATED, Valid: true},
		AMOUNT: sql.NullInt32{Int32: int32(rcd.AMOUNT), Valid: true},
		CATEGORY: sql.NullString{String: rcd.CATEGORY, Valid: true},
		RECORD_DESCRIPTION: sql.NullString{String: rcd.RECORD_DESCRIPTION,Valid: true},
		EXPENSE_TYPE: sql.NullString{String: rcd.EXPENSE_TYPE, Valid: true},
	}
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE records SET
		date_created = :date_created,
		amount = :amount,
		category = :category,
		record_description = :description,
		expense_type = :expense_type 
		WHERE id = :id`,
		rcdRow,
	)
	if err != nil {
		return record.Record{}, fmt.Errorf("failed to update record: %w", err)
	}
	if err := rows.Close(); err != nil {
		return record.Record{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return convertRecordRowToRecord(rcdRow), nil
}