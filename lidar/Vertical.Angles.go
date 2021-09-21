package lidar

type VerticalAngle struct {
	angleDeg             int16
	verticalCorrectionMM float32
}

var VerticalAngleMap = [16]VerticalAngle{
	{-15, 11.2},
	{1, -0.7},
	{-13, 9.7},
	{3, -2.2},
	{-11, 8.1},
	{5, -3.7},
	{-9, 6.6},
	{7, -5.1},
	{-7, 5.1},
	{9, -6.6},
	{-5, 3.7},
	{11, -8.1},
	{-3, 2.2},
	{13, -9.7},
	{-1, 0.7},
	{15, -11.2},
}
