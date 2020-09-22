package main  
import (  
   "fmt"  
)  
func main() {  
   ptr := new(int)  
   fmt.Println("Before change ptr",*ptr)  
   changePtr(ptr)  
   fmt.Println("After change ptr",*ptr)  
}  
func changePtr(ptr *int)  {  
   *ptr = 10  
}  