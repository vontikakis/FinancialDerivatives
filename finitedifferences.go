/**
 * Finite Difference pricing an american put option
 * single period and multiperiod
 * 
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
 * Amer
 */
func finiteDiffrenceExpicit(S, K, r, sigma, time float64, tNoSteps, sNoSteps int) float64{

    var sigmaSqrt, deltaS, deltaT, r1, r2 float64
    var M, N int

    sigmaSqrt = sigma*sigma

    M = sNoSteps + (sNoSteps%2)
    deltaS = 2.00*S/float64(M)

    N = tNoSteps
    deltaT = time/float64(N)

    sValues  := make([]float64,M+1)

    for m := 0; m <= M; m++ {
        
        sValues[m] = float64(m)* deltaS
    }

    a := make([]float64,M)
    b := make([]float64,M)
    c := make([]float64,M)

    r1 = 1.00/(1.00+r*deltaT)
    r2 = deltaT/(1.00+r*deltaT)

    for j:= 1; j< M; j++ {
        
        a[j] = r2*0.5* float64(j)*(-r+sigmaSqrt* float64(j))
        b[j] = r1*(1.00-sigmaSqrt*float64(j)*float64(j)*deltaT)
        c[j] = r2*0.5*float64(j)*(r+sigmaSqrt*float64(j))
    }

    fNext := make([]float64,M+1)

    for m := 0; m <= M; m++ {
        
        fNext[m] = math.Max(0.00,K-sValues[m])
    }
   
    f := make([]float64,M+1)

    for t := N-1; t >= 0; t-- {
        
        f[0]=K

        for m:= 1; m < M; m++ {
            
            f[m] = a[m]*fNext[m-1]+b[m]*fNext[m]+c[m]*fNext[m+1]
            f[m] = math.Max(f[m],K-sValues[m])

        }

        f[M] = 0

        for m := 0; m <= M; m++ {
        
            fNext[m] = f[m]
        }
    }

    return f[M/2]

}


func main() {

    var resultCallOption float64
    var S, K, r, sigma, time float64 = 50.00, 50.00, 0.10, 0.4, 0.4167
    var sNoSteps, tNoSteps = 20, 11

    resultCallOption = finiteDiffrenceExpicit(S, K, r, sigma, time, sNoSteps, tNoSteps)
    
    fmt.Println("American Put Options Single Period Binomial Option Price is ",resultCallOption)



}
