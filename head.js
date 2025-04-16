const cache = []
const branches = []

class Commit {
  constructor(hash) {
    this.hash = hash
    this.toCache()
  }

  toCache() {
    cache.push(this.hash)
  }
}

class Branch extends Commit {
  constructor(branch, hash) {
    super(hash)
    this.branch = branch
    // this.hash = hash
    this.toBranches()
  }

  toBranches() {
    branches.push(this.branch)
  }
}



function LatestCommit() {
  if (cache.length === 0) {
    console.log("no commits yet")
  }

  const currentNode = cache[(cache.length - 1)]
  console.log( currentNode)
  return currentNode
}


// ROLL BACK HISTORY
function Recurse() {
  if (cache.length === 0) {
    return cache
  }
  if (cache.length <= 1) {
    return cache[(cache.length - 1)]
  }
  return cache.pop()  + Recurse()

}

function Head() {
  if (cache.length === 0) {
    return cache
  }

  return cache[(cache.length - 1)]
}

const commit1 = new Commit("A")
const commitB = new Commit("B")
const commitC = new Commit("C")
const commitD = new Commit("D")
const commitE = new Commit("E")
const feat =  new Branch("feat", "F")
const dev =  new Branch("dev", "G")
const bugFix =  new Branch("bug fix", "H")
console.log(branches)

const currentCommit = LatestCommit()
console.log("latest commit -->", currentCommit)

console.log("current cache -->", cache)
