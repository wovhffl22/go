package main

import "fmt"

type Restuarant struct {
	queue []string
}

// Insert adds a number of elements into array as entered.
func (r *Restuarant) Insert(name ...string) {
	r.queue = append(r.queue, name...)
}

// Delete removes a number of elements from array as entered.
func (r *Restuarant) Delete(num int) {
	r.queue = r.queue[num:]
}

func main() {

	// Create an object.
	myQueue := Restuarant{}

	// A group of people wait in line.
	myQueue.Insert("person A", "person B", "person C", "person D", "person E")

	fmt.Println(myQueue.queue)

	// People who came first went first.
	myQueue.Delete(3) // -3 error

	fmt.Println(myQueue.queue)
}
