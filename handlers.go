package main
import ("fmt"
"context"
"os"
"github.com/google/uuid"
"github.com/samiiac/bootdev-gator/internal/database"
"database/sql"
"time"
"errors"
)

func handlerLogin(s *state,cmd Command) error {
	if len(cmd.args) <= 0 {
       return fmt.Errorf("Username required to login.")
  }
	
	username := cmd.args[0]
	ctx := context.Background()
    _,err := s.db.GetUserByName(ctx,username)

   if err != nil && errors.Is(err,sql.ErrNoRows){
	os.Exit(1)
   } 
   
	err = s.config.SetUser(username)
	if err != nil {
     return err
	}
    fmt.Println("Username has been set.")
	return nil
}


func handlerRegister(s *state,cmd Command) error{
   if len(cmd.args) <= 0 {
       return fmt.Errorf("Username required to register.")
  }
  ctx := context.Background()
  _,err := s.db.GetUserByName(ctx,cmd.args[0])

   if err == nil {
	os.Exit(1)
   }

   if !errors.Is(err,sql.ErrNoRows) {
     return err
   }

   ctx = context.Background()
  newUser,err := s.db.CreateUser(ctx,database.CreateUserParams{
	ID:uuid.New(),
	CreatedAt:time.Now(),
	UpdatedAt:time.Now(),
	Name:cmd.args[0],
  })

  if err != nil {
	return err
  }

 if  err := s.config.SetUser(cmd.args[0]);err !=nil{
	return err
 }
  fmt.Println("User with name " + newUser.Name +" was created.")
  return nil
}

func handlerReset(s *state,cmd Command) error{
	ctx:= context.Background()
	err := s.db.DeleteAllUsers(ctx)

	if err != nil {
		return err
	}
    fmt.Println("Table Users Cleared.")
	return nil
}
