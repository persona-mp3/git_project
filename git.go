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
	Name string
	LatestCommit *HashId
}

type GitStore []Git
type CommitCache []Commit
type BranchCache []Branch

var branches BranchCache
var master Branch
var commitCache CommitCache



func Init() *Commit{
	master.Name = "master"
	branches = append(branches, master)

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
	}
	

	fmt.Println("\n[initialised new git repository]")
	return &repo
}



func ShowCommitHistory() {
	table := table.New(os.Stdout)
	fmt.Println("\n---------COMMIT HISTORY-----------\n")
	table.SetHeaders("hashID", "author", "commitMsg", "branch", "commitedAt", "parentCommit")
	
	for _, commit := range commitCache {
		formattedTime := commit.CommitedAt.Format("Jan 2, 2006 3:04 PM")
		parentCommit := commit.Parent.Id
		table.AddRow(commit.HashId.Id, commit.Author, commit.CommitMsg, commit.Master, formattedTime, parentCommit)
	}

	table.Render()
}


func updateMaster() {
	branches[0].LatestCommit = &commitCache[len(commitCache)-1].HashId
}

func updateCommitCache(commit *Commit) {
	commitCache = append(commitCache, *commit)
	updateGitStore()
}

func (commit *Commit) Commit(msg string) {
	hashId := HashId {
		Id: uuid.New().String(),
	}

	commitTime := time.Now()

	commit.HashId = hashId
	commit.CommitedAt = commitTime
	commit.CommitMsg = msg
	
	if len(commitCache) < 1 {
		commit.Parent = &hashId 
		// assuming master is the firstBranch
		branches[0].LatestCommit = &hashId
		fmt.Println("\n[new commit added ...]")
		updateCommitCache(commit)
		// return *commit
		return
	}
	
	commit.Parent = &commitCache[len(commitCache)-1].HashId
	updateCommitCache(commit)
	updateMaster()
	fmt.Println("\n[new commit added ...]")
	// return *commit
	return
}

func (commit *Commit) CheckoutB(name string) {
	newBranch := Branch {
		Name: name,
		LatestCommit: (branches[len(branches)-1].LatestCommit),
	}

	branches = append(branches, newBranch)
	fmt.Println("[new branch added]")
	fmt.Printf("%+v\n", branches)

}


var gitStore GitStore
func updateGitStore() {
	if len(commitCache) == 0  {
		return
	}

	currCommit := commitCache[len(commitCache)-1]
	newRow := Git {
		Key: &currCommit.HashId,
		Object : &currCommit,
	}

	gitStore = append(gitStore, newRow)
	// fmt.Println("[git store updated]")
	// fmt.Println(gitStore)
	return
	
}

func PrintDb() {
	table := table.New(os.Stdout)
	table.SetHeaders("#id", "Commit")
	
	fmt.Println("\n-----------RENDERING GIT STORE------------\n")

	for _, commit := range gitStore {
		hashId := *commit.Key
		commitMsg := commit.Object.CommitMsg 
		table.AddRow(hashId.Id, commitMsg)
	}
	table.Render()
}




func main() {
	
	repo := Init()

	repo.Commit("inital commit")
	repo.Commit("added new structs:")
	repo.Commit("working on gitStore")

	// getId := repo.Commit("feat: configured branching").HashId
	// fmt.Println(getId)
	// fmt.Println(*branches[0].LatestCommit)
	// [lines 10 and 13 must return the same value]
	
	ShowCommitHistory()
	PrintDb()
}
