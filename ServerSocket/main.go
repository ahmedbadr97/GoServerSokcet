package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)
import "strings"

type location struct {
	description string
	neighboursCount int
}
var maxLocation location
var locations = make(map[string]int)


func main () {
	// listen on port 8000
	go listenForClientRequests()
	fmt.Println("please Enter one of the next choices\n 1-Print location that most neighbors agrees on now \n 2-print all locations \n 0-exit\n ")
	reader := bufio.NewReader(os.Stdin)
	var input int
	for{
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		input, _ = strconv.Atoi(text)

		switch input {
		case 1:
			if maxLocation.neighboursCount!=0{
				fmt.Printf("the location that most neighbors agrees on now is  %s\niam here salama------------------------------------^^",maxLocation.description)
			}
			break
		case 2:
			printLocations(locations)
			break
		case 0:
			os.Exit(0)
		default:
			continue

		}
	}


}
func listenForClientRequests()  {
	fmt.Println("where are you Gehad ...")

	ln, _ := net.Listen("tcp", ":5000")
	for{
		conn, _ := ln.Accept()
		go serveClient(conn)
	}

}
func printLocations(locations map[string]int)  {
	for s, i := range locations {
		fmt.Printf("Location:%s  noOfNeighbours:%d \n",s,i)
	}
}
func addLocation(locations map[string]int, location string)   {
	/* test if entry is present in the map or not*/
	lCaseLocation :=strings.ToLower(location)
	count, found := locations[lCaseLocation]
	/* if ok is true, entry is present otherwise entry is absent*/
	if found {
		locations[lCaseLocation]=count+1
		if (count>maxLocation.neighboursCount) {
			maxLocation.description=location
			maxLocation.neighboursCount=count
		}

	} else {
		if len(locations)==0{
			maxLocation.description=location
			maxLocation.neighboursCount=1
		}
		locations[lCaseLocation]=1
	}

}
func serveClient(conn net.Conn)  {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if len(message)!=0{
		addLocation(locations,message)
		}
}
