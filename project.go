package revenuecat

import (
	"net/http"
	"strings"
)

var (
	getAListOfProjectsPath = "/v2/projects"
)

type GetAListOfProjectsResponse struct {
	Items []*GetAListOfProjectsItem `json:"items"`
	//https://api.revnuecat.com/v2/projects?starting_after=proj82bc1db5&limit=1
	NextPage string `json:"next_page"`
	Object   string `json:"object"`
	Url      string `json:"url"`
}
type GetAListOfProjectsItem struct {
	CreatedAt int64  `json:"created_at"`
	Id        string `json:"id"`
	Name      string `json:"name"`
	Object    string `json:"object"`
}

// GetAListOfProjects
// nextPage
//case: ""
//case: https://api.revnuecat.com/v2/projects?starting_after=proj82bc1db5&limit=1
//https://www.revenuecat.com/reference/list-projects
func (c *Client) GetAListOfProjects(nextPage string) (resp *GetAListOfProjectsResponse, err error) {
	var path string
	if strings.HasPrefix(nextPage, Host) {
		path = strings.ReplaceAll(nextPage, Host, "")
	} else {
		path = PathQuery(getAListOfProjectsPath, map[string]string{
			"starting_after": "",
			"limit":          "20",
		})
	}
	resp = new(GetAListOfProjectsResponse)
	err = c.call(c.ApiKeyV2,
		http.MethodGet,
		path,
		nil,
		"",
		&resp)
	return
}
