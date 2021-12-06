// package main

// import (
// 	"encoding/xml"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// )

// // XML to Go

// func main() {
// 	xmlFile, err := os.Open("Employees.xml")
// 	if err != nil {
// 		fmt.Println("Opening file error : ", err)
// 		return
// 	}
// 	defer xmlFile.Close()

// 	xmlData, _ := ioutil.ReadAll(xmlFile)

// 	// Unmarshalling
// 	var c Company
// 	// xmlData'ya gelen byte[]'si formatındaki veriyi Company'nin nesne örneğine aktarır/dönüştür.
// 	xml.Unmarshal(xmlData, &c)

// 	// fmt.Println(c.Persons)
// 	for i := 0; i < 3; i++ {
// 		fmt.Println(c.Persons[i])
// 	}

// }

// type Person struct {
// 	XMLName   xml.Name `xml:"person"`
// 	ID        int      `xml:"id"`
// 	FirstName string   `xml:"firstname"`
// 	LastName  string   `xml:"lastname"`
// 	UserName  string   `xml:"username"`
// }

// type Company struct {
// 	XMLName xml.Name `xml:"company"`
// 	Persons []Person `xml:"person"`
// }

// func (p Person) String() string {
// 	return fmt.Sprintf("\t ID: %d - FirstName : %s %s -> UserName : %s",
// 		p.ID, p.FirstName, p.LastName, p.UserName)
// }
//*********************************** JSON işlemleri ************************************
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// JSON Operations in Go

func main() {

	person := Person{
		ID:        7,
		FirstName: "Cihan",
		LastName:  "Özhan",
		UserName:  "cihanozhan",
		Gender:    "Erkek",
		Name:      Name{Family: "xyz", Personal: "Cihan"},
		Email: []Email{
			{ID: 1, Kind: "work", Address: "cihan@deeplab.co"},
			{ID: 2, Kind: "home", Address: "cihan.ozhan@hotmail.com"},
		},
		Interest: []Interest{
			{ID: 1, Name: "C#"},
			{ID: 2, Name: "Go"},
			{ID: 3, Name: "Python"},
			{ID: 4, Name: "SQL Server"},
			{ID: 5, Name: "PostgreSQL"},
		},
	}

	WriteMessage("Reading Operation Started.")

	WriteMessage("Personal Fullname")
	WriteStarLine()
	res := GetPerson(&person)
	WriteMessage(res)
	WriteStarLine()

	WriteMessage("\n")

	WriteMessage("Personal Email with Index")
	WriteStarLine()
	resEmail := GetPersonEmailAddress(&person, 0)
	WriteMessage(resEmail)
	WriteStarLine()

	WriteMessage("\n")

	WriteMessage("Personal Email Object with Index")
	WriteStarLine()
	resEmail2 := GetPersonEmail(&person, 0)
	fmt.Println(resEmail2)
	WriteStarLine()

	WriteMessage("Reading Operation Ended.")
	WriteMessage("\n")
	WriteMessage("Writing Operation Started.")
	SaveJSON("person.json", person)
	WriteMessage("Writing operation ended.")
}

// Struct Nesnelerimiz

type Person struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Gender    string
	Name      Name
	Email     []Email
	Interest  []Interest
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	ID      int
	Kind    string
	Address string
}

type Interest struct {
	ID   int
	Name string
}

// Function'lar

func GetPerson(p *Person) string {
	return p.FirstName + " " + p.LastName
}

func GetPersonEmailAddress(p *Person, ind int) string {
	return p.Email[ind].Address
}

func GetPersonEmail(p *Person, ind int) Email {
	return p.Email[ind]
}

func WriteMessage(msg string) {
	fmt.Println(msg)
}

func WriteStarLine() {
	fmt.Println("****************")
}

func SaveJSON(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encoder(key)
	checkError(err)
	outFile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error : ", err.Error())
		os.Exit(1)
	}
}
