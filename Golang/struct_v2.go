package main

type Kid struct {
	name string
	age  int
}

type Parent struct {
	name string
	job  string
	age  int
	kids []Kid
}

// Method is function of the structure
func (p Parent) getKids() string {
	lc := ""
	for _, child := range p.kids {
		lc += child.name + ","
	}
	return lc
}

func main() {

	// PHP
	// $parent  = {
	// 		"name": "longHa"
	// 		"age": 23
	// 		"job": "IT"
	// 		"kids" : [{
	// 			"name":  "be A",
	// 			"age": 1
	// 			},
	// 			"name":  "be B",
	// 			"age": 2
	// 			},
	// 		]
	// };

	Parent001 := Parent{
		name: "LongHa",
		job:  "IT PHP",
		age:  23,
		kids: []Kid{
			{name: "be Longha A", age: 1},
			{name: "be Longha B", age: 1},
			{name: "be Longha C", age: 1},
		},
	}

	Parent002 := Parent{
		name: "ChiChi",
		job:  "IT PHP",
		age:  23,
		kids: []Kid{
			{name: "be ChiChi A", age: 1},
			{name: "be ChiChi B", age: 1},
			{name: "be ChiChi C", age: 1},
		},
	}

	// fmt.Println(Parent001)
	// fmt.Println(Parent002)

	parentInClass := []Parent{}
	parentInClass = append(parentInClass, Parent001, Parent002)

	//fmt.Println(parentInClass)

	lc := Parent001.getKids()
	println(lc)

}
