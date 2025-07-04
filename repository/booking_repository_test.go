package repository

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/rizkideveloper/ticket-booking/helper"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func TestMain(m *testing.M) {
	var err error
	//buka koneksi ke database 
	db, err = sql.Open("mysql","root@tcp(127.0.0.1:3306)/ticket-booking")
	helper.PanicIfError(err)

	insertEvent(1,100,"Event 1")
	insertSeat(2,1,"A1","AVAILABLE")
	insertUser(3, "rizki")

	exitCode := m.Run()
	db.Close()
	os.Exit(exitCode)
}

//buat seeder event
func insertEvent(id, quota int, name string)  {
	SQL := "INSERT INTO events (id, name, quota) VALUES(?,?,?) ON DUPLICATE KEY UPDATE quota = VALUES(quota)"

	_,err := db.Exec(SQL,id,name,quota)
	helper.PanicIfError(err)
}

//buat seeder seat
func insertSeat(id, event_id int, seat_number, status string)  {
	SQL := "INSERT INTO seats (id, event_id,seat_number, status) VALUES(?,?,?,?) ON DUPLICATE KEY UPDATE status = VALUES(status)"

	_,err := db.Exec(SQL,id,event_id,seat_number,status)
	helper.PanicIfError(err)
}

//buat seeder user
func insertUser(id int, name string)  {
	SQL := "INSERT INTO users (id, name) VALUES(?,?) ON DUPLICATE KEY UPDATE name = VALUES(name)"

	_,err := db.Exec(SQL,id,name)
	helper.PanicIfError(err)
}

func TestGetSeatStatus(t *testing.T)  {
	//initiate
	repo := NewBookingRepository()
	ctx := context.Background()
	tx,err := db.BeginTx(ctx , nil)
	assert.NoError(t, err)

	//anonymous func
	defer func ()  {
		r := recover()
		if r != nil {
			tx.Rollback()
			t.Errorf("panic %v", r)
		}else{
			tx.Commit()
		}
	}()

	status := repo.GetSeatStatus(ctx, tx, 2)
	assert.Equal(t,"AVAILABLE", status)
} 