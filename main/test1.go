package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// func main() {

// 	var data Data

// 	EmployeeResp, err := CallEmployeeDetails()

// 	if err != nil {
// 		return
// 	}

// 	EmployerResp, err := CallEmployerDetails()

// 	data.Employee = EmployeeResp
// 	data.Employer = EmployerResp

// 	SaveData(data)

// }

// // calling api function for employee details

// // get the response

// //// calling api function for employee employer data
// // get the response

// // then aggregate both details and stored to 3rd api

// func CallEmployeeDetails() (Employee, error) {

// 	req, err := http.NewRequest("GET", "api/v1/employee/:id", nil)

// 	if err != nil {
// 		fmt.Println(err)
// 		return Employee{}, err
// 	}

// 	client := &http.Client{}

// 	resp, err := client.Do(req)

// 	if err != nil {
// 		fmt.Println(err)
// 		return Employee{}, err
// 	}

// 	response, err := io.ReadAll(resp.Body)

// 	if err != nil {
// 		fmt.Println(err)
// 		return Employee{}, err
// 	}
// 	fmt.Println(response)
// 	// binding json to struct
// 	// var data Employee

// 	return Employee{}, nil

// }

// func CallEmployerDetails() (Employer, error) {

// 	req, err := http.NewRequest("GET", "api/v1/employer/:id", nil)

// 	if err != nil {
// 		fmt.Println(err)
// 		return Employer{}, err
// 	}

// 	client := &http.Client{}

// 	resp, err := client.Do(req)

// 	if err != nil {
// 		fmt.Println(err)
// 		return Employer{}, err
// 	}

// 	response, err := io.ReadAll(resp.Body)

// 	if err != nil {
// 		fmt.Println(err)
// 		return Employer{}, err
// 	}
// 	fmt.Println(response)
// 	// binding json to struct
// 	// var data Employer

// 	return Employer{}, nil

// }

// func SaveData(data Data) error {

// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	// req, err := http.NewRequest("POST", "api/v1/database", jsonData)

// 	client := &http.Client{}

// 	resp, err := client.Do(req)

// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	fmt.Println(resp)

// }

// type Employee struct {
// 	Name  string `json:"Name"; binding: required`
// 	Age   int    `json:"Age"; binding: required`
// 	Email string `json:"Email"; binding: ommitempty`
// }

// type Employer struct {
// 	Org      string `json:Org; binding: required`
// 	location string `json:location; binding: required`
// }

// type Data struct {
// 	Employee
// 	Employer
// }
