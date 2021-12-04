// Challenge : Structs
// Define a Square, which has two fields : center of type point and
// length of type int.
package main

import (
	"fmt"
	"log"
)

// Point is a 2d point
type Point struct {
	X int
	Y int
}

// Move moves the point
func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

// Square
type Square struct {
	Center Point;
	Length int;
}

// Move moves the Square
func (s *Square) Move(dx int, dy int) {
	// move the point at the center of the Square
	s.Center.Move(dx, dy);
}

// Area calculates the area of the square
func (s *Square) Area() int {
	return s.Length * s.Length;
}

// NewSquare creates a new Square from the parameters
func NewSquare(x int, y int, length int) (*Square, error) {
	// check if parameters are provided
	if length <= 0 {
		return nil, fmt.Errorf("Square length must be positive (was %d)", length);
	}

	newSquare := &Square{
		Length: length,
		Center : Point{
			X : x,
			Y : y,
		},
	}

	return newSquare, nil;
}

func main() {
	fmt.Println("Challenge : Structs")

	s, err := NewSquare(45,158,2);

	if err != nil {
		//fmt.Printf("Error : could not create a new Square - %s\n", err);
		//os.Exit(1);
		log.Fatalf("Error : could not create a new Square - %s\n", err)
	}

	fmt.Println("Successfully allocated new Square !")
	fmt.Printf("%+v\n", s);
	fmt.Printf("Area : %d\n", s.Area())
	s.Move(30,50);
	fmt.Printf("%+v\n", s);
	fmt.Printf("Area : %d\n", s.Area())

}
