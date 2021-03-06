package main 
  
import "fmt"
  
// Defining a struct type 
type Address struct { 
    Name    string 
    city    string 
    Pincode int
} 
  
func main() { 
  
    // Declaring a variable of a `struct` type 
    // All the struct fields are initialized  
    // with their zero value 
    var a Address  
    fmt.Println(a) 
  
    // Declaring and initializing a 
    // struct using a struct literal 
    a1 := Address{"Maheswar Reddy Kamasani", "Galiveedu", 516216} 
  
    fmt.Println("Address1: ", a1) 
  
    // Naming fields while  
    // initializing a struct 
    a2 := Address{Name: "Chinnapa Reddy Kanakanti", city: "Rayachoti", 
                                 Pincode: 516214} 
  
    fmt.Println("Address2: ", a2) 
  
    // Uninitialized fields are set to 
    // their corresponding zero-value 
    a3 := Address{Name: "Kadapa"} 
    fmt.Println("Address3: ", a3) 
}