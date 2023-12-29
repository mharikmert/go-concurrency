package main

import (
	"fmt"
	"sync"
)

// func print(s string) {
// 	fmt.Println(s)
// }

func printV2(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(s)

}

func main() {

	// // GO KEYWORD AND ROUTINES
	// print("This is the first thing to be printed!") // what this go keyword does -> tells the compiler run this in its own routine

	// // runs in the main routine. (main function is also a routine.)
	// go print("This is the first thing to be printed with the go keyword!")

	// // when we run these two lines above, we see only the second statement to be printed.
	// // what happens is the main routine is not waiting the separate routine started in line 14. (should wait it some how)

	// // another routine (there are 3 routines with the main routine now.)
	// go print("This is the second thing to be printed with the go keyword!")
	// print("This is the second thing to be printed!")

	// time.Sleep(time.Second * 1)
	// // when we run these lines above with sleeping for a second, we see all the statements to be printed. (Not just wait the routine in line 23, also wait 14 this time.)
	// // time.sleep is the worst way of waiting a routine btw.

	// WAIT GROUPS
	strings := []string{
		"first",
		"second",
		"third",
		"fourth",
		"fifth",
	}

	// create a wait group
	var wg sync.WaitGroup

	// declare how many things the wait group should wait for -> len(strings) for this case
	wg.Add(len(strings))

	// iterate over the strings by spinning up a new routine for each string
	for i, s := range strings {
		go printV2(fmt.Sprintf("%d: %s", i+1, s), &wg)
	}

	// wait for the routines until they're done.
	wg.Wait()

	// printV2("this is the first thing with the v2", &wg)

}
