package  main


import (
	"fmt"
	"os"
	"io/ioutil"
	)

func check(err error){
	if err!=nil{
	panic(err)
	}
	}

func readFile(name string)([]byte){
	content,err:=ioutil.ReadFile(name);
	check(err);
	return content
	}

func main(){
	if len(os.Args)==1{
		fmt.Println("Please Provide a valid filename to parse")
		fmt.Println("For e.g $./csv filename_of_file_to_be_parsed")
	}else if len(os.Args)>1{
		parse(readFile(string(os.Args[1])))
		}
		}


func parse(arr[]byte){
	comma_counter:=0
	i:=0
	for string(arr[i])!="\n"{
		if string(arr[i])==","{
			comma_counter++;
			}
			i++;
			}
		fmt.Println(comma_counter);
	}


