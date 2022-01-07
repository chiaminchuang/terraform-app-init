package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func is_terraform_app() {}

func clear_all() {}

func create_root(curr string) {
	// Create modules/, env/, main.tf, variables.tf, outputs.tf

	var err error
	err0 := os.Mkdir(path.Join(curr, "modules"), 0755)  // 
	if err0 != nil {
		log.Fatal(err)
	}

	err1 := os.Mkdir(path.Join(curr, "env"), 0755)
	if err1 != nil {
		log.Fatal(err)
	}

	_, err2 := os.Create(path.Join(curr, "main.tf"))
	if err2 != nil {
		log.Fatal(err)
	}

	_, err3 := os.Create(path.Join(curr, "variables.tf"))
	if err3 != nil {
		log.Fatal(err)
	}

	_, err4 := os.Create(path.Join(curr, "outputs.tf"))
	if err4 != nil {
		log.Fatal(err)
	}


}	

func main() {
	fmt.Println("Create Terraform app template")

	// curr := os.Getwd()
	curr := "test-wd"
	create_root(curr)

}