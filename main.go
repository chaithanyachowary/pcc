// // package main

// // import (
// // 	"fmt"
// // 	"runtime"
// // )

// // func main() {
// // 	fmt.Println("Starting")
// // 	fmt.Println("hello world")
// // 	fmt.Println(runtime.GOOS) //it will print the Operating system version
// // }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// Get the current day of the week
// 	today := time.Now().Weekday()
// 	fmt.Println(time.Now().AddDate(1999,03,25)) // in this line we added 1999 years , 3 months and 25 days to the current date

// 	// Check if today is Saturday or Sunday (weekend)
// 	isWeekend := today == time.Saturday || today == time.Sunday

// 	if isWeekend {
// 		fmt.Println("It's the weekend!")
// 	} else {
// 		fmt.Println("It's a weekday.")
// 	}
// 	var a = fmt.Println // alias for printing where we can use a instead of fmt.Println
// 	b := 10
// 	a(b)
// }

package main

import (
	"fmt"
	// "time"
	// "math"
	"sync"
	//    "time"
)

// func someTask(id int, data chan int) {
//    for taskId := range data {
//       time.Sleep(2 * time.Second)
//       fmt.Printf("Worker: %d executed Task %d\n", id, taskId)
//    }
// }

// func main() {
//    // Creating a channel
//    channel := make(chan int)

//    // Creating 10.000 workers to execute the task
//    for i := 0; i < 100; i++ {
//       go someTask(i, channel)
//    }

//    // Filling channel with 100.000 numbers to be executed
//    for i := 0; i < 100; i++ {
//       channel <- i
//    }

// }

// type shape interface{
// 	area() float64
// 	perimeter() float64
// }

// type rectangle struct{
// 	length float64
// 	breadth float64
// }

// type circle struct{
// 	radius float64
// }

// func (c circle) area()float64{
// 	return math.Pi*c.radius*c.radius
// }

// func (r rectangle) area()float64{
// 	return r.length*r.breadth
// }

// func (c circle)perimeter()float64{
// 	return 2*math.Pi*c.radius
// }

// func (r rectangle)perimeter() float64{
// 	return 2*(r.length+r.breadth)
// }

// func main(){
// c := circle{2.5}
// r := rectangle{4.5,3}
// fmt.Printf("The perimeter of rectangle is %f \n",shape.perimeter(r))
// fmt.Printf("The perimeter of circle is %f \n",shape.perimeter(c))
// fmt.Printf("The perimeter of rectangle is %f \n",shape.area(r))
// fmt.Printf("The perimeter of circle is %f \n",shape.area(c))
// var a interface{} = 23
// fmt.Printf("%T \n",a)
// a = "34"
// fmt.Printf("%T",a)
// }

func print(wg *sync.WaitGroup, a int) {
	defer wg.Done()
	fmt.Println(a)
}
func main() {
	var wg sync.WaitGroup

	wg.Add(30)

	for i := 1; i <= 10; i++ {
		go print(&wg, i)
		//  time.Sleep(time.Millisecond)
	}
	for i := 1; i <= 20; i++ {
		go print(&wg, i)
		//  time.Sleep(time.Millisecond)
	}
	// time.Sleep(time.Millisecond)
	wg.Wait()
	fmt.Println("loop is terminated ")
}

// without sync and time packages we can print the numbers using  for loop and ananymous function but they print randomly not in an order.
