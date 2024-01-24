package revenuecat

import (
	"net/http"
	"strings"
)

var (
	getAListOfProjectsPath = "/v2/projects"
)

type GetAListOfProjectsResponse struct {
	Items []*GetAListOfProjectsItem `json:"items,omitempty"`
	//https://api.revnuecat.com/v2/projects?starting_after=proj82bc1db5&limit=1
	NextPage string `json:"next_page,omitempty"`
	Object   string `json:"object,omitempty"`
	Url      string `json:"url,omitempty"`
}
type GetAListOfProjectsItem struct {
	CreatedAt int64  `json:"created_at,omitempty"`
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Object    string `json:"object,omitempty"`
}

// GetAListOfProjects
// nextPage
//case: ""
//case: https://api.revnuecat.com/v2/projects?starting_after=proj82bc1db5&limit=1
//https://www.revenuecat.com/reference/list-projects
func (c *Client) GetAListOfProjects(nextPage string) (resp *GetAListOfProjectsResponse, err error) {
	var path string
	if strings.HasPrefix(nextPage, defaultBaseURL) {
		path = strings.ReplaceAll(nextPage, defaultBaseURL, "")
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
