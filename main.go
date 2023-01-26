package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/open-policy-agent/opa/rego"
)

func main() {
	// Define the input data
	var age = 27
	input := map[string]interface{}{
		"age":     age,
		"company": "averlon",
	}

	// Read the contents of the Rego file
	b, err := ioutil.ReadFile("C:/Users/puneeth.sharma_averl/hello/Go/regoquery/example.rego")
	if err != nil {
		fmt.Println("Error reading Rego file:", err)
		return
	}

	// Parse the Rego file
	r := rego.New(
		rego.Query("data.example.allow_access"),
		rego.Module("example.rego", string(b)),
		rego.Input(input),
	)

	rs, err := r.Eval(context.Background())

	if err != nil {
		fmt.Println("Error evaluating Rego file:", err)
		return
	}
	//fmt.Println(age)
	// Print the result
	if len(rs) == 0 {
		fmt.Println("You are not eligible.")
	} else {
		fmt.Println("You have access.")
	}

}
