package dbrepo

import (
	"errors"
	"time"

	"github.com/akshay/bookings/internal/models"
)

// Inserts reservation to the db
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	if res.RoomID == 2 {
		return 0, errors.New("some error")
	}
	return 1, nil
}

// Inserts room restriction to the db
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	if r.RoomID == 1000 {
		return errors.New("some error")
	}
	return nil
}

// rerturns boolean value depeding if availabilty exists or not
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	if start.Year() == 2060 {
		return false, errors.New("database error")
	}
	return false, nil
}

func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	if start.Year() == 2060 {
		return nil, errors.New("database error")
	}
	var rooms []models.Room
	return rooms, nil
}
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room

	if id > 2 {
		return room, errors.New("some error")
	}
	return room, nil
}
