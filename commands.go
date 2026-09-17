package main
import ("fmt")

type Command struct {
	name   string
	args   []string
}

type Commands struct {
	cmdRegistry  map[string]func(*state,Command)error
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
  if c.cmdRegistry == nil {
        c.cmdRegistry = make(map[string]func(*state, Command) error)
  }

  _,ok := c.cmdRegistry[name]
  if ok {
	return fmt.Errorf("Command already exists.")
  }
  c.cmdRegistry[name] = handler
  fmt.Println("Command has been added.")
  return nil
}

func handlerLogin(s *state,cmd Command) error {
	if len(cmd.args) <= 0 {
       return fmt.Errorf("Username required to login.")
  }
	
	username := cmd.args[0]
	err := s.config.SetUser(username)
	if err != nil {
     return err
	}
    fmt.Println("Username has been set.")
	return nil
}