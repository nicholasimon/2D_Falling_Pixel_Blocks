package main

import (
	"math/rand"

	"time"

	rl "github.com/lachee/raylib-goplus/raylib"
)

var (
	pixmap                                                          = make([]pix, mona)
	mona                                                            = monh * monw
	monw                                                            = 1280
	monh                                                            = 720
	fps                                                             = 60
	framecount                                                      int
	nextblockon, pauseon, numberson, debugon, gridon, centerlineson bool
	onoff2, onoff3, onoff6, onoff10, onoff15, onoff30               bool
	imgs                                                            rl.Texture2D
	camera, camera2x, camera4x, camera8x                            rl.Camera2D
)

type pix struct {
	color rl.Color
	activ bool
}

func raylib() { // MARK: raylib
	rl.InitWindow(monw, monh, "blokkzz")
	rl.SetExitKey(rl.KeyEnd)          // key to end the game and close window
	imgs = rl.LoadTexture("imgs.png") // load images
	rl.SetTargetFPS(fps)
	rl.HideCursor()
	//rl.ToggleFullscreen()
	for !rl.WindowShouldClose() {

		framecount++

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)
		drawlayers()
		rl.EndMode2D()

		drawnocamera()

		update()

		rl.EndDrawing()

	}
	rl.CloseWindow()
}
func drawlayers() { // MARK: drawlayers

}
func drawnocamera() { // MARK: drawnocamera

	for a := 0; a < len(pixmap); a++ {

		if pixmap[a].activ {
			x := a % monw
			y := a / monw
			rl.DrawPixel(x, y, pixmap[a].color)
		}

	}

}
func update() { // MARK: update

	timers()
	input()
	updatepix()

	if debugon {
		debug()
	}
}
func updatepix() { // MARK: updatepix

	for a := len(pixmap) - monw; a > 0; a-- {

		if pixmap[a].activ && !pixmap[a+monw].activ {
			pixmap[a+monw] = pixmap[a]
			pixmap[a].activ = false

		}

	}

}
func createshape() { // MARK: createshape

	choose := rInt(1, 5)

	position := rInt(128, monw-128)
	position += (monw * 40)

	switch choose {
	case 1:
		count := 0

		// square
		for a := 0; a < 1024; a++ {
			pixmap[position].color = randomgreen()
			pixmap[position].activ = true
			position++
			count++
			if count == 32 {
				count = 0
				position += monw
				position -= 32
			}
		}
	case 2:
		// rectangle
		count := 0

		length := 128
		length2 := 32
		area := length * length2
		for a := 0; a < area; a++ {
			pixmap[position].color = randombluelight()
			pixmap[position].activ = true
			position++
			count++
			if count == length {
				count = 0
				position += monw
				position -= length
			}
		}
	case 3:
		// triangle

		length := 64
		for {

			for a := 0; a < length; a++ {
				pixmap[position].color = randomyellow()
				pixmap[position].activ = true
				position++
			}
			position -= length
			position++
			position += monw
			length -= 2

			if length < 2 {
				break
			}
		}
	case 4:
		// diamond

		length := 64
		for {
			for a := 0; a < length; a++ {
				pixmap[position].color = randomyellow()
				pixmap[position].activ = true
				position++
			}
			position -= length
			position++
			position += monw
			length -= 2

			if length < 2 {
				break
			}
		}

		length = 64
		for {
			for a := 0; a < length; a++ {
				pixmap[position].color = randomorange()
				pixmap[position].activ = true
				position++
			}
			position -= length
			position++
			position -= monw
			length -= 2

			if length < 2 {
				break
			}
		}

	}

}
func main() { // MARK: main
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLogLevel(rl.LogError) // hides info window
	rl.InitWindow(monw, monh, "setscreen")
	setscreen()
	rl.CloseWindow()
	setinitialvalues()
	raylib()
}
func timers() { // MARK: timers

	if framecount%2 == 0 {
		if onoff2 {
			onoff2 = false
		} else {
			onoff2 = true
		}
	}
	if framecount%3 == 0 {
		if onoff3 {
			onoff3 = false
		} else {
			onoff3 = true
		}
	}
	if framecount%6 == 0 {
		if onoff6 {
			onoff6 = false
		} else {
			onoff6 = true
		}
	}
	if framecount%10 == 0 {
		if onoff10 {
			onoff10 = false
		} else {
			onoff10 = true
		}
	}
	if framecount%15 == 0 {
		if onoff15 {
			onoff15 = false
		} else {
			onoff15 = true
		}
	}
	if framecount%30 == 0 {
		if onoff30 {
			onoff30 = false
		} else {
			onoff30 = true
		}
	}

}
func input() { // MARK: input

	if rl.IsKeyPressed(rl.KeySpace) {

		createshape()
	}

	if rl.IsKeyPressed(rl.KeyKpMultiply) {
		if centerlineson {
			centerlineson = false
		} else {
			centerlineson = true
		}
	}

	if rl.IsKeyPressed(rl.KeyKpDecimal) {
		if debugon {
			debugon = false
		} else {
			debugon = true
		}
	}

	if rl.IsKeyPressed(rl.KeyKpAdd) {
		if camera.Zoom == 1.0 {
			camera.Zoom = 2.0
		} else if camera.Zoom == 2.0 {
			camera.Zoom = 3.0
		} else if camera.Zoom == 3.0 {
			camera.Zoom = 4.0
		}
	}
	if rl.IsKeyPressed(rl.KeyKpSubtract) {
		if camera.Zoom == 2.0 {
			camera.Zoom = 1.0
		} else if camera.Zoom == 3.0 {
			camera.Zoom = 2.0
		} else if camera.Zoom == 4.0 {
			camera.Zoom = 3.0
		}
	}
	if rl.IsKeyPressed(rl.KeyPause) {
		if pauseon {
			pauseon = false
		} else {
			pauseon = true
		}
	}

}
func debug() { // MARK: debug
	rl.DrawRectangle(monw-300, 0, 500, monw, rl.Fade(rl.Black, 0.8))
	rl.DrawFPS(monw-290, monh-100)

	//rl.DrawText(xtext, monw-290, 10, 10, rl.White)
	//rl.DrawText("blok9[0].x", monw-150, 10, 10, rl.White)

}
func setinitialvalues() { // MARK: setinitialvalues

}
func setscreen() { // MARK: setscreen

	rl.SetWindowSize(monw, monh)

	camera.Zoom = 1.0
	camera.Target.X = 0
	camera.Target.Y = 0

	camera2x.Zoom = 2.0
	camera4x.Zoom = 4.0
	camera8x.Zoom = 8.0

	camera4x.Target.X = 0
	camera4x.Target.Y = 0
}

