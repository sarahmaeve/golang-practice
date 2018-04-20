package main

import "fmt"

type Grade struct {
	last_name  string
	first_name string
	grade      int
	subject    string
}

// this can prob be done with two shorter statements
// based on dividing the grade
func (g Grade) letter_grade() string {
	switch {
	case g.grade < 60:
		return "F"
	case g.grade >= 60 && g.grade <= 63:
		return "D-"
	case g.grade >= 64 && g.grade <= 66:
		return "D"
	case g.grade >= 67 && g.grade <= 69:
		return "D+"
	case g.grade >= 70 && g.grade <= 73:
		return "C-"
	case g.grade >= 74 && g.grade <= 76:
		return "C"
	case g.grade >= 77 && g.grade <= 79:
		return "C+"
	case g.grade >= 80 && g.grade <= 83:
		return "B-"
	case g.grade >= 84 && g.grade <= 86:
		return "B"
	case g.grade >= 87 && g.grade <= 89:
		return "B+"
	case g.grade >= 90 && g.grade <= 93:
		return "A-"
	case g.grade >= 94 && g.grade <= 96:
		return "A"
	case g.grade >= 97 && g.grade <= 105:
		return "A+"
	}
	return "I"
}

func (g *Grade) change_grade(new_grade int) {
	g.grade = new_grade
}

// TODO: function that shows a table of all subjects (alphabetized)
// and the grades for each student (alpha order by Last name, first name)
// by letter grade and raw score;
// name           History   Chemistry
// Schmoe, Joe    B+ (88)   C (75)

// TODO: read a CSV file and add to map.

func main() {
	// key : integer record ids or similar
	// playing with map of structs
	grades := make(map[int]*Grade)
	grades[1] = &Grade{
		first_name: "Joe",
		last_name:  "Bloggs",
		grade:      82,
	}

	// can update record this way after creation
	grades[1].grade = 90

	// allocate to zero values
	grades[2] = new(Grade)
	// now add record
	grades[2].grade = 70
	grades[2].first_name = "Sam"
	grades[2].last_name = "Wexford"

	// instead of direct manipulation, use the method this time
	grades[2].change_grade(80)
	fmt.Println(grades[2])
	fmt.Println(grades[2].letter_grade())
}
