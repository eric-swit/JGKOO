package main

import "fmt"

func main() {
	var p_log_ida, p_log_idb string

	log_idb := "2202100524500U4l4GL"
	log_ida := "220211062560ft05aL6"

	if log_ida < log_idb {
		p_log_ida = log_ida
		p_log_idb = log_idb
	} else {
		p_log_ida = log_idb
		p_log_idb = log_ida
	}

	fmt.Println(p_log_ida, p_log_idb)
}
