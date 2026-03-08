// package a

// import (
// 	"fmt"
// 	"image"
// 	"log"
// 	"math/rand"
// 	"os"
// 	"time"

// 	"github.com/fogleman/gg"
// )


// // func squareMaker(imageBase, imageKindaRefined, x, y, ) {
	
// // }

// func a() {
//     // Create a new source with current time as seed
//     source := rand.NewSource(time.Now().UnixNano())
//     r := rand.New(source)

//     imageRaw, err := os.Open("waffle.png")
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer imageRaw.Close()
//     imageKindaRefined, _, err := image.Decode(imageRaw)
//     if err != nil {
//         log.Fatal(err)
//     }

//     bounds := imageKindaRefined.Bounds()
//     rawWidth := bounds.Max.X
//     rawHeight := bounds.Max.Y
//     imageBase := gg.NewContext(rawWidth, rawHeight)
//     // Fill background with white (or any color you prefer)
//     imageBase.SetRGB(1, 1, 1)
//     imageBase.Clear()
//     // Adjust step size based on your rectangle size
//     // 
//     // 

//     density := 8

//     randStep := density 
//     for i := -200; i < rawWidth; i += randStep {
//     randStep = density + r.Intn(density)
//         for a := 0; a < rawHeight; a += randStep{
//             imageBase.SetRGBA(10, 10, 10, 1)
//             imageBase.DrawRectangle(float64(i), float64(a), float64(i), float64(a))
//             imageBase.Fill()
//         }

//         for j := -300; j < rawHeight; j += randStep {
//             randomNum := 0 + r.Float64() * float64(randStep*density+20) 
//             transparancy := (randomNum+float64(density))/randomNum
//             r, g, b, _ := imageKindaRefined.At(i + randStep, j).RGBA()

//             if g >= 200 && (r < 15 && b < 15){
//                   g = 0
//             }

//             // imageBase.SetRGBA(0, 0, 0, 0)
//             // imageBase.DrawRectangle(float64(i), float64(j), float64(i), float64(j))
//             // imageBase.Fill()
//             imageBase.SetRGBA(float64(r)/65535, float64(g)/65535, float64(b)/65535, transparancy)
//             imageBase.DrawRectangle(float64(i), float64(j), randomNum, randomNum)
//             imageBase.Fill()
//         }
//         fmt.Print(fmt.Sprintf("Done with Width: %d", i))
//         fmt.Print("\n")
//     }


//     err = imageBase.SavePNG("waffleexport.png")
//     if err != nil {
	//         log.Fatal(err)
	//     }
	// }