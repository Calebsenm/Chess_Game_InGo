// this a example of the concept of the 
// os module for clear the console 



package main;

import(
  "os/exec"
  "os"
  "fmt"

)

func main (){

  fmt.Println("Text1");
  
  cmd := exec.Command("clear");
  cmd.Stdout = os.Stdout
  cmd.Run();
  fmt.Println("text2");

  im := exec.Command("clear")
  im.Stdout = os.Stdout;
  cmd.run();


}

