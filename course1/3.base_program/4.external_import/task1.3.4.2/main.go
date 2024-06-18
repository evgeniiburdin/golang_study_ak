// Примечание: библиотека github.com/ksrof/gocolors больше не доступна
// Последняя доступная версия 0.1.0, в ней нет gocolors.NewRGBColor
// Документация есть только на pkg.go.dev, но там только перечислены функции буквально, без типов, без констант и т.п.

package main

import (
	"fmt"
	"github.com/ksrof/gocolors"
)

func main() {
	fmt.Println(ColorizeRed("red"))
	fmt.Println(ColorizeOrange("orange"))
	fmt.Println(ColorizeYellow("yellow"))
	fmt.Println(ColorizeGreen("green"))
	fmt.Println(ColorizeCyan("cyan"))
	fmt.Println(ColorizeBlue("blue"))
	fmt.Println(ColorizePurple("purple"))
	fmt.Println(ColorizeMagenta("magenta"))
	fmt.Println(ColorizeWhite("white"))
	fmt.Println(ColorizeCustom("red", 172, 229, 238))
}

func ColorizeRed(a string) string {
	return gocolors.Red(a, "")
}

func ColorizeOrange(a string) string {
	color := fmt.Sprintf("\x1b[38;2;255;146;24m")
	return gocolors.Color(a, color, "")
}

func ColorizeYellow(a string) string {
	return gocolors.Yellow(a, "")
}

func ColorizeGreen(a string) string {
	return gocolors.Green(a, "")
}

func ColorizeCyan(a string) string {
	return gocolors.Cyan(a, "")
}

func ColorizeBlue(a string) string {
	return gocolors.Blue(a, "")
}

func ColorizePurple(a string) string {
	color := fmt.Sprintf("\x1b[38;2;128;0;255m")
	return gocolors.Color(a, color, "")
}

func ColorizeMagenta(a string) string {
	return gocolors.Magenta(a, "")
}

func ColorizeWhite(a string) string {
	return gocolors.White(a, "")
}

func ColorizeCustom(a string, r, g, b uint8) string {
	color := fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
	return gocolors.Color(a, color, "")
}
