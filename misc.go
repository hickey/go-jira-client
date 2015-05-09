package gojira

import (
	"encoding/json"
	"fmt"
)


type Priority struct {
    Self                    string
    StatusColor             string
    Description             string
    IconUrl                 string
    Name                    string
    Id                      string
}

type Resolution struct {
    Self                    string
    Id                      string
    Description             string
    Name                    string
}

type JiraServerInfo struct {
    BaseUrl                 string
    Version                 string
    VersionNumbers          []uint
    BuildNumber             uint
    BuildDate               string
    ServerTime              string
    ScmInfo                 string
    ServerTitle             string
    HealthChecks            []struct {
        Name                    string
        Description             string
        Passed                  bool
    }
}

type Status struct {
    Self                    string
    Description             string
    IconUrl                 string
    Name                    string
    Id                      string
    StatusCategory          *StatusCategory
}

type StatusCategory struct {
    Self                    string
    Id                      uint
    Key                     string
    ColorName               string
    Name                    string
}

func (j *Jira) Priorities() []Priority {

	url := j.BaseUrl + j.ApiPath + "/priority"
	contents := j.buildAndExecRequest("GET", url)

	var priorities []Priority
	err := json.Unmarshal(contents, &priorities)
	if err != nil {
		fmt.Println("%s", err)
	}

	return priorities
}


func (j *Jira) Resolutions() []Resolution {

	url := j.BaseUrl + j.ApiPath + "/resolution"
	contents := j.buildAndExecRequest("GET", url)

	var resolutions []Resolution
	err := json.Unmarshal(contents, &resolutions)
	if err != nil {
		fmt.Println("%s", err)
	}

	return resolutions
}


func (j *Jira) ServerInfo(doHealthChecks bool) JiraServerInfo {

	url := j.BaseUrl + j.ApiPath + fmt.Sprintf("/serverInfo?doHealthChecks=%s", doHealthChecks)
	contents := j.buildAndExecRequest("GET", url)

	var serverinfo JiraServerInfo
	err := json.Unmarshal(contents, &serverinfo)
	if err != nil {
		fmt.Println("%s", err)
	}

	return serverinfo
}


func (j *Jira) Statuses() []Status {

	url := j.BaseUrl + j.ApiPath + "/status"
	contents := j.buildAndExecRequest("GET", url)

	var statuses []Status
	err := json.Unmarshal(contents, &statuses)
	if err != nil {
		fmt.Println("%s", err)
	}

	return statuses
}


func (j *Jira) StatusCategories() []StatusCategory {

	url := j.BaseUrl + j.ApiPath + "/statuscategories"
	contents := j.buildAndExecRequest("GET", url)

	var statuscategories []StatusCategory
	err := json.Unmarshal(contents, &statuscategories)
	if err != nil {
		fmt.Println("%s", err)
	}

	return statuscategories
}

