package main

import (
	"fmt"
	"os"

	"github.com/libgit2/git2go"
)

func main() {
	path := "/home/jorge/go/src/github.com/gogit"
	repo, err := git.InitRepository(path, false)
	if err != nil {
		os.Exit(1)
	}

	opts := &git.StatusOptions{}
	opts.Show = git.StatusShowIndexAndWorkdir
	opts.Flags = git.StatusOptIncludeUntracked | git.StatusOptRenamesHeadToIndex | git.StatusOptSortCaseSensitively

	statusList, err := repo.StatusList(opts)

	entry, err := statusList.ByIndex(0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%v\n", entry.IndexToWorkdir)
	fmt.Printf("%v\n", entry.HeadToIndex)
}
