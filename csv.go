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
	no_col:=comma_counter+1;
	no_rows:=(strings.Count(string(arr),",")/comma_counter)-1;
	fmt.Printf("no columns:- %d, no rows:- %d\n",no_col,no_rows);
	f:=func (c rune)bool{
		return string(c)==","
		}
/*	g:=func (c rune)bool{
		return string(c)=="\n"
		}*/
//	data:=make(map[string][9]float64)
	fmt.Println(strings.FieldsFunc(string(arr),f)[8:]);

}


