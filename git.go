package main

import (
	"fmt"
	"time"

	"os"
	"github.com/aquasecurity/table"
)

type Git struct {
	Key *HashId
	Object *Commit
}



type Commit struct {
	Parent *HashId
	HashId HashId
	Author string
	Snapshot *Tree
	CommitMsg string
	Content Tree
	Master string
	ExternalBranch *Branch
	CommitedAt time.Time 
}

type Tree struct {
	Files []string
}

type HashId struct {
	Id string
}

type Branch struct {
	ParentCommit *Commit 
}

type GitDB []Git
type CommitCache []Commit


func Init() *Commit{
	// now := time.Now()
	hashId := HashId {
		Id: "[no commits yet]",
	}

	repo := Commit {
		Parent: nil,
		HashId: hashId, 
		Author: "personamp3",
		Snapshot: nil,
		CommitMsg:"[no commits yet]" ,
		Master: "[master]",
		ExternalBranch: nil,
		// CommitedAt: time.Time, 
	}
	

	fmt.Println("[initialised new git repository]")
	return &repo
}



var commitCache CommitCache

func ShowCommitHistory() {
	table := table.New(os.Stdout)
	table.SetHeaders("hashID", "author", "commitMsg", "branch", "commitedAt")
	
	for _, commit := range commitCache {
		formattedTime := commit.CommitedAt.Format("Jan 2, 2006 3:04 PM")
		table.AddRow(commit.HashId.Id, commit.Author, commit.CommitMsg, commit.Master, formattedTime)
	}

	table.Render()
}


func (commit *Commit) Commit(msg string) {
	hashId := HashId {
		Id: fmt.Sprintf("m%07d", time.Now().UnixNano() % 10000000),
	}

	commitTime := time.Now()
	commit.HashId = hashId
	commit.CommitMsg = msg
	commit.CommitedAt = commitTime
	// commit.CommitedAt = commitTime.Format("Jan 2, 2006 3:04 PM")

	commitCache = append(commitCache, *commit)
	fmt.Println("[new commit added]")
	fmt.Println("COMMIT CACHE")
	// fmt.Println(commitCache)
}












func main() {
	repo := Init()
	repo.Commit("inital commit")
	repo.Commit("added new structs:")
	repo.Commit("cache as global var:")
	ShowCommitHistory()
	fmt.Printf("%+v\n", repo)
}
