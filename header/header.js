// lets make the header point to an a branch in branches
// well make it a class that has a method for checking the branches and use the branch
//
const branches = [] 
const commits = []

class Commit {
  static id = 1;

  constructor( msg) {
    this.id = Commit.id++
    this.commit = msg;
    this.hash = this.hashFn()
    this.parent = Commit.parent()

    Commit.updateCommits(this)
  }

  hashFn() {
    let chars = "!abcdefghijklmnopqrstuvABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890*#"
    let hashId = ""

    for (let i=0; i <=12; i++) {
      hashId += chars.charAt(Math.floor(Math.random()*chars.length))
    }

    return hashId 
  }


  static parent() {
    if (commits.length === 0) {
      return null
    }

    return commits[commits.length -1].hash
  }

  static updateCommits(commit) {
    commits.push(commit)
    console.log("\n---new commit added---")
    return commit
  }
}



class Branch {
  constructor(name) {
    this.branch = Branch.updateBranches(name)
  }
  
  static updateBranches(name) {
    branches.push(name)
    console.log("[---updated branch---]")
    
    return name
  }
}


class Header {
  constructor(header, branch) {
    this.header = header
    this.branch = branch
  }

  checkout(branch) {
    if (branches.length === 0) {
      return
    }

    for (let i=0; i < branches.length; i++) {
      if (branches[i] === branch) {
        this.branch = branch
        return branch
      } else {
        this.branch = null
        return "no branch found"
      }
    }
  }
}





const goku = new Commit("goku @kingkai")
const vegeta = new Commit("saibamen &&nappa")
const yamcha = new Commit("yamcha &&meme")


console.log(commits)
const kory = new Branch("kory") 
const master = new Branch("master")
console.log(branches)
