package main

import "fmt"

// towerOfHanoi function solves the Tower of Hanoi puzzle
func towerOfHanoi(n int, source, destination, auxiliary string) {
	if n == 1 {
		fmt.Printf("Move disk 1 from %s to %s\n", source, destination)
		return
	}

	towerOfHanoi(n-1, source, auxiliary, destination)
	fmt.Printf("Move disk %d from %s to %s\n", n, source, destination)
	towerOfHanoi(n-1, auxiliary, destination, source)
}

func main() {
	n := 3                         // Number of disks
	towerOfHanoi(n, "A", "C", "B") // A, B and C are names of rods
}
