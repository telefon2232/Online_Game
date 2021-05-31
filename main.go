
// +build example

package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"net"
	"time"
)

const (
	screenWidth  = 640
	screenHeight = 480
	size_person = 64
)

var flag bool = true

type Game struct {
	noiseImage *image.RGBA
}

type client struct {
	id uint
	x float64
	y float64
}

var x_player float64 = 0
var y_player float64 = screenHeight



func connect_to(data client) {
	conn, err := net.Dial("udp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	//fmt.Println(data.id)
	fmt.Fprint(conn,data)
	//fmt.Fprintf(conn, data.id)

	conn.Close()
}


var img *ebiten.Image
var bullet *ebiten.Image

func init() {
	fmt.Println("Init has been started...")
	var err error
	img, _, err = ebitenutil.NewImageFromFile("tile.png")
	bullet, _, err = ebitenutil.NewImageFromFile("bullet.png")
	if err != nil {
		fmt.Println("Image path not found")
		log.Fatal(err)
	}
}
func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x0f, 0xf0, 0xa0, 0xff})
	op := &ebiten.DrawImageOptions{}
	trase := &ebiten.DrawImageOptions{}
	trase.GeoM.Translate(screenHeight/2,screenWidth/2)

	if ebiten.IsKeyPressed(ebiten.KeyW){

		y_player--
	}
	if ebiten.IsKeyPressed(ebiten.KeyA){

		x_player--
	}
	if ebiten.IsKeyPressed(ebiten.KeyS){

		y_player++
	}
	if ebiten.IsKeyPressed(ebiten.KeyD){

		x_player++
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace){
		screen.DrawImage(bullet, trase)
		//d:=client{23,44,55}
		//connect_to(d)


	}
	if y_player<screenHeight/2 {
		y_player = screenHeight / 2
	}
	if y_player > screenHeight - size_person{
		y_player = screenHeight - size_person
	}


	op.GeoM.Translate(x_player,y_player)
	screen.DrawImage(img,op)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f\n%s", ebiten.CurrentFPS(),time.Now().Format(time.StampMilli)))

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