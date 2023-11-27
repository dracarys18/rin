package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	rp "rin/libs/rope"
)

const width = 1600
const height = 900
const textSize = 20

type RinState struct {
	X            int32
	Y            int32
	isFullScreen bool
}

func free(font rl.Font) {
	rl.UnloadFont(font)
	rl.CloseWindow()
}

func (r *RinState) toggleFullScreen() {
	if r.isFullScreen {
		rl.SetWindowSize(width, height)
		r.isFullScreen = false
	} else {
		rl.SetWindowPosition(0, 0)
		rl.SetWindowSize(rl.GetMonitorWidth(rl.GetCurrentMonitor()), rl.GetMonitorHeight(rl.GetCurrentMonitor()))
		r.isFullScreen = true
	}
}
func (r *RinState) handleKeyPressed(key int32, rope *rp.Rope) {
	switch key {
	case rl.KeyEnter:
		r.toggleFullScreen()
	default:
		if (key >= 32) && (key <= 125) {
			str := string(rune(key))
			newRope := rp.NewRope(str)
			cursorWidth := rl.MeasureText(str, textSize)
			fmt.Println(cursorWidth)
			r.X += cursorWidth
			*rope = rope.Concatenate(newRope)
		}
	}
}

func main() {
	rope := rp.Rope{}

	r := RinState{X: 10, Y: 10}
	rl.InitWindow(width, height, "Rin")

	font := rl.LoadFontEx("/Users/karthikey.hegde/Library/Fonts/HasklugNerdFontMono-Bold.otf", 20, nil)

	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	defer free(font)
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		r.handleKeyPressed(rl.GetCharPressed(), &rope)
		text := rope.String()
		if int(rl.GetTime()*2)%2 == 0 {
			rl.DrawText("_", r.X, r.Y, 20, rl.RayWhite)
		}

		rl.DrawTextEx(font, text, rl.Vector2{X: 10, Y: 10}, 20, 0, rl.RayWhite)
		rl.EndDrawing()
	}
}
