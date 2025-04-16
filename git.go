package main

import (
	"fmt"
	"time"

	"os"
	"github.com/aquasecurity/table"
	"github.com/google/uuid" 
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
	table.SetHeaders("hashID", "author", "commitMsg", "branch", "commitedAt", "parentCommit")
	
	for _, commit := range commitCache {
		formattedTime := commit.CommitedAt.Format("Jan 2, 2006 3:04 PM")
		parentCommit := commit.Parent.Id
		table.AddRow(commit.HashId.Id, commit.Author, commit.CommitMsg, commit.Master, formattedTime, parentCommit)
	}

	table.Render()
}


func (commit *Commit) Commit(msg string) {
	hashId := HashId {
		Id: fmt.Sprintf("m%s", uuid.New().String()),
	}

	// we need to first check cache history if theres any one before it 
	// [node1, node2, node3]
	// node4??
	// 
	// if len(cache) === 0; commit.Parent = nil
	// parentCommit = cache[(len(cache) -1)]
	// commit.Parent = &parentCommit
	// parent := commit.Parent
	// var hashId2 Commit
	commitTime := time.Now()

	if len(commitCache) == 0 {
		commit.Parent = &hashId 
		commit.HashId = hashId
		commit.CommitMsg = msg
		commit.CommitedAt = commitTime
	} else {
		prevCommit := commitCache[(len( commitCache) -1) ]
		// fmt.Println("PREVIOUS COMMIT?????")
		// fmt.Printf("%+v\n",prevCommit)
		parentId := prevCommit.HashId
		commit.Parent = &parentId
		commit.HashId = hashId
		commit.CommitMsg = msg
		commit.CommitedAt = commitTime
	}

	commitCache = append(commitCache, *commit)
	fmt.Println("[new commit added]")
	// fmt.Println("COMMIT CACHE")
	// fmt.Println(commitCache)
}












func main() {
	repo := Init()
	repo.Commit("inital commit")
	repo.Commit("added new structs:")
	repo.Commit("cache as global var:")
	ShowCommitHistory()
	// fmt.Printf("%+v\n", repo)
}
