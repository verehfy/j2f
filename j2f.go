package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func exitGracefully(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func visit(current_object map[string]interface{}, base_folder string) {
	for k, v := range current_object {
		switch vv := v.(type) {
		case []interface{}:
			for i, u := range vv {
				fmt.Println(base_folder+k+"/"+strconv.Itoa(i), "=>", u)
				write_file(base_folder+k+"/"+strconv.Itoa(i), fmt.Sprint(i), fmt.Sprint(u))
			}
		case map[string]interface{}:
			visit(vv, base_folder+k+"/")
		default:
			fmt.Println(base_folder+k, "=>", vv)
			write_file(base_folder, k, fmt.Sprint(vv))
		}
	}
}

func write_file(file_path string, file_name string, file_value string) {
	os.MkdirAll(file_path, 0744)
	err := ioutil.WriteFile(file_path+file_name, []byte(file_value), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	jsonval, err := ioutil.ReadFile("test.json")
	check(err)

	var dat map[string]interface{}
	if err := json.Unmarshal(jsonval, &dat); err != nil {
		panic(err)
	}
	visit(dat, "artifacts/")

	exitGracefully(err)
}
