package main

import (
	"fmt"
	"errors"
)

type Branch struct {
	Name string
	LatestCommit *HashId 
}

type HashId struct {
	Id string
}

type Header struct {
	BranchName string
	ActiveBranch *Branch
}

type Commit struct {
	Id *HashId
	Msg string
	Author string
}

type Commits []Commit
type Branches []Branch

var commits Commits
var branches Branches
var headerPtr Header

func Init(author string) (*Commit, error) {
	if len(author) == 0 {
		return nil, errors.New("Empty parameters passed in")
	}

	
	newRepo := Commit {
		Author: author,
	}

	// for every new repo, defaut master is initialised
	master := Branch {
		Name: "master",
	}

	branches = append(branches, master)

	headerPtr.ActiveBranch = &master
	headerPtr.BranchName = master.Name

	return &newRepo, nil
}

func CheckoutB(name string) {
	newBranch := Branch{
		Name: name,
	}
	// get lastest branch
	latestBranch := branches[len(branches) - 1]
	newBranch.LatestCommit = latestBranch.LatestCommit
	
	branches = append(branches, newBranch)
	fmt.Print("\nnew branch added")

	HeaderPtr2(name)
	
}



func HeaderPtr2(name string)  {
	var newBranch Branch
	for i, branch := range branches {
		if branch.Name == name {
			newBranch = branches[i]	
			// return 
		}
	}


	headerPtr.ActiveBranch = &newBranch
	headerPtr.BranchName = newBranch.Name
}

func (commit *Commit) NCommit (msg, hash string) error{
	if len(msg) == 0 || len(hash) == 0{
		return errors.New("Empty parameters passed in")
	}

	hashId := HashId {
		Id: hash,
	}
	
	commit.Msg = msg
	commit.Id = &hashId
	
	commits = append(commits, *commit)
	// update branch pointer
	branches[0].LatestCommit = &hashId

	fmt.Println("\n[new commit added]")
	return nil
}

// travesre to a previous branch/already existing branch
func Switch(name string) error{
	for i, branch := range branches {
		if branch.Name == name {
			headerPtr.ActiveBranch = &(branches[i])
			headerPtr.BranchName = branches[i].Name

			return nil
		} 
	}


	return errors.New("No branch found")
}

func main() {
	aphex, err := Init("aphex") 
	
	if err!= nil {
		fmt.Println(err)
		panic(err)
	}

	aphex.NCommit("xtal by AphexTwin", "xc_234800u24412")
	aphex.NCommit("stoneinFocus", "sfc_094800s24212")

	CheckoutB("soundcloudId")
	CheckoutB("spotify")

	Switch("master")
	fmt.Printf("\n%+v", headerPtr)
	
}
