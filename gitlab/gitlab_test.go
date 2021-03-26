package gitlab

import (
    `fmt`
    "github.com/xanzy/go-gitlab"
    `testing`
)

func initGitlabClient() *gitlab.Client {
    host := ""
    token := "Y-QuueNqkzbLC2sNE6nQ"
    gitlabClient, _ := gitlab.NewClient(token, gitlab.WithBaseURL(host))
    return gitlabClient
}

func TestGitlab_Test(t *testing.T){
    client := initGitlabClient()
    groups, resp, err := client.Groups.ListGroups(&gitlab.ListGroupsOptions{
        ListOptions:          gitlab.ListOptions{
            Page: 1,
            PerPage: 10000,
        },
    })
    if err != nil {
        panic(err)
    }
    fmt.Println(*resp)

    for _, group := range groups {
        if group.ParentID == 0 {
            continue
        }
    }
}
