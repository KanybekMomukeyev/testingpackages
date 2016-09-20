package controllers

import (
	//"regexp/syntax"
	"fmt"
	"encoding/json"
)

type Person struct {
	First string
	Last string
	Age int
}

type Saiyan struct {
	Name string
	Power int
	Father *Saiyan
}

type Saiyan2 struct {
	Name int
	Power int
}



func Super(s *Person) {
	s.Age += 10000
}

func Super2(s Person) {
	s.Age += 10000
}

func NewPerson(first string, last string, age int) *Person {
	return &Person{
		First: first,
		Last: last,
		Age: age,
	}
}

func NewPerson2(first string, last string, age int) Person {
	return Person{
		First: first,
		Last: last,
		Age: age,
	}
}
type person struct {
	name string
	age int
}

func SomeFunctionCalled2(k int) int {

	fmt.Println("")
	fmt.Println("")
	fmt.Println("STARTING")


	var P person  // p is person type
	P.name = "Astaxie"  // assign "Astaxie" to the field 'name' of p
	P.age = 25  // assign 25 to field 'age' of p
	fmt.Printf("The person's name is %s\n", P.name)


	goku1000 := new(Saiyan2)
	fmt.Println(goku1000 )

	goku2000 := &Saiyan2{}
	fmt.Println(goku2000 )


	gokurbek := &Saiyan {
		Name: "goku",
		Power: 9000,
	}

	gohannek := &Saiyan {
		Name: "Gohan",
		Power: 1000,
		Father: &Saiyan {
			Name: "Goku",
			Power: 9001,
			Father: nil,
		},
	}

	fmt.Println(gokurbek)
	fmt.Println(gohannek)


	per1 := NewPerson("koke","boke",1000)
	per1.First = "10"
	per1.Last = "10"
	per1.Age = 0

	per2 := NewPerson2("koke","boke",1000)

	fmt.Println(*per1)
	fmt.Println(per1)
	fmt.Println(per2)

	goku := Person{"Goku","terr", 9000}
	bs, _ := json.Marshal(goku)
	fmt.Println(string(bs))

	var p12 Person
	json.Unmarshal(bs, &p12)
	fmt.Println(p12)

	fmt.Println(goku)
	Super(&goku)
	fmt.Println(goku.Age)
	Super2(goku)
	fmt.Println(goku.Age)

	pp := make([]Person, 3)
	p11:= Person{"koke","boke",23}
	pp = append(pp, p11)
	fmt.Println("pp :", pp)


	dict := map[string]int {"foo" : 1, "bar" : 2}
	value, ok := dict["baz"]

	if ok {
		fmt.Println("value: ", value)
	} else {
		fmt.Println("key not found")
	}

	if value, ok := dict["baz"]; ok {
		fmt.Println("value: ", value)
	} else {
		fmt.Println("key not found")
	}

	s := make([]string, 3)
	s = append(s, "Hello world")

	s1 := []string{}
	fmt.Println("emp1:", s)
	fmt.Println("emp2:", s1)


	mm := map[string]int{
		"kev":1,
		"kevev":2,
	}

	for kev, value := range mm {
		fmt.Println(kev, "----", value)
	}
	
	fmt.Println("mm:", mm)

	m := map[string]int{}
	//m := make(map[string]int)

	fmt.Println("map:", m)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	return 1123
}