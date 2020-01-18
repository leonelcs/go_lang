package main

import (
	"fmt"
	"strconv"
)

func main() {
	//convert boolean
	isNew := true
	fmt.Println("Iniciando Type Convertion " + strconv.FormatBool(isNew))
	//convert int
	number := int64(15)
	fmt.Println("Number: " + strconv.FormatInt(number, 10))

	numberInt := 100
	fmt.Println("Number: " + strconv.Itoa(numberInt))
	//convert float
	floatNumber := 1234.23456923
	fmt.Println("Float Number: " + strconv.FormatFloat(floatNumber, 'f', 5, 64))
	fmt.Println("Float Number: " + strconv.FormatFloat(floatNumber, 'f', -1, 64))

	//parse boolean
	isNewS := "true"
	isNewBool, err := strconv.ParseBool(isNewS)

	if err != nil {
		fmt.Println("failed")
	} else {
		if isNewBool {
			fmt.Println("IsNew")
		} else {
			fmt.Println("Not new")
		}
	}

	//parse integer
	numberS := "2"
	valueInt, err := strconv.ParseInt(numberS, 10, 32)

	if err == nil {
		if valueInt == 2 {
			fmt.Println("parsed")
		}

	} else {
		fmt.Println("Didn't parsed")
	}

	//parse float
	numberFloat := "2.2"
	valueFloat, errFloat := strconv.ParseFloat(numberFloat, 64)
	if errFloat != nil {
		fmt.Print("Error happened.")
	} else {
		if valueFloat == 2.2 {
			fmt.Println("Success")
		}
	}

	//byte array to string
	helloWorldByte := []byte{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100}
	fmt.Println(string(helloWorldByte))

	//the other way aroung
	helloWorld := "Hello, World"
	fmt.Println([]byte(helloWorld))
}
