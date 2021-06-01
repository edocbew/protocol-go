package configure

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const (
	routeTableFileName = "data/route.json"
)

type Route struct {
	IPAddr    IPv4 `json:"ipaddr"`
	TransKeys Keys `json:"keys"`
}

type Routes struct {
	Routes []Route `json:"routes"`
}

/***
Function Name : validateIPAddr, Type: Internal
Utility : Validates format of IP Address of type IPv4 against a regex
Return Type : bool
Created At: 01-06-2021
**/
func GetRouteTable() {
	var routeTable Routes
	routeFile, err := os.OpenFile(routeTableFileName, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("File doesn't exist!")
		os.Exit(1)
	}
	defer routeFile.Close()
	routeJSONBytes, _ := ioutil.ReadAll(io.Reader(routeFile))
	json.Unmarshal(routeJSONBytes, &routeTable)
	fmt.Println(routeTable.Routes)
	for _, route := range routeTable.Routes {
		fmt.Println("IP : ", route.IPAddr.makeIPAddr(), "Key : ", route.TransKeys)
	}
}
