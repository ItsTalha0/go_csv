package  csvread


import (
	"fmt"
	"io/ioutil"
	"strings"
	)

func check(err error){
	if err!=nil{
	panic(err)
	}
	}

func ReadFile(name string){
	content,err:=ioutil.ReadFile(name);
	check(err);
	parse(content)
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
			fmt.Println()
			}
	}


