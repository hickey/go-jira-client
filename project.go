package gojira

import (
	"encoding/json"
	"fmt"
)


type Project struct {
	Self                    string
	Id                      string
	Key                     string
	Name                    string
	AvatarUrls              map[string]string
}


func (j *Jira) Projects() []Project {

	url := j.BaseUrl + j.ApiPath + "/project"
	contents := j.buildAndExecRequest("GET", url)

	var projects []Project
	err := json.Unmarshal(contents, &projects)
	if err != nil {
		fmt.Println("%s", err)
	}

	return projects
}

func (j *Jira) Project(id string) string {
    return "this is a test"
}