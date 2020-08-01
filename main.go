package  main


import (
	"fmt"
	"os"
	"io/ioutil"
	)
func check(err error){
	panic(err)
	}


func main(){
	fileName:=os.Args[1]
	content,err:=ioutil.ReadFile(fileName);
	check(err);
	fmt.Println(content)
	}

	
