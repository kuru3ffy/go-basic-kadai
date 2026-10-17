package main

import (
	"errors"
	"fmt"
	"math"
)

type ValidationError struct {
	Message string
	Field   string
	Value   float64
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("[バリデーションエラー] %s：%s（Value：%.2f）", e.Field, e.Message, e.Value)
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func validateShape(s Shape) error {
	switch shape := s.(type) {
	case Rectangle:
		if shape.Width <= 0 {
			return &ValidationError{Message: "幅は正の数を指定してください。", Field: "Rectangle.Width", Value: shape.Width}
		}
		if shape.Height <= 0 {
			return &ValidationError{Message: "高さは正の数を指定してください。", Field: "Rectangle.Height", Value: shape.Height}
		}
	case Circle:
		if shape.Radius <= 0 {
			return &ValidationError{Message: "半径は正の数を指定してください。", Field: "Circle.Radius", Value: shape.Radius}
		}
	default:
		return fmt.Errorf("対応していない図形です：%T", s)
	}
	return nil
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 10.0, Height: 5.0},
		Circle{Radius: 3.0},
		Rectangle{Width: -2.0, Height: 5.0},
		Circle{Radius: -1.0},
		Rectangle{Width: 1.0, Height: 0.0},
	}
	var calcShapes []Shape

	fmt.Println("--- 図形のバリデーション ---")
	for i, shape := range shapes {
		if err := validateShape(shape); err != nil {
			var validationErr *ValidationError
			if errors.As(err, &validationErr) {
				fmt.Printf("#%d：%T => NG：%v\n", i+1, shape, validationErr)
				fmt.Printf(" [詳細] Field=%s, Value=%.2f\n", validationErr.Field, validationErr.Value)
			} else {
				fmt.Printf("#%d：%T => 通常エラー：%v\n", i+1, shape, err)
			}
			continue
		}
		fmt.Printf("#%d：%T => バリデーションOK\n", i+1, shape)
		calcShapes = append(calcShapes, shape)
	}

	fmt.Println("\n--- 検証済み図形の一括計算 ---")
	for _, shape := range calcShapes {
		fmt.Printf("図形：%T｜面積：%.2f ｜周長：%.2f\n", shape, shape.Area(), shape.Perimeter())
	}
}
