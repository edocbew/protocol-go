package configure

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type CommandLineArgs struct {
	Key, Value string
}

/***
Function Name : readArgsFileMap, Type: Internal
Utility : Reads the argument map from data/args.json file to compare with User CLI input
Return Type : map[string]int
Created At: 28-05-2021
**/
func readArgsFileMap() (argsDefaultMap map[string]int) {
	argsFile, err := os.OpenFile("data/args.json", os.O_RDONLY, 0444)
	if err != nil {
		fmt.Println("Args File doesn't exist. Please check configuration!")
		os.Exit(1)
	}
	defer argsFile.Close()
	argsFileJSONReader, _ := ioutil.ReadAll(io.Reader(argsFile))
	json.Unmarshal(argsFileJSONReader, &argsDefaultMap)
	return
}

/***
Function Name : getCommandLineArgs, Type: Internal
Utility : Reads the arguments passed by User in CLI and creates an array of maps for parsing
Return Type : []CommandLineArgs (type struct)
Created At: 28-05-2021
**/
func getCommandLineArgs(args []string) (argsMap []CommandLineArgs) {
	var key, value string
	var param bool = false
	for _, s := range args {
		if s[0] == byte('-') && len(s) > 1 {
			argsMap = validateOptions(param, value, argsMap, key)
			key, param = s, true
		} else {
			value, param = s, false
			argsMap = append(argsMap, CommandLineArgs{key, value})
		}
	}
	argsMap = validateOptions(param, value, argsMap, key)
	return
}

/***
Function Name : validateOptions, Type: Internal
Utility : Validates the arguments and their coressponding parameters to append a map to an array
Return Type : []CommandLineArgs (type struct)
Created At: 28-05-2021
**/
func validateOptions(param bool, value string, argsMap []CommandLineArgs, key string) []CommandLineArgs {
	if param {
		value, param = "", false
		argsMap = append(argsMap, CommandLineArgs{key, value})
	}
	return argsMap
}

/***
Function Name : ParseCommandLineArgs, Type: External
Utility : Calls the readArgsFileMap and getCommandLineArgs to parse the Command line arguments and invoke required functions
Return Type : nil
Created At: 31-05-2021
**/
func ParseCommandLineArgs(args []string) {
	argsDefaultMap := readArgsFileMap()
	argsMap := getCommandLineArgs(args)
	for _, args := range argsMap {
		switch mappedVal := argsDefaultMap[args.Key]; mappedVal {
		case 1:
			fmt.Println("Init")
		case 2:
			fmt.Println("IP Address")
		}
	}
}
