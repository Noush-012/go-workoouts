package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// calling api function for employee details

// get the response

//// calling api function for employee employer data
// get the response

// then aggregate both details and stored to 3rd api
type Employee struct {
	ID    int    `json:"id" binding:"ommitempty"`
	Name  string `json:"name" binding:"required,string"`
	Age   int    `json:"age" binding:"required,numeric"`
	Email string `json:"email" binding:"ommitempty"`
	// OrgID int    `json:"orgid" binding:"ommitempty"`
}

type Employer struct {
	OrgID    int    `json:"orgid" binding:"ommitempty"`
	Org      string `json:"org" binding:"required,string"`
	Location string `json:"location" binding:"required,string"`
}

type Data struct {
	Employee
	Employer
}

type Make struct {
	MakeID   int    `json:"Make_ID" binding:"required,numeric"`
	MakeName string `json:"Make_Name" binding:"required,string"`
}

type MakeData struct {
	Results []Make `json:"Results"`
}

func GetAllCarMakes() (Employee, error) {

	url := "https://vpic.nhtsa.dot.gov/api/vehicles/getallmakes?format=json"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return Employee{}, err
	}
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return Employee{}, err
	}

	response, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return Employee{}, err
	}

	str := string(response)
	fmt.Println(str[:400])
	// binding json to struct
	var data MakeData

	json.Unmarshal(response, &data)

	for i, v := range data.Results {

		fmt.Println("1", v)
		fmt.Println("2", i)
		fmt.Println("3", v.MakeID)
		fmt.Println("4", v.MakeName)

		if i == 3 {
			break
		}

	}

	return Employee{}, nil

}

func CallEmployeeDetails(id string) (Employee, error) {

	url := fmt.Sprintf("api/v1/employee/%s", id)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return Employee{}, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return Employee{}, err
	}
	response, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return Employee{}, err
	}
	fmt.Println(response)
	// binding json to struct
	var data Employee

	err = json.Unmarshal(response, &data)
	if err != nil {
		return Employee{}, err
	}
	return data, nil
}

func CallEmployerDetails(id string) (Employer, error) {

	url := fmt.Sprintf("api/v1/employer/%s", id)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return Employer{}, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return Employer{}, err
	}
	response, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return Employer{}, err
	}
	fmt.Println(response)
	// binding json to struct
	var data Employer
	err = json.Unmarshal(response, &data)
	if err != nil {
		return Employer{}, err
	}
	return data, nil
}

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

// type Emoployee struct {
// 	Name  string
// 	Age   int
// 	EmpID int
// }

// type Employer struct {
// 	ID       int
// 	ORG      string
// 	Location string
// }
