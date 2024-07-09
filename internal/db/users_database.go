package db

import (
	"github.com/mugund10/grpcuserservice/internal/model"
)

var UsersList []model.User

func init() {
	UsersList = []model.User{
		{Id: 1, Fname: "Rohit Sharma", City: "MUMBAI", Phone: 903248732, Height: 5.2, Married: true},
		{Id: 2, Fname: "Yashasvi Jaiswal", City: "BHADOHI", Phone: 9014313478, Height: 5.3, Married: false},
		{Id: 3, Fname: "Virat Kohli", City: "DELHI", Phone: 9275843489, Height: 6.4, Married: true},
		{Id: 4, Fname: "Rishabh Pant", City: "ROORKEE", Phone: 9323891212, Height: 5.8, Married: false},
		{Id: 5, Fname: "Shubman Gill", City: "FAZILKA", Phone: 9212902389, Height: 5.1, Married: false},
		{Id: 6, Fname: "Jasprit Bumrah", City: "AHMEDABAD", Phone: 9467633467, Height: 6.6, Married: false},
		{Id: 7, Fname: "Ravindra Jadeja", City: "AHMEDABAD", Phone: 9219028912, Height: 5.3, Married: true},
		{Id: 8, Fname: "Kuldeep Yadav", City: "KANPUR", Phone: 9876543223, Height: 5.9, Married: false},
		{Id: 9, Fname: "Ashwin ravichandran", City: "CHENNAI", Phone: 9023347832, Height: 6.2, Married: false},
		{Id: 10, Fname: "Axar Patel", City: "AHMEDABAD", Phone: 9493423245, Height: 6.0, Married: true},
		{Id: 11, Fname: "Hrithik Roshan", City: "MUMBAI", Phone: 9032433732, Height: 5.2, Married: true},
		{Id: 12, Fname: "Shivam Dube", City: "BHADOHI", Phone: 9144313789, Height: 5.3, Married: false},
		{Id: 13, Fname: "ShahRukh Khan", City: "DELHI", Phone: 9278448912, Height: 6.4, Married: true},
		{Id: 14, Fname: "Ajay Devgan", City: "DELHI", Phone: 9322321278, Height: 5.8, Married: true},
		{Id: 15, Fname: "Kapil Dev", City: "FAZILKA", Phone: 9034622342, Height: 5.1, Married: true},
		{Id: 16, Fname: "Navdeep Asija", City: "FAZILKA", Phone: 9464563467, Height: 6.6, Married: true},
		{Id: 17, Fname: "Pradip Baijal", City: "ROORKEE", Phone: 9217832542, Height: 5.3, Married: false},
		{Id: 18, Fname: "Ram Nath Kovind", City: "KANPUR", Phone: 9356324460, Height: 5.9, Married: true},
		{Id: 19, Fname: "Sundar Pichai", City: "CHENNAI", Phone: 9023435832, Height: 6.2, Married: true},
		{Id: 20, Fname: "Lady Andal amma", City: "CHENNAI", Phone: 9432423245, Height: 6.0, Married: true},
	}
}
