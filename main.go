package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

)

//User defines model for storing account details in database
type office struct {

	// office
	OfficeCode   string `json:"officeCode,omitempty"`
	OfficeName   string `json:"officeName,omitempty"`
	City         string `json:"city,omitempty"`
	AddressLine1 string `json:"addressLine1,omitempty"`
	Address      string `json:"address,omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	PostalCode   string `json:"postalCode,omitempty"`
	Street       string `json:"street,omitempty"`
	Country      string `json:"country,omitempty"`
	PhoneNumber  int    `json:"phonenumber,omitempty"`
}

type employee struct {
	// employee
	EmployeeNumber int    `json:"employeeNumber,omitempty"`
	OfficeCode     string `json:"officeCode,omitempty"`
	FirstName      string `json:"firstName,omitempty"`
	LastName       string `json:"lastName,omitempty"`
	Email          string `json:"email,omitempty"`
	JobTitle       string `json:"job,omitempty"`
	PhoneNumber    int    `json:"phonenumber,omitempty"`
}

func main() {
	fmt.Println("Schema Change up")
	mux := http.NewServeMux()

	mux.HandleFunc("/", changeRequestForUpdatedVersion)

	http.ListenAndServe(":50002", mux)
}

func changeRequestForUpdatedVersion(w http.ResponseWriter, r *http.Request) {
	log.Println("POST request received from localhost:50000")
	//log.Println(r)
	//log.Println(r.Body)
	if r.URL.Path == "/addFormOffice" {
		dbOffice := office{} //initialize
		//Parse json request body and use it to set fields on db
		err := json.NewDecoder(r.Body).Decode(&dbOffice)
		fmt.Println(dbOffice)
		if err != nil {
			panic(err)
		}
		log.Println("Apply office table DB changes Shcema for POST Request")
		// apply schema changes
		dbOffice.Address = dbOffice.AddressLine1 + " " + dbOffice.AddressLine2
		dbOffice.AddressLine1 = ""
		dbOffice.AddressLine2 = ""

		//Marshal or convert user object back to json and write to response
		dbJson, err := json.Marshal(dbOffice)
		if err != nil {
			panic(err)
		}
		responseBody := bytes.NewBuffer(dbJson)
		url := "http://localhost:3000" + r.URL.String()
		resp, err := http.Post(url, "application/json", responseBody)
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer resp.Body.Close()

	}

	if r.URL.Path == "/addFormEmployee" {
		dbEmp := employee{} // initailize
		err := json.NewDecoder(r.Body).Decode(&dbEmp)
		fmt.Println(dbEmp)

		if err != nil {
			panic(err)
		}
		log.Println("Apply Employee table DB changes Shcema for POST Request")
		dbEmp.PhoneNumber = 07100000
		//Marshal or convert user object back to json and write to response
		dbJson, err := json.Marshal(dbEmp)
		if err != nil {
			panic(err)
		}
		responseBody := bytes.NewBuffer(dbJson)
		url := "http://localhost:3000" + r.URL.String()
		resp, err := http.Post(url, "application/json", responseBody)
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer resp.Body.Close()
	}
}
