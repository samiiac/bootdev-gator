package main
import ("fmt")

type Command struct {
	name   string
	args   []string
}

type Commands struct {
	cmdRegistry  map[string]func(*state,Command)error
}

func newCommands() *Commands {
  c := &Commands{
    cmdRegistry : make(map[string]func(*state,Command)error),
  }
  err := c.register("login",handlerLogin)

    if err != nil {
	  fmt.Println(err)
    return nil
   }

   err = c.register("register",handlerRegister)
   if err != nil {
	  fmt.Println(err)
    return nil
   }

    err = c.register("reset",handlerReset)
   if err != nil {
	  fmt.Println(err)
    return nil
   }

   err = c.register("users",handlerGetUsers)
   if err != nil {
	  fmt.Println(err)
    return nil
   }

   err = c.register("agg",handlerGetFeed)
   if err != nil {
	  fmt.Println(err)
    return nil
   }

   err = c.register("addfeed",middlewareLoggedIn(handlerAddFeed))
   if err != nil {
	  fmt.Println(err)
    return nil
   }

   err = c.register("feeds",handlerDisplayFeed)
   if err != nil {
	  fmt.Println(err)
    return nil
   }

   err = c.register("follow",middlewareLoggedIn(handlerFollowFeed))
   if err != nil {
	  fmt.Println(err)
    return nil
   }

   err = c.register("following",middlewareLoggedIn(handlerGetFollowingFeeds))
   if err != nil {
	  fmt.Println(err)
    return nil
   }

    err = c.register("unfollow",middlewareLoggedIn(handlerUnfollowFeed))
   if err != nil {
	  fmt.Println(err)
    return nil
   }

    err = c.register("browse",middlewareLoggedIn(handlerPostFromFollowing))
   if err != nil {
	  fmt.Println(err)
    return nil
   }

   return c
}

func (c *Commands) run(s *state,cmd Command) error {
  handler,ok:= c.cmdRegistry[cmd.name]
  if !ok {
	return fmt.Errorf("No such command.")
  }
  err := handler(s,cmd)
  if err != nil {
	return err
  }
  return nil
}

func (c *Commands) register(name string,handler func(*state,Command)error) error {

  _,ok := c.cmdRegistry[name]
  if ok {
	return fmt.Errorf("Command already exists.")
  }
  c.cmdRegistry[name] = handler

  return nil
}

