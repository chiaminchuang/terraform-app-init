package main

import (
	"fmt"
	"terraform-app-init/pkg/initializer"
) 

func is_terraform_app() {}

func clear_all() {}


func main() {
	fmt.Println("Create Terraform app template")

	// root := os.Getwd()
	root := "test-wd"
	initializer.InitApp(root)
}