package bitbucket

import (
	"net/url"
	"time"
)

var apiBaseURL = "https://bitbucket.org/api/2.0"

func GetApiBaseURL() string {
	return apiBaseURL
}

var v1ApiBaseURL = "https://bitbucket.org/!api/1.0"

func GetV1ApiBaseURL() string {
	return v1ApiBaseURL
}

var internalApiBaseURL = "https://bitbucket.org/!api/internal"

func GetInternalApiBaseURL() string {
	return internalApiBaseURL
}

func SetApiBaseURL(urlStr string) {
	apiBaseURL = urlStr
}

type users interface {
	Get(username string) (interface{}, error)
	Followers(username string) (interface{}, error)
	Following(username string) (interface{}, error)
	Repositories(username string) (interface{}, error)
}

type user interface {
	Profile() (interface{}, error)
	Emails() (interface{}, error)
}

type pullrequests interface {
	Create(opt PullRequestsOptions) (interface{}, error)
	Update(opt PullRequestsOptions) (interface{}, error)
	List(opt PullRequestsOptions) (interface{}, error)
	Get(opt PullRequestsOptions) (interface{}, error)
	Activities(opt PullRequestsOptions) (interface{}, error)
	Activity(opt PullRequestsOptions) (interface{}, error)
	Commits(opt PullRequestsOptions) (interface{}, error)
	Patch(opt PullRequestsOptions) (interface{}, error)
	Diff(opt PullRequestsOptions) (interface{}, error)
	Merge(opt PullRequestsOptions) (interface{}, error)
	Decline(opt PullRequestsOptions) (interface{}, error)
}

type repository interface {
	Get(opt RepositoryOptions) (*Repository, error)
	Create(opt RepositoryOptions) (*Repository, error)
	Delete(opt RepositoryOptions) (interface{}, error)
	ListWatchers(opt RepositoryOptions) (interface{}, error)
	ListForks(opt RepositoryOptions) (interface{}, error)
}

type repositories interface {
	ListForAccount(opt RepositoriesOptions) (interface{}, error)
	ListForTeam(opt RepositoriesOptions) (interface{}, error)
	ListPublic() (interface{}, error)
}

type commits interface {
	GetCommits(opt CommitsOptions) (interface{}, error)
	GetCommit(opt CommitsOptions) (interface{}, error)
	GetCommitComments(opt CommitsOptions) (interface{}, error)
	GetCommitComment(opt CommitsOptions) (interface{}, error)
	GetCommitStatus(opt CommitsOptions) (interface{}, error)
	GiveApprove(opt CommitsOptions) (interface{}, error)
	RemoveApprove(opt CommitsOptions) (interface{}, error)
}

type branchrestrictions interface {
	Gets(opt BranchRestrictionsOptions) (interface{}, error)
	Get(opt BranchRestrictionsOptions) (interface{}, error)
	Create(opt BranchRestrictionsOptions) (interface{}, error)
	Update(opt BranchRestrictionsOptions) (interface{}, error)
	Delete(opt BranchRestrictionsOptions) (interface{}, error)
}

type diff interface {
	GetDiff(opt DiffOptions) (interface{}, error)
	GetPatch(opt DiffOptions) (interface{}, error)
}

type webhooks interface {
	Gets(opt WebhooksOptions) (interface{}, error)
	Get(opt WebhooksOptions) (interface{}, error)
	Create(opt WebhooksOptions) (interface{}, error)
	Update(opt WebhooksOptions) (interface{}, error)
	Delete(opt WebhooksOptions) (interface{}, error)
}

type teams interface {
	List(role string) (interface{}, error) // [WIP?] role=[admin|contributor|member]
	Profile(teamname string) (interface{}, error)
	Members(teamname string) (interface{}, error)
	Followers(teamname string) (interface{}, error)
	Following(teamname string) (interface{}, error)
	Repositories(teamname string) (interface{}, error)
}

type RepositoriesOptions struct {
	Owner string `json:"owner"`
	Team  string `json:"team"`
	Role  string `json:"role"` // role=[owner|admin|contributor|member]
}

type RepositoryOptions struct {
	Owner     string `json:"owner"`
	Repo_slug string `json:"repo_slug"`
	Scm       string `json:"scm"`
	//	Name        string `json:"name"`
	Is_private  string `json:"is_private"`
	Description string `json:"description"`
	Fork_policy string `json:"fork_policy"`
	Language    string `json:"language"`
	Has_issues  string `json:"has_issues"`
	Has_wiki    string `json:"has_wiki"`
	Project     string `json:"project"`
}

