// package main

// import "fmt"

// // Struct

// type Human struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// }

// func main() {
// v1
// humX := Human{FirstName: "Murtaza", LastName: "Sayın"}
// fmt.Println("Data :" + humX.FirstName)
// fmt.Println("Data :" + humX.LastName)
//humX := Human{"Murtaza","Sayın", 36}   // Buşekilde kullanıldığında tüm parametreler girilmeli

// humX := Human{}
// humX.FirstName = "Mehmet"
// fmt.Println("Data :" + humX.FirstName)
// fmt.Println("Data :" + humX.LastName)

// v2

// humY := new(Human)
// humY.FirstName = "Ali"
// humY.LastName = "BAL"
// fmt.Println("Data: " + humY.FirstName)

// v3
// 	xx := NewHuman("CO", "Ozhan")
// 	fmt.Println(xx.FirstName + " " + xx.LastName)

// }
// func NewHuman() *Human{			// Bu da parametreleri defaul değerleri ile döndürür.
// 	return new(Human)
// }

// func NewHuman(FirstName, LastName string) *Human {
// 	x := new(Human)
// 	x.FirstName = FirstName
// 	x.LastName = LastName
// 	return x
// }
//*****************************************************************
// package main

// import "fmt"

// func main() {
// 	x := new(User)
// 	x.FirstName = "E"
// 	x.LastName = "B"
// 	 result := GetFullName(x)
//  result := x.GetFullName()
//  fmt.Println(result)
// 	x.SetFirstName("co")
// 	fmt.Println(x.FirstName)

// }

// type User struct {
// 	ID        string
// 	FirstName string
// 	LastName  string
// 	UserName  string
// 	Age       int
// }

// func NewUser() *User {
// 	return new(User)
// }

// Function Version
// func GetFullName(user *User) string {
// 	return user.FirstName + " " + user.LastName
// }
// Method Version
//Methodlar bir nesneye aittir.
// func (u *User) GetFullName() string {
// 	return u.FirstName + " " + u.LastName
// }

// * kullanmadığımız versiyonda hafızadaki hangi referansın güncellendiğini bilemeyiz.
// En azından bizim nesne örneğimizin güncellenmeyeceği garantidir.
// func (u *User) SetFirstName(firstname string) {
// 	u.FirstName = firstname
// }
//******************************************
// package main

// import (
// 	"fmt"

//"github.com/cihanozhan/userpayment/models"------> Genellikle bu mod kullanılacak
// models "github.com/cihanozhan/userpayment/models"
//x "github.com/cihanozhan/userpayment/models"
// 	"github.com/cihanozhan/userpayment/models"
// 	"github.com/cihanozhan/userpayment/utils"
//_ "github.com/cihanozhan/userpayment/models"
// )

// func main() {

// v1

// 	u1 := &User{
// 		ID:        5,
// 		FirstName: "Ali",
// 		LastName:  "Veli",
// 		UserName:  "aliveli",
// 		Age:       29,
// 		Pay: &Payment{
// 			Salary: 3.555,
// 			Bonus:  522,
// 		},
// 	}

// 	fmt.Println(u1.GetFullName())
// 	fmt.Println(u1.GetPayment())

// }

// v2

// User Creation
// user := models.NewUser()
// user.FirstName = "Cihan"
// user.LastName = "Özhan"
// user.Age = 33
// user.UserName = "cihanozhan"

//Payment Creation v1

// user.Pay.Salary=50
// user.Pay.Bonus = 5

//Printing
// fmt.Println(user.Pay)
// fmt.Println(user.GetPayment())

// Payment Creation v2
// user.Pay = &Payment{Salary: 45, Bonus: 3}
// fmt.Println(user.GetFullName())
// fmt.Println(user.GetPayment())

// randX := utils.Random(10, 99)
// fmt.Println("Rand value: ", +strconv.Itoa(randX))
// 	fmt.Println(GetSalary())
// }
// func GetSalary() int {
// 	return 500 + utils.Random(10, 99)
// }

// Constructor
//******************************************* INTERFACE  *********************
package main

import (
	"fmt"
	"strconv"
)

// Polymoprhism (Çok Biçimli)
// Interface
// Type Embeddings

func main() {

	// Ferrari
	ferr := new(Ferrari)
	ferr.Brand = "Ferrari"
	ferr.Model = "F50"
	ferr.Color = "Red"
	ferr.Speed = 340
	ferr.Price = 1.4
	ferr.Special = true
	// fmt.Println(ferr.Information())

	// Mercedes
	merc := new(Mercedes)
	merc.Brand = "Mercedes"
	merc.Model = "CLX"
	merc.Color = "Black"
	merc.Speed = 290
	merc.Price = 0.4
	// fmt.Println(merc.Information())

	CarExecute(ferr)
	CarExecute(merc)
}

func CarExecute(c Carface) {
	fmt.Println("\n")
	fmt.Println("Araç Bilgi: \n" + c.Information())
	fmt.Println("\n")

	msg := ""

	isRun := c.Run()
	if isRun {
		msg = "çalışıyor"
	} else {
		msg = "çalışmıyor"
	}
	fmt.Println("Araç " + msg + ".")

	isStop := c.Stop()
	if isStop {
		msg = "durdu"
	} else {
		msg = "durmuyor, fren tutmuyor!"
	}
	fmt.Println("Araç " + msg + ".")
}

// Struct'lar

type Car struct {
	Brand string
	Model string
	Color string
	Speed int
	Price float64
}

type Carface interface {
	Run() bool
	Stop() bool
	Information() string
}

type SpecialProduction struct {
	Special bool
}

// Struct Nesneleri : Ferrari

type Ferrari struct {
	Car
	SpecialProduction
	/*
	   Type Embeddings yöntemi.
	   Object Composition(Inheritance)
	*/
}

func (_ Ferrari) Run() bool {
	return true
}

func (_ Ferrari) Stop() bool {
	return true
}

func (f *Ferrari) Information() string {
	retVal := "\t" + f.Brand + " " + f.Model + "\n" + "\t" + "Color: " + f.Color + "\n" + "\t" + "Speed: " + strconv.Itoa(f.Speed) + "\n" + "\t" + "Price: " + strconv.FormatFloat(f.Price, 'g', -1, 64) + " Million"
	add := "Yes"
	if f.Special {
		retVal += "\n" + "\t" + "Special: " + add
	}
	return retVal
}

// Struct Nesneleri : Lamborghini

type Lamborghini struct {
	Car
	SpecialProduction
}

func (_ Lamborghini) Run() bool {
	return true
}

func (_ Lamborghini) Stop() bool {
	return true
}

func (f *Lamborghini) Information() string {
	retVal := "\t" + f.Brand + " " + f.Model + "\n" + "\t" + "Color: " + f.Color + "\n" + "\t" + "Speed: " + strconv.Itoa(f.Speed) + "\n" + "\t" + "Price: " + strconv.FormatFloat(f.Price, 'g', -1, 64) + " Million"
	add := "Yes"
	if f.Special {
		retVal += "\n" + "\t" + "Special: " + add
	}
	return retVal
}

// Struct Nesneleri : Mercedes

type Mercedes struct {
	Car
}

func (_ Mercedes) Run() bool {
	return true
}

func (_ Mercedes) Stop() bool {
	return true
}

func (f *Mercedes) Information() string {
	return "\t" + f.Brand + " " + f.Model + "\n" + "\t" + "Color: " + f.Color + "\n" + "\t" + "Speed: " + strconv.Itoa(f.Speed) + "\n" + "\t" + "Price: " + strconv.FormatFloat(f.Price, 'g', -1, 64) + " Million"
}
