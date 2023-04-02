package main

import "fmt"

type EmployeeInfo struct {
	age        int
	name       string
	job        string
	likes      []string
	friends    []string
	oldCompany []oldCompany
}

type oldCompany struct {
	name      string
	position  string
	joinedAt  string
	seniority int
}

func main() {

	employee := []EmployeeInfo{
		{
			age:     28,
			name:    "Long Ha",
			job:     "Software Engineer",
			likes:   []string{"Motorcycles", "car", "Mechanics Keyboard", "Technical Development"},
			friends: []string{"Long1", "long2", "long3", "long4", "long5", "long6"},
			oldCompany: []oldCompany{
				{
					name:      "s3corp",
					position:  "Senior developer",
					joinedAt:  "2010-11-11",
					seniority: 3,
				},
				{
					name:      "s4corp",
					position:  "Senior developer 2",
					joinedAt:  "2012-12-12",
					seniority: 4,
				},
			},
		},
		{
			age:     29,
			name:    "Long Ha 222",
			job:     "Software Engineer",
			likes:   []string{"Motorcycles", "car", "Mechanics Keyboard", "Technical Development"},
			friends: []string{"Long1", "long2", "long3", "long4", "long5", "long6"},
			oldCompany: []oldCompany{
				{
					name:      "s5corp",
					position:  "Senior developer 2",
					joinedAt:  "2012-12-12",
					seniority: 4,
				},
			},
		},
	}

	// fmt.Println(longha.name)
	// fmt.Println(longha.age)
	// fmt.Println(longha.friends)
	// fmt.Println(longha.oldCompany)

	for _, value := range employee {

		fmt.Println("\nEmployee :" + value.name)

		for _, oldCompany := range value.oldCompany {

			fmt.Println(oldCompany.name)
		}

	}
}
