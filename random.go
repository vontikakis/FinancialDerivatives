/**
 * Random variable generator
 *
 * @link https://www.vontikakis.com/
 * @copyright Copyright &copy; 2015-2017 Vasilis Vontikakis
 * @author Vasilis Vontikakis
 *
 * randoma variables generator
 * 
 */

package main

import(
    "fmt"
    "math"
    "math/rand"
    "time"
)

/**
 *  get generate random uniform variable in space [0,scale]
 */
func randomUniform(scale ...float32) float32{
    
    var scalarVariable float32 = 1.0000

    if len(scale) > 0 {

        scalarVariable = float32(scale[0])
    }
    
    randMax := float32(math.MaxInt32)

    return (float32(rand.Intn(math.MaxInt32))/randMax)*scalarVariable

}

func main() {

    rand.Seed(time.Now().Unix())
    
    result := randomUniform()
    
    fmt.Println(result)
}
