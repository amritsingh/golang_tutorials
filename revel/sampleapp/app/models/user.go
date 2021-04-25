package models

import (
	"time"
	"sampleapp/app"
)

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// This function looks for the email in the table, if entry exits update the name.
func (u User) UpdateOrCreate(name string, email string) *User {
	user := User{}
	app.DB.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		user = User{Name: name, Email: email, CreatedAt: time.Now(), UpdatedAt: time.Now()}
		app.DB.Create(&user)
	} else {
		user.Name = name
		user.UpdatedAt = time.Now()
		app.DB.Save(&user)
	}
	return (&user)
}

// Get an entry by ID
func (u User) GetByID(id uint64) *User {
	user := User{}
	app.DB.Where("id = ?", id).First(&user)
	return (&user)
}

// Find an entry by ID and delete it
func (u User) DeleteUser(id uint64) bool{
    user := User{}.GetByID(id)
	status := false
	if user.ID != 0 {
		app.DB.Delete(user)
		status = true
	}
	return status
}

