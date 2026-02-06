package models

import (
	"errors"
	"time"

	"github.com/warlock1729/first-go-project/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (event *Event) Save() error {
	query := `
	INSERT INTO events(name,description,location,datetime,user_id) values (?,?,?,?,? )`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	event.ID = id
	defer stmt.Close()
	return nil
}

func (event *Event) RegisterUser(userID int64) error {
	query := `
	  	Insert into registrations (user_id,event_id)
		values (?,?)
	  `
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userID, event.ID)
	defer stmt.Close()
	if err != nil {
		return err
	}
	return err
}

func (event *Event) CancelRegistration(userID int64) error {
	query := `
	 
	Delete from registrations where event_id =? and user_id =?
	  `
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(event.ID, userID)
	defer stmt.Close()
	if err != nil {
		return err
	}

	rowsAffected , _:=result.RowsAffected()
	if rowsAffected==0{
		return errors.New("No registration found to cancel !")
	}
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `
	Select * from events
	`
	rows, err := db.DB.Query(query)
	var events []Event

	if err != nil {
		return events, err
	}

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	defer rows.Close()
	return events, nil
}

func GetEventByID(ID int64) (*Event, error) {
	query := `
	select * from events where id = ?
	`
	row := db.DB.QueryRow(query, ID)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event *Event) Update() error {
	query := `
	Update events
	 set name=?, description=?,location=?,datetime=? 
	 where id=?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	if err != nil {
		return err
	}
	return nil
}

func (event Event) DeleteEvent() error {
	query := `
	delete 
	from events
	where id= ? 
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(event.ID)
	defer stmt.Close()
	return err
}
