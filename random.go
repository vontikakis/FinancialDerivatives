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
 *  generates random uniform variable in space [0,scale]
 */
func randomUniform(scale ...float32) float32{
    
    var scalarVariable float32 = 1.0000

    if len(scale) > 0 {

        scalarVariable = float32(scale[0])
    }
    
    randMax := float32(math.MaxInt32)

    return (float32(rand.Intn(math.MaxInt32))/randMax)*scalarVariable

}

/**
 *  generates random normal distributed variable in space [-1,1]
 *  using Marsaglia polar method
 */
func randomNormal() float32{

    var v1 float32
    var v2 float32
    var v  float32

    v = 2.00

    for v>1 {
        
        v1 = 2*randomUniform()-1
        v2 = 2*randomUniform()-1
            
        v =  v1*v1 + v2*v2
    }

    return (v1)*(float32(math.Pow(-2*math.Log(float64(v)),0.5)))

}

func main() {

    rand.Seed(time.Now().Unix())
    
    resultUniform := randomUniform()
    
    fmt.Println(resultUniform)

    resultNormal := randomNormal()
    
    fmt.Println(resultNormal)
}
