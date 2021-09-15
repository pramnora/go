package main

import "fmt"

func main() {
     for tablesNo:=1;tablesNo<=12;tablesNo++{
       fmt.Println(tablesNo, " X tables:-")
       fmt.Println()
       for timesNo:=1;timesNo<=12;timesNo++{
	  fmt.Println(timesNo," X ",tablesNo," = ", timesNo*tablesNo)
	}
	fmt.Println()
     }
}
