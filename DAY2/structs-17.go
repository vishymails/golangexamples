/*
In Go, Struct can be used to create user-defined types.

Struct is a composite type means it can have different properties and each property can have their own type and value.

Struct can represent real-world entity with these properties. We can access a property data as a single entity. 

It is also valued types and can be constructed with the new() function.

*/



package main  
import (  
   "fmt"  
)  
type person struct {  
   firstName string  
   lastName  string  
   age       int  
}  





type person1 struct {  
	fname string  
	lname string	
}  

 type employee struct {  
	person1  
	empId int  
 }  
 func (p person1) details() {  
	fmt.Println(p, " "+" I am a person")  
 }  
 func (e employee) details() {  
	fmt.Println(e, " "+"I am a employee")  
 }  







func main() {  
   x := person{age: 30, firstName: "John", lastName: "Anderson", }  
   fmt.Println(x)  
   fmt.Println(x.firstName)  





   p1 := person1{"Raj", "Kumar"}  
   p1.details()  
   e1 := employee{person1:person1{"John", "Ponting"}, empId: 11}  
   e1.details()  





}  



