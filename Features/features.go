package features

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name   string
	RollNo string
	Marks  int
}

type Class struct {
	Students []Student
}

func (c *Class) NewClass() {

	// fmt.Println(entry)
	var studentsArray []Student
	file, err := os.Open("db.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		marks, _ := strconv.Atoi(fields[2])

		entry := Student{
			Name:   fields[0],
			RollNo: fields[1],
			Marks:  marks,
		}

		studentsArray = append(studentsArray, entry)

	}

	// fmt.Println(studentsArray)
	c.Students = studentsArray

}

func (C *Class) writeToFile() {
	// convert the data

	var data []string

	for _, v := range C.Students {
		line := fmt.Sprintf("%v %v %v", v.Name, v.RollNo, v.Marks)
		data = append(data, line)
	}
	err := os.WriteFile("db.txt", []byte(strings.Join(data, "\n")), 644)
	if err != nil {
		fmt.Println(err.Error())
	}

}

// CRUD: Create Read Update and Delete

func (c *Class) AddStudent() {

	fmt.Println("Enter your first name")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Enter your roll no")
	var roll string
	fmt.Scanln(&roll)

	fmt.Println("Enter your marks")
	var marks int
	fmt.Scanln(&marks)

	entry := Student{
		Name:   name,
		RollNo: roll,
		Marks:  marks,
	}

	c.Students = append(c.Students, entry)

	c.writeToFile()

	fmt.Println("Student Added Successfully")
	fmt.Println()

	//Update this data in the db.txt
}

func (c *Class) ShowStudents() {
	fmt.Println()
	for _, student := range c.Students {
		name, roll, marks := student.Name, student.RollNo, student.Marks
		fmt.Printf("Name: %s\nRoll No: %s\nMarks: %d\n", name, roll, marks)
		fmt.Println()
	}

	c.writeToFile()

}

func (c *Class) UpdateStudent() {

	var roll string
	fmt.Println("Enter the roll no")
	fmt.Scanln(&roll)
	var i int
	for idx, student := range c.Students {
		if student.RollNo == roll {
			i = idx
			break
		}
	}

	fmt.Println("Enter your first name")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Enter your roll no")
	var rollNo string
	fmt.Scanln(&rollNo)

	fmt.Println("Enter your marks")
	var marks int
	fmt.Scanln(&marks)

	c.Students[i] = Student{
		Name:   name,
		RollNo: rollNo,
		Marks:  marks,
	}

	fmt.Println("Data updated successfully")
	fmt.Println()

	c.writeToFile()

	//Update it in db.txt
}

func (c *Class) DeleteStudent() {
	var roll string
	fmt.Println("Enter the roll no")
	fmt.Scanln(&roll)
	var i int
	for idx, student := range c.Students {
		if student.RollNo == roll {
			i = idx
			break
		}
	}

	dummyArr := c.Students

	dummyArr = append(dummyArr[:i], dummyArr[i+1:]...)
	c.Students = dummyArr
	fmt.Println("Data deleted successfully")
	fmt.Println()

	c.writeToFile()

	//Update the data in db.txt

}

func (C *Class) GetStart() {
	// sum ...
	sum := C.sum()
	// avg ..
	avg := C.average()
	//highes
	highest := C.highest()
	//lowest
	lowest := C.lowest()

	fmt.Println()
	fmt.Printf("\nsum is: %d\navg marks of all students is: %f\nhighest marks is: %d\nlowest marks is: %d\n", sum, avg, highest, lowest)
	fmt.Println()
	C.writeToFile()

}

func (C *Class) sum() int {

	sum := 0
	for _, v := range C.Students {
		sum += v.Marks
	}

	return sum
}

func (C *Class) average() float32 {

	return float32(C.sum()) / float32(len(C.Students))
}

func (C *Class) highest() int {

	highest := 0
	for _, v := range C.Students {
		highest = max(highest, v.Marks)
	}

	return highest
}

func (C *Class) lowest() int {

	lowest := 500
	for _, v := range C.Students {
		lowest = min(lowest, v.Marks)
	}

	return lowest
}
