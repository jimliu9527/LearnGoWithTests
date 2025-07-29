package structinterface

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{1.1, 2.1}
	got := Perimeter(rectangle)
	want := 6.4

	if got != want {
		t.Errorf("got is '%.2f', but want '%.2f'", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got is '%.2f' but want is '%.2f'", shape, got, want)
		}
	}

	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 1.0, Length: 2.0}, want: 2},
		{"Circle", Circle{Radius: 4}, 50.24},
		{name: "Triangle", shape: Triangle{Base: 1.0, Height: 2.0}, want: 1},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			checkArea(t, tt.shape, tt.want)
		})
	}
}
