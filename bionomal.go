/**
 * Bionomial option pricing in european call option
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
func singlePeriodBinomial(S, K, r, u, d float64, noPeriod int) float64{

    probUp := (math.Exp(r)-d)/(u-d)

    probDown := 1.00 - probUp

    callUp :=   math.Max(0.00,(u*S-K))

    callDown :=  math.Max(0.00,(d*S-K))

    return math.Exp(-r)*(probUp*callUp+probDown*callDown)

}


func main() {

    var S, K, r, u, d float64 = 100.00, 100.00, 0.025, 1.05,(1/1.05)
    var noPeriod int = 1

    resultCallOption := singlePeriodBinomial(S, K, r, u, d, noPeriod)
    
    fmt.Println("European Call Option Price is ",resultCallOption)

}
