package main

import (
	"fmt"
	"net/http"
	"path"
	"time"
	"io/ioutil"
	"regexp"
)

const status_file_dir = "/var/check-receiver/status-files/"

func handler(w http.ResponseWriter, r *http.Request) {
	if (r.Method != "POST") {
		w.WriteHeader(405)
		fmt.Fprintf(w, "send POST to update status_file")
		return
	}

	status_name := r.URL.Path[1:]
	matched, err := regexp.MatchString("^[a-zA-Z0-9\\-\\_]+$", status_name)
	if (! matched || err != nil) {
		w.WriteHeader(400)
		fmt.Fprintf(w, "status_name may only consist of 'a-z A-Z 0-9 - _'")
		return
	}

	current_time := time.Now().UTC().Format(time.RFC3339)

	status_file := path.Join(status_file_dir, status_name)

	fmt.Println("updating status-file:", status_file)
	err = ioutil.WriteFile(status_file, []byte(current_time), 0644)
	if (err != nil) {
		fmt.Println("failed to update status-file ", status_file, ": ", err)
		w.WriteHeader(500)
		fmt.Fprintf(w, "failed to update status_file")
		return
	}

	fmt.Fprintf(w, "UTC: " + current_time)
	return
}

func main() {
	fmt.Println("STATUS_FILE_DIR:", status_file_dir)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
