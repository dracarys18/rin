package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	rp "rin/libs/rope"
	"strings"
)

const width = 1600
const height = 900

var isFullScreen = false

func free(font rl.Font) {
	rl.UnloadFont(font)
	rl.CloseWindow()
}

func toggleFullScreen() {
	if isFullScreen {
		rl.SetWindowSize(width, height)
		isFullScreen = false
	} else {
		rl.SetWindowSize(rl.GetMonitorWidth(rl.GetCurrentMonitor()), rl.GetMonitorHeight(rl.GetCurrentMonitor()))
		isFullScreen = true
	}
}
func handleKeyPressed(key int32, rope *rp.Rope) {
	switch key {
	case rl.KeyEnter:
		toggleFullScreen()
	default:
		if (key >= 32) && (key <= 125) {
			str := string(rune(key))
			if !(rl.IsKeyDown(rl.KeyCapsLock) || rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift)) {
				str = strings.ToLower(str)
			}
			newRope := rp.NewRope(str)
			*rope = rope.Concatenate(newRope)
		}
	}
}
func main() {
	rope := rp.Rope{}

	rl.InitWindow(width, height, "Rin")

	font := rl.LoadFontEx("/Users/karthikey.hegde/Library/Fonts/HasklugNerdFontMono-Bold.otf", 20, nil)

	rl.SetWindowPosition(0, 0)
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	defer free(font)
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		handleKeyPressed(rl.GetKeyPressed(), &rope)
		rl.DrawTextEx(font, rope.String(), rl.Vector2{X: 10, Y: 10}, 20, 0, rl.RayWhite)
		rl.EndDrawing()
	}
}
