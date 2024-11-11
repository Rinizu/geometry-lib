package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/Rinizu/geometry-lib/model"
)

type Shape interface {
	Area() float32
	Perimeter() float32
}

func printShapeResult(shape Shape) {
	fmt.Printf("Area of shape is: %.2f\n", shape.Area())
	fmt.Printf("Perimeter of shape is: %.2f\n", shape.Perimeter())
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(`Chooso a shape:
	1. Rectangle
	2. Square
	3. Triangle
	4. Circle
	5. Parallelogram
	`)
	scanner.Scan()
	shapeType := scanner.Text()
	shapeChoice, err := strconv.Atoi(shapeType)
	if err != nil {
		fmt.Println("Invalid input for shape: ", err)
		return
	}

	var shape Shape

	switch shapeChoice {
	case 1:
		fmt.Println("Enter width of rectangle: ")
		scanner.Scan()
		widthInput := scanner.Text()
		width, err := strconv.ParseFloat(widthInput, 32)
		if err != nil {
			fmt.Println("Invalid input for width: ", err)
			return
		}

		fmt.Println("Enter length of rectangle: ")
		scanner.Scan()
		lengthInput := scanner.Text()
		length, err := strconv.ParseFloat(lengthInput, 32)
		if err != nil {
			fmt.Println("Invalid input for length: ", err)
			return
		}

		rectangle := model.Rectangle{Width: float32(width), Length: float32(length)}
		shape = &rectangle

	case 2:
		fmt.Println("Enter side of square: ")
		scanner.Scan()
		sideInput := scanner.Text()
		side, err := strconv.ParseFloat(sideInput, 32)
		if err != nil {
			fmt.Println("Invalid input for side: ", err)
			return
		}

		square := model.Square{Side: float32(side)}
		shape = &square

	case 3:
		fmt.Println("Enter base of triangle: ")
		scanner.Scan()
		baseInput := scanner.Text()
		base, err := strconv.ParseFloat(baseInput, 32)
		if err != nil {
			fmt.Println("Invalid input for base: ", err)
			return
		}

		fmt.Println("Enter height of triangle: ")
		scanner.Scan()
		heightInput := scanner.Text()
		height, err := strconv.ParseFloat(heightInput, 32)
		if err != nil {
			fmt.Println("Invalid input for height: ", err)
			return
		}

		fmt.Println("Enter left side length of triangle: ")
		scanner.Scan()
		leftSideInput := scanner.Text()
		leftSide, err := strconv.ParseFloat(leftSideInput, 32)
		if err != nil {
			fmt.Println("Invalid input for left side: ", err)
			return
		}

		fmt.Println("Enter right side length of triangle: ")
		scanner.Scan()
		rightSideInput := scanner.Text()
		rightSide, err := strconv.ParseFloat(rightSideInput, 32)
		if err != nil {
			fmt.Println("Invalid input for right side: ", err)
			return
		}

		triangle := model.Triangle{Base: float32(base), Height: float32(height), SideLeft: float32(leftSide), SideRight: float32(rightSide)}
		shape = &triangle

	case 4:
		fmt.Println("Enter radius of circle: ")
		scanner.Scan()
		radiusInput := scanner.Text()
		radius, err := strconv.ParseFloat(radiusInput, 32)
		if err != nil {
			fmt.Println("Invalid input for radius: ", err)
			return
		}

		circle := model.Circle{Radius: float32(radius)}
		shape = &circle

	case 5:
		fmt.Println("Enter base of parallelogram: ")
		scanner.Scan()
		baseInput := scanner.Text()
		base, err := strconv.ParseFloat(baseInput, 32)
		if err != nil {
			fmt.Println("Invalid input for base: ", err)
			return
		}

		fmt.Println("Enter height of parallelogram: ")
		scanner.Scan()
		heightInput := scanner.Text()
		height, err := strconv.ParseFloat(heightInput, 32)
		if err != nil {
			fmt.Println("Invalid input for height: ", err)
			return
		}

		fmt.Println("Enter side length of parallelogram: ")
		scanner.Scan()
		sideInput := scanner.Text()
		side, err := strconv.ParseFloat(sideInput, 32)
		if err != nil {
			fmt.Println("Invalid input for side: ", err)
			return
		}

		parallelogram := model.Parallelogram{Base: float32(base), Height: float32(height), Side: float32(side)}
		shape = &parallelogram

	default:
		fmt.Println("Invalid input")
		return
	}

	printShapeResult(shape)
	fmt.Print("\nDo you want to try enter another shape? (y/n): ")
	scanner.Scan()
	choice := scanner.Text()
	if choice == "y" || choice == "Y" {
		main()
	}

	fmt.Println("Exiting the program......")
}
