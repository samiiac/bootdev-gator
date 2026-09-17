package main 
import ("github.com/samiiac/bootdev-gator/internal/config"
"fmt"

)


type state struct{
	config  *config.Config
}


func main() {
   name,args := args()
   
 
   commands := Commands{
	 cmdRegistry : make(map[string]func(*state,Command)error),
   }
   command := Command{
    name:name,
    args:args,
   }
   err := commands.register("login",handlerLogin)

    if err != nil {
	 //show the err?
   }

   configData,err := config.Read()
   if err != nil {
	 fmt.Errorf("Error while reading the file")
   }
   fmt.Println(configData)

    globalstate := state{
       config : &configData,
   }
 
   err = commands.run(&globalstate,command)
   
   if err != nil {
     fmt.Println(err)
   }

   configData,err = config.Read()
   if err != nil {
	 fmt.Errorf("Error while reading new contents")
   }
   fmt.Println(configData)

}