package main

import (

"github.com/samiiac/bootdev-gator/internal/database"

"context"
)

func middlewareLoggedIn(handler func(s *state, cmd Command, user database.User) error) func(*state, Command) error{
	return func(s *state,cmd Command)error { 
		user,err := s.db.GetUserByName(context.Background(),s.config.UserName)
		if err != nil {
		return err
		}
		return handler(s,cmd,user)
	}
}