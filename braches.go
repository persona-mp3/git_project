package main

import ("fmt"

	"github.com/google/uuid" 
)


type Branch struct {
	Name string
	LatestCommit *HashId
}

type BranchCache []Branch
type CommitCache []Commit

type HashId struct {
	Id string
}

type Commit struct {
	Parent *HashId
	CommitMsg string
	Id HashId
}


type Header struct {
	WorkingBranch *Branch
}


	// var branchCache BranchCache
var commitCache CommitCache
var master Branch  
var branchCache BranchCache

var currDir Header



















func newCommit(msg string) {
	hashId := HashId {
		Id: fmt.Sprintf("m%s", uuid.New().String()), 
	}

	currDir.WorkingBranch = &master
	// fmt.Println("hashId", hashId)

	// branchCache = append(branchCache, master)

	commit := Commit {
		CommitMsg: msg,
		Id: hashId,
	}

	if len(commitCache) < 1 {
		commit.Parent = &hashId
	} else {
		commit.Parent = &(commitCache[len(commitCache) - 1].Id)
	}

	commitCache = append(commitCache, commit)


	updateMasterBranch()


}

 func updateMasterBranch() {
	if len(commitCache) == 0 {
		return
	}
	master.Name = "MASTER"
	latestCommit := commitCache[len(commitCache)-1].Id 
	master.LatestCommit = &latestCommit
}



func NewBranch(branch string) {
	for _, inCache := range branchCache{
		if inCache.Name == branch || master.Name == inCache.Name {
			continue
		}
	}

	// we'll need to transfer to using a map instead of strings instead to just avoid duplicates
	// then we can get what the latest commits these branches have
	branchCache = append(branchCache, master)

	newBranch := Branch {
		Name: branch,
		LatestCommit: master.LatestCommit,
	}

	currDir.WorkingBranch = &newBranch

	branchCache = append(branchCache, newBranch)

	fmt.Println("\nNEW BRANCH CREATED")
	fmt.Println("currently pointing from latest commit", *newBranch.LatestCommit)
	

	fmt.Println("HEADER CURRENTLY POINTING AT BRANCH-->", currDir.WorkingBranch.Name)
}

func  LatestCommit(id int) {
	// if we need to checkout, we need to actually recurse through the commitHistory 	
	// fmt.Println(id)
	fmt.Println("LATEST COMMIT", *master.LatestCommit)
}

func main() {
	newCommit("kagefummi")
	newCommit("darkFantasy")
	newCommit("lumiere")
	newCommit("honwiyomu")
	LatestCommit(5)

	NewBranch("*branching")
	NewBranch("*BUGSS")
	fmt.Println("\nBRANCH CACHE BELOW")
	fmt.Println(branchCache)

	// updateBranch()
}
