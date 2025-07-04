package repository

import (
	"context"
	"database/sql"

	"github.com/rizkideveloper/ticket-booking/helper"
)

//yang di eksport(public) yang ini
type BookingRepository interface {
	GetSeatStatus(ctx context.Context, tx *sql.Tx, seatId int) string
}

type bookingRepository struct {

}

//constructor
func NewBookingRepository() BookingRepository {
	return &bookingRepository{}
}

func (repo *bookingRepository) GetSeatStatus(ctx context.Context, tx *sql.Tx, seatId int) string {
	var status string
	SQL := "SELECT status FROM seats WHERE id = ?"
	err := tx.QueryRowContext(ctx,SQL, seatId).Scan(&status)
	helper.PanicIfError(err)
	
	return status 
}