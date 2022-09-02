

package main

    import (
        "fmt"
    )

    func main(){
        var res = minOpr(8)
        fmt.Print("result = ")
        fmt.Print(res)
    }
    
    func minOpr(n int) (int) {
        totalIdx := n+1
        dp := make([]int, totalIdx)

        
        dp[1] = 0
        
        for i := 1; i <= n; i++ {
            
            dp[i] = 999999
            
            if i % 2 == 0 {
                var x = dp[i / 2]
                if (x + 1 < dp[i]) {
                    dp[i] = x + 1;
                }
            }
            
            
            var x = dp[i-1]
            if x+1 < dp[i] {
                dp[i] = x + 1;
            }
        }
        return dp[n]
    }

