/**
 * Bionomial option pricing in european call option
 * single period and multiperiod
 * 
 * https://en.wikipedia.org/wiki/Call_option
 * 
 * @link https://www.vontikakis.com/
 * @author Vasilis Vontikakis
 * 
 */

package main

import(
    "fmt"
    "math"
)


/** 
 * bionomial single period
 */
func singlePeriodBinomial(S, K, r, u, d float64) float64{

    var callUp, callDown, probUp, probDown float64

    probUp = (math.Exp(r)-d)/(u-d)

    probDown = 1.00 - probUp

    callUp = math.Max(0.00,(u*S-K))

    callDown = math.Max(0.00,(d*S-K))

    return (math.Exp(-r)*(probUp*callUp+probDown*callDown))

}


/** 
 * Bionomial multiperiod pricing models
 */ 
func multiPeriodBinomial(S, K, r, u, d float64, noPeriod int) float64{

    // var Rinv, uu, probUp, probDown float64
    var Rinv, uu, probUp, probDown float64

    prices := make([]float64,noPeriod*2)
    callValues  := make([]float64,noPeriod*2)
    
    Rinv = math.Exp(-r)

    uu = u*u

    probUp = (math.Exp(r)-d)/(u-d)

    probDown = 1.00 - probUp

    prices[0] = S*math.Pow(d,float64(noPeriod))

    for i := 1; i <= noPeriod; i++ {
        prices[i] =  uu*prices[i-1]
    }

    for i := 0; i <= noPeriod; i++ {
        callValues[i] =  math.Max(0.00,(prices[i]-K))
    }

    for step := (noPeriod-1); step>=0; step-- {

        for i := 0; i <= step; i++ {

            callValues[i] = (probUp*callValues[i+1]+probDown*callValues[i])*Rinv
        }

    }

    return callValues[0]

}

func main() {

    var resultCallOption float64
    var S, K, r, u, d float64 = 100.00, 100.00, 0.025, 1.05, 1/1.05

    resultCallOption = singlePeriodBinomial(S, K, r, u, d)
    
    fmt.Println("European Call Single Period Binomial Option Price is ",resultCallOption)

    var noPeriod int = 2

    resultCallOption = multiPeriodBinomial(S, K, r, u, d, noPeriod)
    
    fmt.Println("European Call Multi Period Binomial Option Price is ",resultCallOption)

}
