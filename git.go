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
	

	fmt.Println("[initialised new git repository]")
	return &repo
}


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


func updateMaster() {
	master.LatestCommit = &commitCache[len(commitCache)-1].HashId
}

func updateCommitCache(commit *Commit) {
	commitCache = append(commitCache, *commit)
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
		master.LatestCommit = &hashId
		updateCommitCache(commit)
		return
	}
	
	commit.Parent = &commitCache[len(commitCache)-1].HashId
	updateCommitCache(commit)
	updateMaster()
	fmt.Println("\n[new commit addedd...]\n")
	return
}


var gitStore GitStore

func GitDatabase() {
	// iterates through the cache
	// and then stores data in key value pairs
	// where the key is the hashId of the commit 
	// and the value is the actual commitMetaData
		
	if len(commitCache) == 0 {
		return 
	}

	fmt.Println("GIT STORE")

	for i, _ := range commitCache {
		newRow := Git {
			Key: &commitCache[i].HashId,
			Object : &commitCache[i],
		}
		// this is supposed to be immutable btw
		gitStore = append(gitStore, newRow)
	}
		fmt.Println(gitStore)
}


func SeeGitStore() {
	table := table.New(os.Stdout)
	table.SetHeaders("Key/HashCodes", "Objects/Commits")

	for _, row := range gitStore {

		key := row.Key.Id	
		commitMsg := row.Object.CommitMsg
		table.AddRow(key, commitMsg)
	}

	table.Render()
}



func main() {
	repo := Init()
	repo.Commit("inital commit")
	repo.Commit("added new structs:")
	repo.Commit("feat: caching fixed")
	
	fmt.Println(*master.LatestCommit)
	ShowCommitHistory()
	// GitDatabase()
}
