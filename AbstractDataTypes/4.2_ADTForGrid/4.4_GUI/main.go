package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	grid "guigameoflife.go"
	"image/color"
	"math/rand"
	"time"
)

var (
	rows int
	cols int

	rect *canvas.Rectangle

	//Holds rectangle objects
	segments []fyne.CanvasObject
)

func output() *fyne.Container {
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid.GetGrid()[col][row] == false {
				rect = canvas.NewRectangle(&color.RGBA{B: 200, R: 200, G: 200, A: 255})
			} else {
				rect = canvas.NewRectangle(&color.RGBA{B: 0, R: 255, G: 0, A: 255})
			}

			rect.Resize(fyne.NewSize(10, 10))
			rect.Move(fyne.NewPos(float32(row*11), float32(col*11)))
			segments = append(segments, rect)
		}
	}
	return container.NewWithoutLayout(segments...)
}

func main() {
	getGrid := grid.GetGrid()
	getGrid.InitializeGrid(25, 25)

	getNewGrid := grid.GetNewGrid()
	getNewGrid.InitializeGrid(25, 25)

	for numberCritters := 0; numberCritters < 4; numberCritters++ {
		r := 5 + rand.Intn(10)
		c := 5 + rand.Intn(10)

		getGrid.BringAlive(r, c)
		getGrid.BringAlive(r+1, c)
		getGrid.BringAlive(r+1, c+1)
		getGrid.BringAlive(r-1, c)
		getGrid.BringAlive(r-2, c-1)
	}

	a := app.New()
	w := a.NewWindow("GAME OF LIFE - Hit Any Key To Quit")

	w.Resize(fyne.NewSize(300, 300))
	w.SetFixedSize(true)

	go func() {
		for {
			newContainer := output()
			w.SetContent(newContainer)

			time.Sleep(1 * time.Second)
			getGrid.EvolveGrid()
		}
	}()

	w.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		//Shuts down simulation
		w.Close()
	})

	w.ShowAndRun()
}
