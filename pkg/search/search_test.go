package search_test


import (
	"fmt"
	"testing"
	"github.com/mugund10/gRPC_userService/internal/db"
	"github.com/mugund10/gRPC_userService/pkg/search"
)

func Test(t *testing.T) {
	
	users := db.UsersList
	search := search.Details{users}
	
	// test search by city
	city := "MUMBAI"
	byCity := search.ByCriteria("city", city)
	fmt.Printf("users in %s:\n", city)
	for _, user := range byCity {
		fmt.Println(user.Id, user.Fname, user.Phone)
	}

	// test search by marital status as true
	gotmarried := true
	bymarital := search.ByCriteria("married", gotmarried)
	fmt.Printf("\nMarried users:\n")

	for _, user := range bymarital {
		fmt.Println(user.Id, user.Fname, user.City, user.Married)
	}
// test search by marital status as false 
	gotmarried = false
	bymarital = search.ByCriteria("married", gotmarried)
	fmt.Printf("\nUnmarried users:\n")

	for _, user := range bymarital {
		fmt.Println(user.Id, user.Fname, user.City, user.Married)
	}
}
