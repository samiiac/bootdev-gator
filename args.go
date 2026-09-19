package main 
import ("os"
"fmt"
)

func args()(string,[]string) {
   args := os.Args
   if len(args) < 2 {
    fmt.Errorf("No commands provided")
    os.Exit(1)
   }

   args = args[1:]
   name := args[0]
   args = args[1:]
   fmt.Println(args)
   return name,args
}

