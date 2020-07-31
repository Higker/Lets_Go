package main

import "fmt"

var(
	dp1,dp2 = 1,2
)

func main(){

	fmt.Println(climbStairs(3))

}

func climbStairs(n int) int  {
	switch n {
	case 1 :
		return 1
	case 2:
		return 2
	default:
		for i:=2;i<n;i++{
			dp1,dp2 = dp1,dp1+dp2
		}
		return dp2
	}
}