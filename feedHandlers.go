package main 

import ("fmt"
"context"
"os"
"log"
"syscall"
"os/signal"
"database/sql"
"github.com/google/uuid"
"github.com/samiiac/bootdev-gator/internal/database"
"time"
"strconv"
)

func handlerGetFeed(s *state,cmd Command) error{
  if len(cmd.args) <= 0 {
    return fmt.Errorf("Time duration required.")
  }
  dur,err := time.ParseDuration(cmd.args[0])
  if err != nil {
      return err
    }
  
    //initial fetch
  if err:= scrapeFeeds(s);err!=nil {
        return err
  }

  sigChan := make(chan os.Signal,1)
  signal.Notify(sigChan,os.Interrupt,syscall.SIGTERM)
  defer signal.Stop(sigChan)

  ticker := time.NewTicker(dur)
  defer ticker.Stop()

  fmt.Println("Scraper started. Press CTRL+C to exit willingly.")

  //infinite loop
   for {
    select {
    case  <-sigChan:
      fmt.Println("Shutting down scraper!")
     return nil
    case <-ticker.C:
      if err:= scrapeFeeds(s);err!=nil {
        log.Printf("Error while scraping feed : %v ",err)
      }
    }
   }
    
  return nil
}

func scrapeFeeds(s *state)error {
  ctx := context.Background()
  feedToFetch,err := s.db.GetNextFeedToFetch(ctx)
  if err!=nil {
    return err
  }
   err = s.db.MarkedFeedFetched(ctx,database.MarkedFeedFetchedParams{
    LastFetchedAt:sql.NullTime{
      Time:time.Now(),
      Valid:true,
    },
    UpdatedAt:time.Now(),
    ID:feedToFetch.ID,
  })
  if err!=nil {
    return err
  }

  feed,err := fetchFeed(context.Background(),feedToFetch.Url)
 
  
  for _,i := range feed.Channel.Item{
    pubTime, err := time.Parse(time.RFC1123Z, i.PubDate)
  if err != nil {
    pubTime, err = time.Parse(time.RFC1123, i.PubDate)
  }

  var pubDateNull sql.NullTime
  if err == nil {
    pubDateNull = sql.NullTime{
      Time:  pubTime,
      Valid: true,
    }
  }
   _,err = s.db.CreatePost(ctx,database.CreatePostParams{
    ID:uuid.New(),
    Title:i.Title,
    CreatedAt:time.Now(),
    UpdatedAt:time.Now(),
    Url:i.Link,
    Description:sql.NullString{
      String:i.Description,
      Valid:i.Description != "",
    },
    PublishedAt:pubDateNull,
    FeedID:feedToFetch.ID,
   })

   if err != nil {
    log.Printf("Failed to create post : %v \n",err)
   }
   
  }
  return nil;
}

func handlerAddFeed(s *state,cmd Command,user database.User) error{
  if len(cmd.args) <=1 {
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

   fmt.Println("Feed has been created.\n")
   fmt.Printf("Name: %v \n Url: %v\n",newFeed.Name,newFeed.Url)
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

func handlerPostFromFollowing(s *state,cmd Command,user database.User) error{
var limit int32
 if len(cmd.args) > 0{
  l,err := strconv.Atoi(cmd.args[0])
  if err != nil {
  return err
 }
 limit = int32(l)
 }else{
  limit = 2
 }

 posts,err := s.db.GetPostForUser(context.Background(),database.GetPostForUserParams{
  UserID:user.ID,
  Limit:limit,
 })
 if err != nil {
  return err
 }
 fmt.Println("---POSTS---")
 for _,p := range posts{
  fmt.Printf("-Title : %v \n-URL : %v \n-Description : %v\n\n",p.Title,p.Url,p.Description)
 }
 return nil
}