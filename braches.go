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

	// var branchCache BranchCache
	var commitCache CommitCache
	var master Branch  

func newCommit(msg string) {
	hashId := HashId {
		Id: fmt.Sprintf("m%s", uuid.New().String()), 
	}
	fmt.Println("hashId", hashId)

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


	master.Name = "[[MASTER]]"
	latestCommit := commitCache[len(commitCache)-1].Id 
	master.LatestCommit = &latestCommit
	// fmt.Println("---- POINTING AT COMMIT ---")
	// fmt.Printf("%+v\n", *master.LatestCommit)


}

func  CheckoutC(id int) {
	
	fmt.Println(id)
	fmt.Println("LATEST COMMIT", *master.LatestCommit)
}


func main() {
	newCommit("kagefummi")
	newCommit("darkFantasy")
	newCommit("lumiere")
	newCommit("honwiyomu")
	defer CheckoutC(5)
	

	// updateBranch()
}
