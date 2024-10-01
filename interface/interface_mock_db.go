package main

import "fmt"

func mockDB() {
	fmt.Println("=== Interface Mock DB ===")

	db := MockDataStore{
		Data: map[int]User{
			1: {ID: 1, Name: "User 1"},
			2: {ID: 2, Name: "User 2"},
		},
	}

	service := Service{
		Repository: db,
	}

	user, err := service.GetUserByID(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}

	err = service.AddUser(User{ID: 3, Name: "User 3"})
	if err != nil {
		fmt.Println(err)
	}

	user, err = service.GetUserByID(3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}
}

type User struct {
	ID   int
	Name string
}

type MockDataStore struct {
	Data map[int]User
}

type Service struct {
	Repository UserRepository
}

type UserRepository interface {
	GetUserByID(id int) (User, error)
	AddUser(user User) error
}

func (s Service) GetUserByID(id int) (User, error) {
	return s.Repository.GetUserByID(id)
}

func (s Service) AddUser(user User) error {
	return s.Repository.AddUser(user)
}

func (db MockDataStore) GetUserByID(id int) (User, error) {
	user, ok := db.Data[id]
	if !ok {
		return User{}, fmt.Errorf("User with ID %d not found", id)
	}
	return user, nil
}

func (db MockDataStore) AddUser(user User) error {
	_, ok := db.Data[user.ID]
	if ok {
		return fmt.Errorf("User with ID %d already exists", user.ID)
	}
	db.Data[user.ID] = user
	return nil
}
