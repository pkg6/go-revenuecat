package revenuecat

import (
	"fmt"
	"net/http"
	"strings"
)

//https://www.revenuecat.com/reference/list-apps

var (
	getAlistOfAppsPath = "/v2/projects/%s/apps"
	getAnAppPath       = "/v2/projects/%s/apps/%s"
)

type GetAlistOfAppsResponse struct {
	Items    []*AppsItem `json:"items"`
	NextPage string      `json:"next_page"`
	Object   string      `json:"object"`
	Url      string      `json:"url"`
}

type AppsItem struct {
	AppStore  *AppStore `json:"app_store"`
	CreatedAt int64     `json:"created_at"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Object    string    `json:"object"`
	ProjectId string    `json:"project_id"`
	Type      string    `json:"type"`
}

// GetAlistOfApps
//https://www.revenuecat.com/reference/list-apps
func (c *Client) GetAlistOfApps(nextPageOrProjectId string) (resp *GetAlistOfAppsResponse, err error) {
	var path string
	if strings.HasPrefix(nextPageOrProjectId, Host) {
		path = strings.ReplaceAll(nextPageOrProjectId, Host, "")
	} else {
		path = PathQuery(fmt.Sprintf(getAlistOfAppsPath, nextPageOrProjectId), map[string]string{
			"starting_after": "",
			"limit":          "20",
		})
	}
	resp = new(GetAlistOfAppsResponse)
	err = c.call(c.ApiKeyV2,
		http.MethodGet,
		path,
		nil,
		"",
		&resp)
	return
}

type GetAnAppResponse struct {
	AppStore  AppStore `json:"app_store"`
	CreatedAt int64    `json:"created_at"`
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Object    string   `json:"object"`
	ProjectId string   `json:"project_id"`
	Type      string   `json:"type"`
}

// GetAnApp
//https://www.revenuecat.com/reference/get-app
func (c *Client) GetAnApp(projectId, appId string) (resp *GetAnAppResponse, err error) {
	resp = new(GetAnAppResponse)
	err = c.call(c.ApiKeyV2,
		http.MethodGet,
		fmt.Sprintf(getAnAppPath, projectId, appId),
		nil,
		"",
		&resp)
	return
}
