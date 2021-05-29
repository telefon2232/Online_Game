
// +build example

package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	_ "image/png"
	"log"
	"time"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	noiseImage *image.RGBA
}

type player_coordinate struct {
	x int
	y int
}
var img *ebiten.Image
func init() {
	fmt.Println("Init has been started...")
	var err error
	img, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		fmt.Println("Image path not found")
		log.Fatal(err)
	}
}
func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(screenWidth/3, screenHeight/2)
	screen.DrawImage(img,op)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f\n%s", ebiten.CurrentFPS(),time.Now().Format(time.RFC1123)))

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Golang game v0.0.1")

	g := &Game{	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}