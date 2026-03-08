package main

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/fogleman/gg"
)


// func squareMaker(imageBase, imageKindaRefined, x, y, ) {
	
// }

func main() {
    // Create a new source with current time as seed
    source := rand.NewSource(time.Now().UnixNano())
    r := rand.New(source)

    imageRaw, err := os.Open("mountain.png")
    if err != nil {
        log.Fatal(err)
    }
    defer imageRaw.Close()
    imageKindaRefined, _, err := image.Decode(imageRaw)
    if err != nil {
        log.Fatal(err)
    }

    bounds := imageKindaRefined.Bounds()
    rawWidth := bounds.Max.X
    rawHeight := bounds.Max.Y
    imageBase := gg.NewContext(rawWidth, rawHeight)
    // Fill background with white (or any color you prefer)
    imageBase.SetRGB(1, 1, 1)
    imageBase.Clear()
    // Adjust step size based on your rectangle size
    // 
    // 

    density := 20


    randStep := density 
    for aa := 0; aa < 2; aa++ {
        for i := -200; i < rawWidth; i += randStep {
            randStep = density + r.Intn(density)
                for j := -300; j < rawHeight; j += randStep {
                    randomNum := 0 + r.Float64() * float64(randStep*density+20) 
                    transparancy := (randomNum+float64(density))/randomNum
                    // Change the multiplication of the random float for cool
                    // r, g, b, _ := imageKindaRefined.At(i + randStep + int(r.Float64()*1), j).RGBA()
                    r, g, b, _ := imageKindaRefined.At(i + randStep, j + randStep).RGBA()

                    if g >= 200 && (r < 15 && b < 15){
                        g = 0
                    }

                    if r < 5 {
                        r = r+20
                    }

                    if b < 5{
                        b = b+20
                    }

                    if g < 5{
                        g = g+20
                    } 

                    imageBase.SetRGBA(float64(r)/65535, float64(g)/65535, float64(b)/65535, transparancy)
                    imageBase.DrawRectangle(float64(i), float64(j), randomNum, randomNum)
                    imageBase.Fill()
                }
                fmt.Print(fmt.Sprintf("Done with Width: %d", i))
                fmt.Print("\n")
        }
    }



    err = imageBase.SavePNG("output.png")
    if err != nil {
        log.Fatal(err)
    }
}


