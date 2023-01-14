package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// DeployRequest struct for the json request body
type DeployRequest struct {
	Project  string `json:"project"`
	Version  string `json:"version"`
	Time     string `json:"time"`
	JobID    int    `json:"jobId"`
	Token    string `json:"token"`
}

var allowedProjects []string

func init() {
	// read allowed projects from config file
	content, err := ioutil.ReadFile("projects.conf")
	if err != nil {
		log.Fatal(err)
	}
	allowedProjects = strings.Split(string(content), "\n")
}
func main() {
	http.HandleFunc("/api/deploy", deployHandler)
	http.HandleFunc("/", notFoundHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func deployHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		log.Printf("Invalid request method: %s", r.Method )

		return
	}

	var deployReq DeployRequest
	err := json.NewDecoder(r.Body).Decode(&deployReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Invalid request body: %s", r.Body )
		return
	}

	// validate request
	if !validateRequest(deployReq) {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// execute command
	cmd := fmt.Sprintf("cd '/opt/%s' && make VERSION='%s' refresh", deployReq.Project, deployReq.Version)
	output, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Printf("Error executing command: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("Command output: %s", output)
	fmt.Fprint(w, "Deployment successful")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not found", http.StatusNotFound)
}
func validateRequest(req DeployRequest) bool {
	// validate types
	if req.JobID <= 0 {
		log.Printf("Bad jobID: %d", req.JobID)
		return false

	}

	// validate project
	found := false
	for _, p := range allowedProjects {
		if p == req.Project {
			found = true
			break
		}
	}
	if !found {
		log.Printf("Project not supported: %s", req.Project)
		return false
	}

	// validate token
	if req.Token != os.Getenv("API_TOKEN") {
		log.Printf("Bad API_TOKEN: %s", os.Getenv("API_TOKEN"))
		return false
	}

	return true
}

