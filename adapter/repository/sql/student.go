package sql

import (
	"context"
	"database/sql"
	"my-github/clean-code-microservice-golang/domain/model"
	"my-github/clean-code-microservice-golang/usecase/student/repository"

	_ "github.com/go-sql-driver/mysql"
)

const (
	writeStudent = `INSERT INTO mst_student (student_uuid, name, gender, address, date_of_birth, age) VALUES (?,?,?,?,?,?)`
)

type DBConn struct {
	conn *sql.DB
}

func NewSQLRepository(db *sql.DB) repository.StudentSQLRepository {
	return &DBConn{conn: db}
}

func (db *DBConn) WriteStudent(ctx context.Context, req *model.Student) error {
	var err error

	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, writeStudent, req.StudentID, req.Name, req.Gender, req.Address, req.DateOfBirth, req.Age)
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return errRollback
		}
		return err
	}

	return tx.Commit()
}
