/*
	For reference, some of the code below was pulled from the following packages:
	- https://github.com/lucasb-eyer/go-colorful
	- https://github.com/hisamafahri/coco
*/

package nomads

import (
	"image/color"
	"math"
)

func GetColorDifference(color1 color.RGBA, color2 color.RGBA) float64 {
	l1, a1, b1 := convertRGBToLAB(color1)
	l2, a2, b2 := convertRGBToLAB(color2)

	return math.Sqrt(Squared(l1-l2) + Squared(a1-a2) + Squared(b1-b2))
}

func RandomRGBA(colors ...string) color.RGBA {
	randomRed := uint8(RandomWithin(0, 255))
	randomGreen := uint8(RandomWithin(0, 255))
	randomBlue := uint8(RandomWithin(0, 255))

	if colors != nil {
		switch colors[0] {
		case "red":
			return color.RGBA{randomRed, 0, 0, 255}
		case "green":
			return color.RGBA{0, randomGreen, 0, 255}
		case "blue":
			return color.RGBA{0, 0, randomBlue, 255}
		case "gray":
			return color.RGBA{randomRed, randomRed, randomRed, 255}
		}
	}

	return color.RGBA{randomRed, randomGreen, randomBlue, 255}
}

func convertRGBToLAB(rgbColor color.RGBA) (float64, float64, float64) {
	x, y, z := convertRGBToXYZ(rgbColor)

	x /= 95.047
	y /= 100
	z /= 108.883

	if x > 0.008856 {
		x = math.Pow(x, (1.0 / 3.0))
	} else {
		x = (7.787 * x) + (16.0 / 116.0)
	}

	if y > 0.008856 {
		y = math.Pow(y, (1.0 / 3.0))
	} else {
		y = (7.787 * y) + (16.0 / 116.0)
	}

	if z > 0.008856 {
		z = math.Pow(z, (1.0 / 3.0))
	} else {
		z = (7.787 * z) + (16.0 / 116.0)
	}

	l := (116 * y) - 16
	a := 500 * (x - y)
	b := 200 * (y - z)

	return math.Round(l), math.Round(a), math.Round(b)
}

func convertRGBToXYZ(rgbColor color.RGBA) (float64, float64, float64) {
	red := float64(rgbColor.R)
	green := float64(rgbColor.G)
	blue := float64(rgbColor.B)

	red = red / 255
	green = green / 255
	blue = blue / 255

	// Assume sRGB
	if red > 0.04045 {
		red = math.Pow(((red + 0.055) / 1.055), 2.4)
	} else {
		red = (red / 12.92)
	}

	if green > 0.04045 {
		green = math.Pow(((green + 0.055) / 1.055), 2.4)
	} else {
		green = (green / 12.92)
	}

	if blue > 0.04045 {
		blue = math.Pow(((blue + 0.055) / 1.055), 2.4)
	} else {
		blue = (blue / 12.92)
	}

	x := (red * 0.4124564) + (green * 0.3575761) + (blue * 0.1804375)
	y := (red * 0.2126729) + (green * 0.7151522) + (blue * 0.072175)
	z := (red * 0.0193339) + (green * 0.119192) + (blue * 0.9503041)

	return math.Round(x * 100), math.Round(y * 100), math.Round(z * 100)
}

func ConvertToRGBA(colorToConvert color.Color) color.RGBA {
	return color.RGBAModel.Convert(colorToConvert).(color.RGBA)
}