type PullRequestsOptions struct {
	Id                  string   `json:"id"`
	Comment_id          string   `json:"comment_id"`
	Owner               string   `json:"owner"`
	Repo_slug           string   `json:"repo_slug"`
	Title               string   `json:"title"`
	Description         string   `json:"description"`
	Close_source_branch bool     `json:"close_source_branch"`
	Source_branch       string   `json:"source_branch"`
	Source_repository   string   `json:"source_repository"`
	Destination_branch  string   `json:"destination_branch"`
	Destination_commit  string   `json:"destination_repository"`
	Message             string   `json:"message"`
	Reviewers           []string `json:"reviewers"`
	Query               url.Values
	CommentContent      string
	TaskID              string
	TaskContent         string
	TaskResolved        *bool
}

type CommitsOptions struct {
	Owner       string `json:"owner"`
	Repo_slug   string `json:"repo_slug"`
	Revision    string `json:"revision"`
	Branchortag string `json:"branchortag"`
	Include     string `json:"include"`
	Exclude     string `json:"exclude"`
	Comment_id  string `json:"comment_id"`
}

type BranchRestrictionsOptions struct {
	Owner     string            `json:"owner"`
	Repo_slug string            `json:"repo_slug"`
	Id        string            `json:"id"`
	Groups    map[string]string `json:"groups"`
	Pattern   string            `json:"pattern"`
	Users     []string          `json:"users"`
	Kind      string            `json:"kind"`
	Full_slug string            `json:"full_slug"`
	Name      string            `json:"name"`
	Value     interface{}       `json:"value"`
}

type DiffOptions struct {
	Owner     string `json:"owner"`
	Repo_slug string `json:"repo_slug"`
	Spec      string `json:"spec"`
}

type WebhooksOptions struct {
	Owner       string   `json:"owner"`
	Repo_slug   string   `json:"repo_slug"`
	Uuid        string   `json:"uuid"`
	Description string   `json:"description"`
	Url         string   `json:"url"`
	Active      bool     `json:"active"`
	Events      []string `json:"events"` // EX) {'repo:push','issue:created',..} REF) https://goo.gl/VTj93b
}

type Comment struct {
	DisplayName     string `json:"display_name,omitempty"`
	ParentID        int    `json:"parent_id,omitempty"`
	Content         string `json:"content,omitempty"`
	ContentRendered string `json:"content_rendered,omitempty"`
	UserAvatarURL   string `json:"user_avatar_url,omitempty"`
	UtcCreatedOn    string `json:"utc_created_on,omitempty"`
	Username        string `json:"username,omitempty"`
	Deleted         bool   `json:"deleted,omitempty"`
	CommentID       int    `json:"comment_id,omitempty"`
	Comparespec     string `json:"comparespec,omitempty"`
	LineTo          string `json:"line_to,omitempty"`
	IsSpam          bool   `json:"is_spam,omitempty"`
	PullRequestID   int    `json:"pull_request_id,omitempty"`
	UtcLastUpdated  string `json:"utc_last_updated,omitempty"`
	IsEntityAuthor  bool   `json:"is_entity_author,omitempty"`
	ConvertMarkup   bool   `json:"convert_markup,omitempty"`
	LineFrom        string `json:"line_from,omitempty"`
	DestRev         string `json:"dest_rev,omitempty"`
	Anchor          string `json:"anchor,omitempty"`
	IsRepoOwner     bool   `json:"is_repo_owner,omitempty"`
}

type Task struct {
	Comment   TaskComment `json:"comment,omitempty"`
	Links     Links       `json:"links,omitempty"`
	Creator   Author      `json:"creator,omitempty"`
	CreatedOn time.Time   `json:"created_on,omitempty"`
	Content   TaskContent `json:"content,omitempty"`
	State     string      `json:"state,omitempty"`
	UpdatedOn time.Time   `json:"updated_on,omitempty"`
	ID        int         `json:"id,omitempty"`
}

type TaskComment struct {
	Id    int   `json:"id,omitempty"`
	Links Links `json:"links"`
}

type TaskContent struct {
	Raw    string `json:"raw"`
	Markup string `json:"markup"`
	HTML   string `json:"html"`
}

type Links struct {
	HTML HTML `json:"html"`
	Self Self `json:"self"`
}

type HTML struct {
	Href string `json:"href"`
}

type Self struct {
	Self string `json:"href"`
}

type Author struct {
	DisplayName string `json:"display_name"`
	Type        string `json:"type"`
	Username    string `json:"username"`
	UUID        string `json:"uuid"`
	Links       Links  `json:"links"`
}
