package  main


import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
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
	comma_counter++;
	no_col:=comma_counter;
	no_rows:=(strings.Count(string(arr),",")/comma_counter)-1;
	fmt.Printf("no columns:- %d, no rows:- %d\n",no_col,no_rows);
	dumb:=strings.Split(string(arr),"\r")
	data:=[][]string{}
	for p,j:=range dumb {
		if p==0 {continue}
		pop:=strings.Split(j[1:],",")
		data=append(data,pop)
		}
	for _,cou:=range data{
		for _,k:=range cou{
			fmt.Printf("%q ",k)
			}
			fmt.Println("\n")
			}
	}


