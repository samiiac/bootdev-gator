package main 
import ("github.com/samiiac/bootdev-gator/internal/config"
"github.com/samiiac/bootdev-gator/internal/database"
"fmt"
_ "github.com/lib/pq"
"database/sql"
)


type state struct{
   db      *database.Queries
	config  *config.Config
}


func main() {
   name,args := args()
   
 
   //init commands and register handler
   commands := newCommands()
   command := Command{
    name:name,
    args:args,
   }
  
   //read config  data ,open db connection and init state
   configData,err := config.Read()
   if err != nil {
	 fmt.Errorf("Error while reading the file")
   }

    db,err := sql.Open("postgres",configData.DbUrl) 
    if err != nil {
	 fmt.Println(err)
   }

   dbQueries := database.New(db)
  

    globalstate := state{
       config : &configData,
       db:dbQueries,
   }
 
   //run command
   err = commands.run(&globalstate,command)
   
   if err != nil {
     fmt.Println(err)
   }


}