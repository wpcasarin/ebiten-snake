package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"log"
)

const (
	cellSize   int    = 20
	gridLength int    = 30
	width      int    = cellSize * gridLength
	height     int    = cellSize * gridLength
	blockSize  int    = cellSize - 1
	title      string = "Snake"
)

var (
	gophersImage *ebiten.Image
	//block        = ebiten.NewImage(cellSize, cellSize)
	bgColor    = color.Black
	snakeColor = color.RGBA{0x60, 0xa6, 0x65, 0xff}
	fruitColor = color.RGBA{0xcd, 0xb3, 0x5d, 0xff}
	snake      = Snake{body: []Vector2{{3, 10}, {4, 10}, {5, 10}},
		direction: Vector2{1, 0},
		block:     ebiten.NewImage(blockSize-1, blockSize-1)}
)

type Game struct {
}

type Vector2 struct {
	x int
	y int
}

type Snake struct {
	body      []Vector2
	direction Vector2
	block     *ebiten.Image
}

func (snake *Snake) Move() {
	newBody := snake.body[1:]
	newHead := snake.body[len(snake.body)-1]
	newHead.x += snake.direction.x
	newHead.y += snake.direction.y
	if newHead.x > gridLength {
		newHead.x = 1
	}
	if newHead.x <= 0 {
		newHead.x = gridLength
	}
	if newHead.y > gridLength {
		newHead.y = 1
	}
	if newHead.y <= 0 {
		newHead.y = gridLength
	}
	snake.body = append(newBody, newHead)

}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) && snake.direction.y != 1 {
		snake.direction.x = 0
		snake.direction.y = -1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) && snake.direction.y != -1 {
		snake.direction.x = 0
		snake.direction.y = 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && snake.direction.x != 1 {
		snake.direction.x = -1
		snake.direction.y = 0
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) && snake.direction.x != -1 {
		snake.direction.x = 1
		snake.direction.y = 0
	}
	snake.Move()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//screen.Fill(bgColor)
	for _, v := range snake.body {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(v.x*cellSize-cellSize), float64(v.y*cellSize-cellSize))
		screen.DrawImage(snake.block, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	game := &Game{}
	snake.block.Fill(snakeColor)
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)
	ebiten.SetTPS(10)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
