package main

import "fmt"

func main() {
	// 1. Check this generic method will remove integer element by targeting index
	integerInput := []int{1, 2, 3, 4, 5}
	integerUpdated := DeleteElementByIndex(integerInput, 2)
	fmt.Println(integerUpdated) // Output: [1 2 4 5]

	// 2. Check this generic method will remove string element by targeting index
	stringInput := []string{"John Wick", "Jack Sparrow", "Iron Man", "Bat Man"}
	stringUpdated := DeleteElementByIndex(stringInput, 3)
	fmt.Println(stringUpdated) // Output: ["John Wick", "Jack Sparrow", "Iron Man"]

	// 3. It is consuming the DeepEqual which is a recursive function,
	// it will remove the elements which shared same values for the objects.
	var book1 Book
	var book2 Book
	var book3 Book

	book1.id = 1
	book1.subject = "computer science"
	book1.author = "Jack Sparrow"
	book1.title = "How to develop micro services"

	book2.id = 1
	book2.subject = "computer science"
	book2.author = "Jack Sparrow"
	book2.title = "How to develop micro services"

	book3.id = 3
	book3.subject = "social"
	book3.author = "John Wick"
	book3.title = "How to improve your personality"

	booksInput := []Book{book1, book2, book3}
	bookListResult := DeleteElementByValue(booksInput, book1)
	fmt.Println(bookListResult) // Output: [{How to improve your personality John Wick social 3}]

	// 4. Used for removing the comparable elements by value without using DeepEqual recursive function.
	names := []string{"Gary", "Jason", "Daniel", "David", "Bob", "Alex", "Bob"}
	updatedResult := DeleteElementByValueForComparableElements(names, "Bob")
	fmt.Println(updatedResult)
	// Output: [Gary Jason Daniel David Alex]

	// TODO: 抱歉我没太懂什么是缩容
}

type Book struct {
	title   string
	author  string
	subject string
	id      int
}
