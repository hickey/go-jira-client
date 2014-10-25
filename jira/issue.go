package gojira

import (
	"encoding/json"
	//"encoding/xml"
    "net/url"
    "strconv"
	"fmt"
	"time"
)


type Issue struct {
	Id        string
	Key       string
	Self      string
	Expand    string
	Fields    *IssueFields
    CreatedAt time.Time
}

type IssueList struct {
	Expand     string
	StartAt    int
	MaxResults int
	Total      int
	Issues     []*Issue
	Pagination *Pagination
}

type IssueFields struct {
	IssueType               *IssueType
	Summary                 string
	Description             string
    Status                  *Status
	Reporter                *User
	Assignee                *User
    Creator                 *User
	Project                 *JiraProject
	Created                 string
    Updated                 string
    LastViewed              string
    Environment             string
    Versions                []string
    Components              []string
    Labels                  []string
    Progress                *Progress
    Votes                   *Votes
    Priority                *Priority
    Watches                 *Watches
    DueDate                 string
    Resolution              string
    ResolutionDate          string
    TimeSpent               string
    WorkLog                 *WorkLog
    SubTasks                []*SubTasks
    IssueLinks              []*IssueLink
    WorkRatio               int
    Comment                 *Comment
    //Attachment
    //TimeTracking
    //fixVersions
    
}

type IssueType struct {
	Self                    string
	Id                      string
	Description             string
	IconUrl                 string
	Name                    string
	Subtask                 bool
}

type JiraProject struct {
	Self                    string
	Id                      string
	Key                     string
	Name                    string
	AvatarUrls              map[string]string
}

type Priority struct {
    Self                    string
    IconUrl                 string
    Name                    string
    Id                      string
}

type Progress struct {
    Progress                uint
    Total                   uint
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

type Votes struct {
    Self                    string
    Votes                   uint
    HasVoted                bool
}

type Watches struct {
    Self                    string
    WatchCount              uint
    IsWatching              bool
}

type IssueLink struct {
    Id                      string
    Self                    string
    Type                    *IssueLinkType
    InwardIssue             *IssueLinkRelation
    OutwardIssue            *IssueLinkRelation
}

type IssueLinkType struct {
    Id                      string
    Name                    string
    Inward                  string
    Outward                 string
    Self                    string
}

type IssueLinkRelation struct {
    Id                      string
    Key                     string
    Self                    string
    Fields                  *IssueLinkRelationFields    
}

type IssueLinkRelationFields struct {
    Summary                 string
    Status                  *Status
    Priority                *Priority
    IssueType               *IssueType
}

type WorkLog struct {
    StartAt                 uint
    MaxResults              uint
    Total                   uint
    //WorkLogs
}

type SubTasks struct {
    
}

type Comment struct {
    StartAt                 uint
    MaxResults              uint
    Total                   uint
    Comments                []*CommentEntry
}

type CommentEntry struct {
    Self                    string
    Id                      string
    Body                    string
    Author                  *User
    UpdateAuthor            *User
    Created                 string
    Updated                 string
}


// search an issue by its id
func (j *Jira) Issue(id string) Issue {

	url := j.BaseUrl + j.ApiPath + "/issue/" + id
	contents := j.buildAndExecRequest("GET", url)

	var issue Issue
	err := json.Unmarshal(contents, &issue)
	if err != nil {
		fmt.Println("%s", err)
	}

	return issue
}

// func (j *Jira) AddComment(issue *Issue, comment string) error {
//     var cMap = make(map[string]string)
//     cMap["body"] = comment
//     cJson, err := json.Marshal(cMap)
//     if err != nil {
//         return err
//     }
//     uri := j.BaseUrl + j.ApiPath + "/issue/" + issue.Key + "/comment"
//     body := bytes.NewBuffer(cJson)
//     _, err = j.postJson(uri, body)
//     if err != nil {
//         return err
//     }
//     return nil
// }



// search issues assigned to given user
func (j *Jira) IssuesAssignedTo(user string, maxResults int, startAt int) IssueList {

	url := j.BaseUrl + j.ApiPath + "/search?jql=assignee=\"" + url.QueryEscape(user) + "\"&startAt=" + strconv.Itoa(startAt) + "&maxResults=" + strconv.Itoa(maxResults)
	contents := j.buildAndExecRequest("GET", url)

	var issues IssueList
	err := json.Unmarshal(contents, &issues)
	if err != nil {
		fmt.Println("%s", err)
	}

	for _, issue := range issues.Issues {
    	t, _ := time.Parse(dateLayout, issue.Fields.Created)
    	issue.CreatedAt = t
	}

	pagination := Pagination{
		Total:      issues.Total,
		StartAt:    issues.StartAt,
		MaxResults: issues.MaxResults,
	}
	pagination.Compute()

	issues.Pagination = &pagination

	return issues
}


