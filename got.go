package main

import (
	"fmt"
	"os"
	_ "got/aws"
)


func main(){
	// passing yaml file name
	if (len(os.Args) < 2 || len(os.Args) > 2) {
		fmt.Println("Invalid arguments passed")
	} else {
		// check if yaml file is aws or docker
		if (os.Args[1] != "aws.yaml" && os.Args[1] != "docker.yaml"){
			fmt.Println("Invalid YAML config file")
		} else {
			fmt.Println("Building Terraform Config file from "+os.Args[1]+" ...")
		}
	}
}