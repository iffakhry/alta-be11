package main

import (
	"fmt"
	"math"
)

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (z Rect) HitungLuas() float64 {
	return z.width * z.height
}

func (c Circle) HitungLuas() float64 {
	return math.Pi * c.radius * c.radius
}

func HitungKelilingPersegi(panjang, lebar float64) float64 {
	return 2 * (panjang + lebar)
}

func HitungKelilingLingkaran(jari float64) float64 {
	return 0
}

func main() {
	rect := Rect{5.0, 4.0}
	cir := Circle{5.0}
	fmt.Printf("Area of rectangle rect = %0.2f\n", rect.HitungLuas())
	fmt.Printf("Area of circle cir = %0.2f\n", cir.HitungLuas())
	fmt.Println("Keliling:", HitungKelilingPersegi(rect.height, rect.width))
}
