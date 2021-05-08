package main

import (
	"fmt"
	"got/aws"
)

func main(){
	fmt.Println("YAML to Terraform Configuration")
	a := aws.AwsInit()
	fmt.Println(a)
}