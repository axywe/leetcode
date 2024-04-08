package main

func countStudents(students []int, sandwiches []int) int {
	students = append(students, 2)
	flag := 0
	for flag != 2 && len(sandwiches) != 0 {
		if students[0] == sandwiches[0] {
			flag = 0
			students = students[1:]
			sandwiches = sandwiches[1:]
		} else if students[0] == 2 {
			if flag == 1 {
				flag = 2
			} else {
				flag = 1
			}
			var student int = students[0]
			students = students[1:]
			students = append(students, student)
		} else if students[0] != sandwiches[0] {
			var student int = students[0]
			students = students[1:]
			students = append(students, student)
		}
	}
	return len(students) - 1
}
