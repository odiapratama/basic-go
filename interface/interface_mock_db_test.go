package main

import "testing"

func TestGetUserById(t *testing.T) {
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
		t.Errorf("error:  %v", err)
	}

	var actualId = 1
	if user.ID != actualId {
		t.Errorf("expected %v but got %v", actualId, user.ID)
	}
}

func TestAddUser(t *testing.T) {
	db := MockDataStore{
		Data: map[int]User{
			1: {ID: 1, Name: "User 1"},
			2: {ID: 2, Name: "User 2"},
		},
	}

	service := Service{
		Repository: db,
	}

	err := service.AddUser(User{ID: 3, Name: "User 3"})
	if err != nil {
		t.Errorf("error:  %v", err)
	}

	user, err := service.GetUserByID(3)
	if err != nil {
		t.Errorf("error:  %v", err)
	}

	var actualId = 3
	if user.ID != actualId {
		t.Errorf("expected %v but got %v", actualId, user.ID)
	}
}
