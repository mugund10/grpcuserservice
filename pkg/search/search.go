package search

import (
	"fmt"
	"github.com/mugund10/gRPC_userService/internal/model"
)

type Details struct {
	Users []model.User
}
func (d *Details)ByCriteria( criteria string, value any) []model.User {
	var results []model.User
 
	users := d.Users

	for _, user := range users {
		switch criteria {
			
		case "city":
			if user.City == value.(string) {
				results = append(results, user)
			}
		case "phone":
			if user.Phone == value.(int64) {
				results = append(results, user)
			}
		case "married":
			if user.Married == value.(bool) {
				results = append(results, user)
			}
		
		default:
			fmt.Printf("Invalid criteria: %s\n", criteria)
		}
	}

	return results
}
