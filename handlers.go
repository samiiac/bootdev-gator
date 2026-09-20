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

func handlerGetUsers(s *state,cmd Command) error{
  users,err := s.db.GetAllUsers(context.Background())

  if err != nil {
   return fmt.Errorf("Error fetching users : %v",err)
  }
  
  if len(users) == 0 {
    return fmt.Errorf("No users registered.")
  }

  for _,user := range users {
    if user.Name == s.config.UserName {
      fmt.Printf("* %v (current)\n",user.Name)
    }else{
    fmt.Printf("* %v\n",user.Name)
    }
  }
  return nil
}


func handlerGetFeed(s *state,cmd Command) error{
  feed,err := fetchFeed(context.Background(),"https://www.wagslane.dev/index.xml")
  if err!= nil{
    return err
  }
  fmt.Println(feed)
  return nil

}

func handlerAddFeed(s *state,cmd Command) error{
  if len(cmd.args) <= 1 {
       os.Exit(1)
  }
   urlName := cmd.args[0]
   url := cmd.args[1]

   ctx := context.Background()
   user,err := s.db.GetUserByName(ctx,s.config.UserName)
   if err != nil {
    return err
   }
  
   newFeed,err := s.db.CreateFeed(ctx,database.CreateFeedParams{
    ID:uuid.New(),
    Name:urlName,
    Url:url,
    CreatedAt:time.Now(),
    UpdatedAt:time.Now(),
    UserID:user.ID,
   })

   if err != nil {
    return err
   }

   fmt.Println("Feed has been created.")
   fmt.Printf("Name: %v \n Url: %v",newFeed.Name,newFeed.Url)
   return nil;

  
}

func handlerDisplayFeed(s *state,cmd Command)error{
  ctx:= context.Background()
  feeds,err := s.db.GetAllFeed(ctx)
  if err != nil {
    return err
  }
   
  for _,f := range feeds {
    user,err := s.db.GetUserById(ctx,f.UserID)
    if err != nil {
    return err
   }
    fmt.Printf("Name : %v\nUrl: %v\nCreated By : %v\n",f.Name,f.Url,user.Name)
  }
  return nil
}
