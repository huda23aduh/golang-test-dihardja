package main
import "fmt"

func main() {
    var res = checkPalindrome("madama");
    fmt.Println("result = ", res)
    
}

func checkPalindrome(str string) bool {
    for i := 0; i < len(str); i++ {
        j := len(str)-1-i
        if str[i] != str[j] {
            return false   
        }
    }
    return true
}
