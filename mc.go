/**
 * Monte Carlo Simulation to value 
 * European Call Option Financial Derivatives
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
 *  generates random normal distributed variable N(0,1)
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
func europeanCallOption(S, K, r, sigma, time float64, num_sims int) float64{

    var S_T float64

    R := (r-0.05*math.Pow(sigma,2))*time

    SD := sigma * math.Sqrt(time)

    sum_payoffs := 0.00
    
    for i := 0; i < num_sims; i++ {

        S_T = S * math.Exp(R+SD*randomNormal())

        sum_payoffs += math.Max(0.00, S_T-K)
    }

    average_payoff := sum_payoffs/float64(num_sims)

    discount_in_time_payoff := (math.Exp(-r*time))*average_payoff

    return discount_in_time_payoff
}


func main() {

    rand.Seed(time.Now().Unix())
    
    resultUniform := randomUniform()
    
    fmt.Println(resultUniform)

    resultNormal := randomNormal()
    
    fmt.Println(resultNormal)

    resultCallOption := europeanCallOption(9, 10, 0.04, 0.3, 0.25, 10000)
    
    fmt.Println(resultCallOption)

}
