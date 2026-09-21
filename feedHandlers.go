package main 

import ("fmt"
"context"
"os"
"github.com/google/uuid"
"github.com/samiiac/bootdev-gator/internal/database"
"time"
)

func handlerGetFeed(s *state,cmd Command) error{
  feed,err := fetchFeed(context.Background(),"https://www.wagslane.dev/index.xml")
  if err!= nil{
    return err
  }
  fmt.Println(feed)
  return nil

}

func handlerAddFeed(s *state,cmd Command,user database.User) error{
  if len(cmd.args) <= 1 {
       os.Exit(1)
  }
   urlName := cmd.args[0]
   url := cmd.args[1]

   ctx := context.Background()
   
  
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

   _,err = s.db.CreateFollowFeed(ctx,database.CreateFollowFeedParams{
    ID:uuid.New(),
    UserID:user.ID,
    FeedID:newFeed.ID,
    CreatedAt:time.Now(),
    UpdatedAt:time.Now(),
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

func handlerFollowFeed(s *state,cmd Command,user database.User)error {
  if len(cmd.args) <=0 {
    return fmt.Errorf("To be followed url is required.")
  }
  ctx := context.Background()
  feed,err := s.db.GetFeedByUrl(ctx,cmd.args[0])
   if err != nil {
    return err
   }
  newFollowFeed,err := s.db.CreateFollowFeed(ctx,database.CreateFollowFeedParams{
    ID:uuid.New(),
    UserID:user.ID,
    FeedID:feed.ID,
    CreatedAt:time.Now(),
    UpdatedAt:time.Now(),
  })
  if err != nil {
    return err
   }
  fmt.Printf("%v started follwing feed : %v\n",newFollowFeed.UserName,newFollowFeed.FeedName)
  return nil

}


func handlerGetFollowingFeeds(s *state,cmd Command,user database.User)error {
  if s.config.UserName == "" {
    return fmt.Errorf("You need to login first")
  }
  ctx := context.Background()

  following,err := s.db.GetFollowFeedsForUser(ctx,user.ID)
  if err != nil {
    return err
   }
   fmt.Println("Following: ")
   for _,f := range following{
    fmt.Printf("- %v\n",f.FeedName)
   }
   return nil;
}

func handlerUnfollowFeed(s *state,cmd Command,user database.User) error{
  if len(cmd.args) <=0 {
    return fmt.Errorf("Url is required")
  }
  ctx := context.Background()
  feed,err := s.db.GetFeedByUrl(ctx,cmd.args[0])
  if err != nil {
    return err
  }

  err = s.db.UnfollowFeed(ctx,database.UnfollowFeedParams{
    UserID:user.ID,
    FeedID:feed.ID,
  })

  if err != nil {
    return err
  }
  fmt.Printf("Unfollowed %v \n",feed.Name)
  return nil


}