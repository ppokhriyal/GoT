package main

import (
	"fmt"
	"os"
	_ "got/aws"
)

// validate YAML config file
func CheckYamlConfig(yamlconfig string) bool {
	status := true

	if (yamlconfig != "aws.yaml" && yamlconfig != "docker.yaml") {
		status := false
		return status
	} else {
		// check if the YAML config file exists
		fileInfo,err := os.Stat(yamlconfig)
		if (err != nil ){
			status := false
			return status
		} else {
			if (fileInfo.IsDir()){
				status := false
				return status
			} else {
				if (fileInfo.Size() == 0) {
					status := false
					return status
				} else {
					return status
				}
			}
			
		}
		
	}

}

func main(){
	
	// validate the command-line args
	if (len(os.Args) < 2 || len(os.Args) > 2){
		fmt.Println("Invalid argument passeed.\n"+
		"Please pass the below YAML config files :\n"+
	    " \u2713 aws.yaml\n"+
	    " \u2713 docker.yaml")
	} else {
		yamlstatus := CheckYamlConfig(os.Args[1])
		if (yamlstatus == false) {
			fmt.Println("Invalid argument passeed.\n"+
		"Please pass the below YAML config files :\n"+
	    " \u2713 aws.yaml\n"+
	    " \u2713 docker.yaml")
		} else {
			fmt.Println("Building Project ")
		}
	}
}