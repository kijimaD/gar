package gh

import "fmt"

type CallClient struct {
	API clientI
}

func (c *CallClient) PRGitHashs() {
	pr := c.API.PRDetail()
	curBranch := *pr.Head.Ref
	baseBranch := *pr.Base.Ref
	fmt.Println(curBranch, baseBranch)
}

func (c *CallClient) PRCommits() {
	commits := c.API.PRCommits()

	for _, c := range commits {
		fmt.Printf("%v\n", *c.Commit.Message)
	}
}
