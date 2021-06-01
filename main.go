package main

import (
	conf "./configure"
)

func main() {
	//	r := os.Args[1:]
	//	conf.ParseCommandLineArgs(r)
	conf.GetRouteTable()
}
