package sort

import (
	"fmt"
	"sort"
	"testing"
)

type Student struct {
	Age  int
	Name string
}

func TestSort(t *testing.T) {
	students := []Student{
		{Age: 20, Name: "Alice"},
		{Age: 22, Name: "Bob"},
		{Age: 20, Name: "Eve"},
		{Age: 22, Name: "Alice"},
	}

	sort.Slice(students, func(i, j int) bool {
		if students[i].Age != students[j].Age {
			return students[i].Age > students[j].Age // 按年龄从大到小排序
		}
		return students[i].Name < students[j].Name // 按姓名字典序排序
	})

	for _, student := range students {
		fmt.Printf("Student: Age=%d, Name=%s\n", student.Age, student.Name)
	}
}
