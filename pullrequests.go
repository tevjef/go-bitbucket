package bitbucket

import (
	"encoding/json"
	"strconv"

	"github.com/k0kubun/pp"
)

type PullRequests struct {
	c *Client
}

func (p *PullRequests) Create(po *PullRequestsOptions) (interface{}, error) {
	data := p.buildPullRequestBody(po)
	urlStr := p.c.requestUrl("/repositories/%s/%s/pullrequests/", po.Owner, po.Repo_slug)
	return p.c.execute("POST", urlStr, data)
}

func (p *PullRequests) Update(po *PullRequestsOptions) (interface{}, error) {
	data := p.buildPullRequestBody(po)
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.Repo_slug + "/pullrequests/" + po.Id
	return p.c.execute("PUT", urlStr, data)
}

func (p *PullRequests) Gets(po *PullRequestsOptions) (interface{}, error) {
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.Repo_slug + "/pullrequests/"
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) Get(po *PullRequestsOptions) (interface{}, error) {
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.Repo_slug + "/pullrequests/" + po.Id + "?" + po.Query.Encode()
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) Activities(po *PullRequestsOptions) (interface{}, error) {
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.Repo_slug + "/pullrequests/activity"
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) Activity(po *PullRequestsOptions) (interface{}, error) {
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.Repo_slug + "/pullrequests/" + po.Id + "/activity"
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) Commits(po *PullRequestsOptions) (interface{}, error) {
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.Repo_slug + "/pullrequests/" + po.Id + "/commits"
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) Patch(po *PullRequestsOptions) (interface{}, error) {
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.Repo_slug + "/pullrequests/" + po.Id + "/patch"
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) Diff(po *PullRequestsOptions) (interface{}, error) {
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.Repo_slug + "/pullrequests/" + po.Id + "/diff"
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) Merge(po *PullRequestsOptions) (interface{}, error) {
	data := p.buildPullRequestBody(po)
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.Repo_slug + "/pullrequests/" + po.Id + "/merge"
	return p.c.execute("POST", urlStr, data)
}

func (p *PullRequests) Decline(po *PullRequestsOptions) (interface{}, error) {
	data := p.buildPullRequestBody(po)
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.Repo_slug + "/pullrequests/" + po.Id + "/decline"
	return p.c.execute("POST", urlStr, data)
}

func (p *PullRequests) GetComments(po *PullRequestsOptions) (interface{}, error) {
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.Repo_slug + "/pullrequests/" + po.Id + "/comments/"
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) GetComment(po *PullRequestsOptions) (interface{}, error) {
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.Repo_slug + "/pullrequests/" + po.Id + "/comments/" + po.Comment_id
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) CreateComment(po *PullRequestsOptions) (*Comment, error) {
	data := p.buildCommentBody(po)
	urlStr := GetV1ApiBaseURL() + "/repositories/" + po.Owner + "/" + po.Repo_slug + "/pullrequests/" + po.Id + "/comments/"
	resp, err := p.c.execute("POST", urlStr, data)

	respData, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	comment := Comment{}
	err = json.Unmarshal(respData, &comment)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (p *PullRequests) UpdateComment(po *PullRequestsOptions) (interface{}, error) {
	data := p.buildCommentBody(po)
	urlStr := GetV1ApiBaseURL() + "/repositories/" + po.Owner + "/" + po.Repo_slug + "/pullrequests/" + po.Id + "/comments/" + po.Comment_id
	return p.c.execute("PUT", urlStr, data)
}

func (p *PullRequests) DeleteComment(po *PullRequestsOptions) (interface{}, error) {
	urlStr := GetV1ApiBaseURL() + "/repositories/" + po.Owner + "/" + po.Repo_slug + "/pullrequests/" + po.Id + "/comments/" + po.Comment_id
	return p.c.execute("DELETE", urlStr, "")
}

func (p *PullRequests) CreateTask(po *PullRequestsOptions) (*Task, error) {
	data := p.buildTaskBody(po)
	urlStr := GetInternalApiBaseURL() + "/repositories/" + po.Owner + "/" + po.Repo_slug + "/pullrequests/" + po.Id + "/tasks/"

	resp, err := p.c.execute("POST", urlStr, data)
	if err != nil {
		return nil, err
	}

	respData, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	task := Task{}
	err = json.Unmarshal(respData, &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (p *PullRequests) UpdateTask(po *PullRequestsOptions) (*Task, error) {
	data := p.buildTaskBody(po)
	urlStr := GetInternalApiBaseURL() + "/repositories/" + po.Owner + "/" + po.Repo_slug + "/pullrequests/" + po.Id + "/tasks/" + po.TaskID

	resp, err := p.c.execute("PUT", urlStr, data)
	if err != nil {
		return nil, err
	}

	respData, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	task := Task{}
	err = json.Unmarshal(respData, &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (p *PullRequests) GetTasks(po *PullRequestsOptions) (interface{}, error) {
	urlStr := GetInternalApiBaseURL() + "/repositories/" + po.Owner + "/" + po.Repo_slug + "/pullrequests/" + po.Id + "/tasks/" + "?comment=" + po.Comment_id
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) buildCommentBody(po *PullRequestsOptions) string {
	body := Comment{
		Content: po.CommentContent,
	}

	data, err := json.Marshal(body)
	if err != nil {
		pp.Println(err)
	}

	return string(data)
}

func (p *PullRequests) buildTaskBody(po *PullRequestsOptions) string {
	id, _ := strconv.Atoi(po.Comment_id)
	body := Task{
		Content: TaskContent{
			Raw: po.TaskContent,
		},
		Comment: TaskComment{
			Id: id,
		},
	}

	if po.TaskResolved != nil {
		if *po.TaskResolved {
			body.State = "RESOLVED"
		} else {
			body.State = "UNRESOLVED"
		}
	}

	data, err := json.Marshal(body)
	if err != nil {
		pp.Println(err)
	}

	return string(data)
}

func (p *PullRequests) buildPullRequestBody(po *PullRequestsOptions) string {

	body := map[string]interface{}{}
	body["source"] = map[string]interface{}{}
	body["destination"] = map[string]interface{}{}
	body["reviewers"] = []map[string]string{}
	body["title"] = ""
	body["description"] = ""
	body["message"] = ""
	body["close_source_branch"] = false

	if n := len(po.Reviewers); n > 0 {
		for i, user := range po.Reviewers {
			body["reviewers"].([]map[string]string)[i] = map[string]string{"username": user}
		}
	}

	if po.Source_branch != "" {
		body["source"].(map[string]interface{})["branch"] = map[string]string{"name": po.Source_branch}
	}

	if po.Source_repository != "" {
		body["source"].(map[string]interface{})["repository"] = map[string]interface{}{"full_name": po.Source_repository}
	}

	if po.Destination_branch != "" {
		body["destination"].(map[string]interface{})["branch"] = map[string]interface{}{"name": po.Destination_branch}
	}

	if po.Destination_commit != "" {
		body["destination"].(map[string]interface{})["commit"] = map[string]interface{}{"hash": po.Destination_commit}
	}

	if po.Title != "" {
		body["title"] = po.Title
	}

	if po.Description != "" {
		body["description"] = po.Description
	}

	if po.Message != "" {
		body["message"] = po.Message
	}

	if po.Close_source_branch == true || po.Close_source_branch == false {
		body["close_source_branch"] = po.Close_source_branch
	}

	data, err := json.Marshal(body)
	if err != nil {
		pp.Println(err)
	}

	return string(data)
}
