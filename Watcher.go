/*
Coded by jaypyon
scorpion@dgu.ac.kr
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)
//import "time"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func exec_command_sha(hash string) string{
	fmt.Println("SHA-1 : ",hash)
	lsCmd := exec.Command("bash", "-c", `find . -type f -exec sh -c '
   sha1sum "$2" | cut -f 1 -d " " | sed "s|^\\\\||" | grep -Eqi "$1"
' find-sh "`+ hash +`" {} \; -print`)
	lsOut, err := lsCmd.Output()
	check(err)

	return string("SHA-1: "+hash+" : "+string(lsOut)+"\n")
}

func exec_command_md5(hash string) string{
	fmt.Println("MD5 : ",hash)
	lsCmd := exec.Command("bash", "-c", `find . -type f -exec sh -c '
   md5sum "$2" | cut -f 1 -d " " | sed "s|^\\\\||" | grep -Eqi "$1"
' find-sh "`+ hash +`" {} \; -print`)
	lsOut, err := lsCmd.Output()
	check(err)

	return string("MD5: "+hash+" : "+string(lsOut)+"\n")
}
func main() {
	
	var hashes = make([]string,0)
	//fmt.Println(os.Getwd())
	f, err := os.Open("./hash.txt")
	check(err)
	defer f.Close()
	
	fo, err := os.Create("result.txt")
	check(err)
	defer fo.Close()
	
	reader:= bufio.NewReader(f)
	for{
		line, isPrefix, err := reader.ReadLine()
		if isPrefix||err!=nil{
			break
		}
		hashes = append(hashes,string(line))
		
	}
	fmt.Println(hashes)
	
	
	
	for _,value := range hashes {
		if len(value)==40{
			_, err = fo.WriteString(exec_command_sha(value))
			check(err)
		}
		if len(value)==32{
			_, err = fo.WriteString(exec_command_md5(value))
			check(err)
		}
	}
	
}
