/**
 * Monte Carlo Simulation to value 
 * European Call Option Financial Derivatives
 * 
 * https://en.wikipedia.org/wiki/Call_option
 * 
 * @link https://www.vontikakis.com/
 * @copyright Copyright &copy; 2015-2017 Vasilis Vontikakis
 * @author Vasilis Vontikakis
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
func randomUniform(scale ...float64) float64{
    
    var scalarVariable float64 = 1.0000

    if len(scale) > 0 {

        scalarVariable = float64(scale[0])
    }
    
    randMax := float64(math.MaxInt32)

    return (float64(rand.Intn(math.MaxInt32))/randMax)*scalarVariable

}

/**
 *  generates random normal distributed variables N(0,1)
 *  using Marsaglia polar method
 */
func randomNormal() float64{

    var v1 float64
    var v2 float64
    var v  float64

    v = 2.00

    for v>1 {
        
        v1 = 2*randomUniform()-1
        v2 = 2*randomUniform()-1
            
        v =  v1*v1 + v2*v2
    }

    return (v1)*((math.Pow(-2*math.Log(v),0.5)))

}

/** 
 * Monte carlo simulation for European Call Option
 */
func europeanCallOption(S, K, r, sigma, time float64, numSims int) float64{

    var ST float64

    R := (r-0.05*math.Pow(sigma,2))*time

    SD := sigma * math.Sqrt(time)

    sumPayoffs := 0.00
    
    for i := 0; i < numSims; i++ {

        ST = S * math.Exp(R+SD*randomNormal())

        sumPayoffs += math.Max(0.00, ST-K)
    }

    averagePayoff := sumPayoffs/float64(numSims)

    discountInTimePayoff := (math.Exp(-r*time))*averagePayoff

    return discountInTimePayoff
}


func main() {

    rand.Seed(time.Now().Unix())

    var stockPrice, strikePrice, riskFreeRate, sigma, durationTime float64 = 9.00, 10.00, 0.04, 0.3, 0.25
    var numSimulations int = 10000

    resultCallOption := europeanCallOption(stockPrice, strikePrice, riskFreeRate, sigma, durationTime, numSimulations)
    
    fmt.Println("European Call Option Price is ",resultCallOption)

}
