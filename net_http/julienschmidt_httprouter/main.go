package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func getCommandOutput(command string, args ...string) string {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(command, args...)

	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	err := cmd.Start()
	if err != nil {
		log.Fatal(fmt.Sprintf("%s : %s \n", err.Error(), stderr.String()))
	}

	err = cmd.Wait()

	if err != nil {
		log.Fatal(fmt.Sprintf("%s : %s \n", err.Error(), stderr.String()))
	}

	return stdout.String()
}

func getGoVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprint(w, getCommandOutput("/usr/local/go/bin/go", "version"))
}

func getFileContents(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprint(w, getCommandOutput("/usr/bin/cat", params.ByName("name")))
}

func main() {
	mux := httprouter.New()
	mux.GET("/api/go_version", getGoVersion)
	mux.GET("/api/read_file/:name", getFileContents)
	mux.ServeFiles("/static/*filepath", http.Dir("./static"))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