// random colors https://www.rapidtables.com/web/color/RGB_Color.html
func randomgrey() rl.Color {
	color := rl.NewColor(uint8(rInt(105, 192)), uint8(rInt(105, 192)), uint8(rInt(105, 192)), 255)
	return color
}
func randombluelight() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 180)), uint8(rInt(120, 256)), uint8(rInt(120, 256)), 255)
	return color
}
func randombluedark() rl.Color {
	color := rl.NewColor(0, 0, uint8(rInt(120, 250)), 255)
	return color
}
func randomyellow() rl.Color {
	color := rl.NewColor(255, uint8(rInt(150, 256)), 0, 255)
	return color
}
func randomorange() rl.Color {
	color := rl.NewColor(uint8(rInt(250, 256)), uint8(rInt(60, 210)), 0, 255)
	return color
}
func randomred() rl.Color {
	color := rl.NewColor(uint8(rInt(128, 256)), uint8(rInt(0, 129)), uint8(rInt(0, 129)), 255)
	return color
}
func randomgreen() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 170)), uint8(rInt(100, 256)), uint8(rInt(0, 50)), 255)
	return color
}
func randomcolor() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 256)), uint8(rInt(0, 256)), uint8(rInt(0, 256)), 255)
	return color
}

// random numbers
func rF32(min, max float32) float32 {
	return (rand.Float32() * (max - min)) + min
}
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	a := int32(rand.Intn(max-min) + min)
	return a
}
func rFloat32(min, max int) float32 {
	a := float32(rand.Intn(max-min) + min)
	return a
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}
