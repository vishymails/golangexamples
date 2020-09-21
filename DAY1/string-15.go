/*
The Go string is a sequence of variable-width characters.

Go strings and text files occupy less memory or disk space. Since, UTF-8 is the standard, Go doesn't need to encode and decode strings.

Go Strings are value types and immutable. 
It means if you create a string, you cannot modify the contents of the string. The initial value of a string is empty "" by default.
*/


package main  
import ("fmt"  
	  "reflect"  
	  "strings"
)  
func main()  {  
   var x string = "Hello World"  
   fmt.Println(x)  
   fmt.Println(reflect.TypeOf(x))  



   str := "I love my country"  
   fmt.Println(len(str))  



   fmt.Println("Ascii value of A is ","A"[0])  


   str1 := "isro"  
   fmt.Println(strings.ToUpper(str1))  

/*
   str := "ISRO LIMITED"  
   fmt.Println(strings.ToLower(str)) 
   
   
   s := "ISRO"  
   fmt.Println(strings.HasPrefix(s,"IS"))  



   fmt.Println(strings.HasSuffix(s,"RO")) 
   
   

   var arr = []string{"a","b","c","d"}  
   fmt.Println(strings.Join(arr,"*"))  




   var str = "New "  
   fmt.Println(strings.Repeat(str,4))  



   str:= "Hi...there"  
   fmt.Println(strings.Contains(str,"th"))  



   str:= "Hi...there"  
   fmt.Println(strings.Index(str,"th"))  




   str:= "Hi...there"  
   fmt.Println(strings.Count(str,"e"))  



   str:= "Hi...there"  
   fmt.Println(strings.Replace(str,"e","Z",2))  



   str := "I,love,my,country"  
   var arr []string = strings.Split(str, ",")  
   fmt.Println(len(arr))  
   for i, v := range arr {  
	  fmt.Println("Index : ", i, "value : ", v)  
   }  




		fmt.Printf("%q\n", strings.Split("x,y,z", ","))  
		fmt.Printf("%q\n", strings.Split(" John and Jack and Johnny and Jinn ", "and"))  
		fmt.Printf("%q\n", strings.Split(" abc ", ""))  
		fmt.Printf("%q\n", strings.Split("", "Hello")) 
		
		

		fmt.Println(strings.Compare("a", "b"))  
		fmt.Println(strings.Compare("a", "a"))  
		fmt.Println(strings.Compare("b", "a"))  


		fmt.Println(strings.TrimSpace(" \t\n I love my country  \n\t\r\n"))  



   fmt.Println(strings.ContainsAny("Hello", "A"))  
   fmt.Println(strings.ContainsAny("Hello", "o & e"))  
   fmt.Println(strings.ContainsAny("Hello", ""))  
   fmt.Println(strings.ContainsAny("", ""))  

*/
}  