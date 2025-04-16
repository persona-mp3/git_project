package main

import (
	"fmt"
	"os"
	"github.com/aquasecurity/table"
)


// func cacheFn() {
// 	// userStore := []int
//
// 	users := []int{1, 2, 3, 4, 5, 6}
// 	users = append(users, 8)
// 	fmt.Println(users)
//
// }


type User struct {
	Name string
	Language string
}

type Objects struct {
	Name *string
	User *string
}


type Cache []User
type ObjectDB []Objects


func CacheUser(){
	user := User {
		Name: "theprimeagen",
		Language: "OCaml",
	}

	userTwo := User {
		Name: "tj_derrev",
		Language: "Lua",
	}

	var cache Cache;
	cache = append(cache, user)
	cache = append(cache, userTwo)
	
	var objectDb ObjectDB
	for _, obj := range cache{
		currUser := Objects {
			Name: &obj.Language,
			User: &obj.Name,
		}

		objectDb = append(objectDb, currUser)

	}

	fmt.Println(objectDb)
	// userCheck := objectDb[1]
	// fmt.Println(*userCheck.User)
}

type Brand struct {
	Brand string
	Origin string
	Owners []string
} 

type SpreadSheet struct {
	BrandName string
	MetaData Brand
}

type BrandCache []Brand
var brandCache BrandCache

func Spreadsheet() {
	fmt.Println("SPREADSHEET*")
	table := table.New(os.Stdout)
	// table.SetRowLines(false)
	table.SetHeaders("Brand", "Metadata")
	for _, brand := range brandCache {
		newRow := SpreadSheet {
			BrandName: brand.Brand,
			MetaData: brand,
		}
		table.AddRow(brand.Brand, brand.Origin)

		fmt.Println(newRow)
	}
	table.Render()
}


func  CreateBrand(brand string, origin string, owners []string) Brand{
	newBrand := Brand {
		Brand: brand,
		Origin: origin,
		Owners: owners,
	}

	// CacheBrand(&newBrand)
	brandCache = append(brandCache, newBrand)
	return newBrand
}

// this function has no state as new brand will call the same functino
// and a new strut is created for a new branf
// so what i need to do is make cache global, i asked chatgpt
// and it said to either use Make or New methods in go 
// make initalises a new ma
	


func main(){
	origin := []string{"ken ijima", "mario", "commawear"}
	acneOrigin := []string{
		"DENMA",
		"MAISON MARGEILA",
		"APHEX TWIN",
	}
	vujade := CreateBrand("vujade", "AMERICA", origin)
	acneStudios := CreateBrand("ACNE STUDIOS", "JAPAN", acneOrigin)

	kapitalOrigin := []string{"????unknown???"}
	kapital := CreateBrand("KAPITAL", "JAPAN", kapitalOrigin)
	fmt.Println(vujade, acneStudios, kapital)

	fmt.Println("CACHE STORAGE")
	fmt.Println( brandCache)	
	Spreadsheet()


}
