package configure

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	regexIPv4 = `(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`
	regexIPv6 = `((?:[0-9A-Fa-f]{1,4}))((?::[0-9A-Fa-f]{1,4}))*::((?:[0-9A-Fa-f]{1,4}))((?::[0-9A-Fa-f]{1,4}))*|((?:[0-9A-Fa-f]{1,4}))((?::[0-9A-Fa-f]{1,4})){7}`
)

type IPv4 [4]byte

/***
Function Name : matchIPAddr, Type: Internal
Utility : Comapre two IP Address of type IPv4 to return a boolean value
Return Type : bool
Created At: 01-06-2021
**/
func (p1 IPv4) matchIPAddr(p2 IPv4) bool {
	return strings.Compare(p1.makeIPAddr(), p2.makeIPAddr()) == 0
}

/***
Function Name : validateIPAddr, Type: Internal
Utility : Validates format of IP Address of type IPv4 against a regex
Return Type : bool
Created At: 01-06-2021
**/
func (ip IPv4) validateIPAddr() (b bool) {
	b, ok := regexp.MatchString(regexIPv4, ip.makeIPAddr())
	if ok != nil {
		fmt.Println("Error marching with IP regex!")
		os.Exit(1)
	}
	return
}

/***
Function Name : makeIPAddr, Type: Internal
Utility : Converts IPv4 to string type
Return Type : string
Created At: 01-06-2021
**/
func (ip IPv4) makeIPAddr() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

/***
Function Name : unmakeIPAddr, Type: Internal
Utility : Converts stirng to IPv4 type
Return Type : IPv4
Created At: 01-06-2021
**/
func unmakeIPAddr(str string) (ip IPv4) {
	if _, ok := fmt.Sscanf(str, "%d.%d.%d.%d", &ip[0], &ip[1], &ip[2], &ip[3]); ok != nil {
		fmt.Println("Error is parsing IP Address!")
		os.Exit(1)
	}
	return
}
