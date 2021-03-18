package dbutil

import (
	"errors"
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func TestTXHandlerBeginError(t *testing.T) {

	mockDB, mock, mockErr := sqlmock.New()
	if mockErr != nil {
		t.Fatalf("unexpected error: %v", mockErr)
	}
	defer mockDB.Close()
	mock.ExpectBegin().WillReturnError(errors.New("could not begin transaction"))

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	err := TXHandler(sqlxDB, func(tx *sqlx.Tx) error {
		tx.DriverName()
		return nil
	})

	if err == nil {
		t.Fatalf("Expected error but got nil")
	}

	if got, want := err.Error(), "start transaction failed: could not begin transaction"; want != got {
		t.Fatalf("Want %v but got %v", want, got)
	}
}

func TestTXHandlerRollbackSuccessWhenError(t *testing.T) {
	mockDB, mock, mockErr := sqlmock.New()
	if mockErr != nil {
		t.Fatalf("unexpected error: %v", mockErr)
	}
	defer mockDB.Close()
	mock.ExpectBegin()
	mock.ExpectRollback()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	err := TXHandler(sqlxDB, func(tx *sqlx.Tx) error {
		tx.DriverName()
		return errors.New("some error")
	})

	if err == nil {
		t.Fatalf("Expected error but got nil")
	}

	if got, want := err.Error(), "transaction: operation failed: some error"; want != got {
		t.Fatalf("Want %v but got %v", want, got)
	}
}

func TestTXHandlerRollbackSuccessWhenPanic(t *testing.T) {
	mockDB, mock, mockErr := sqlmock.New()
	if mockErr != nil {
		t.Fatalf("unexpected error: %v", mockErr)
	}
	defer mockDB.Close()
	mock.ExpectBegin()
	mock.ExpectRollback()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	err := TXHandler(sqlxDB, func(tx *sqlx.Tx) error {
		tx.DriverName()
		panic("panic")
	})

	if got, want := err.Error(), "recovered: panic"; want != got {
		t.Fatalf("Want %v but got %v", want, got)
	}
}

func TestTXHandlerRollbackError(t *testing.T) {
	mockDB, mock, mockErr := sqlmock.New()
	if mockErr != nil {
		t.Fatalf("unexpected error: %v", mockErr)
	}
	defer mockDB.Close()
	mock.ExpectBegin()
	mock.ExpectRollback().WillReturnError(errors.New("could not rollback transaction"))

	var exitCode int
	fakeExit := func(int) {
		exitCode = 1
		return
	}

	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	TXHandler(sqlxDB, func(tx *sqlx.Tx) error {
		tx.DriverName()
		return errors.New("some error")
	})

	if got, want := exitCode, 1; want != got {
		t.Fatalf("Want %v but got %v", want, got)
	}
}
