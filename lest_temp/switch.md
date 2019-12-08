###  1.My proposals and needs:

-  I hope that the switch statement in the new version of go language will be as easy to use as the switch language in Java 13 version.Thank you for hoping that the official will accept my opinion!

- This is how we currently write switch statements:

```go
package main

import "fmt"

var (
	marks = 90
	grade = "D"
)

func main() {
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Println("Your grade is:", grade)
}
```

- I suggest designing a simplified or more advanced switch statement:

```go
       <T> grade = switch (marks) {
    	        case 90 -> "A"
		case 80 -> "B"
		..... 
    	        default -> "D"
	};
       fmt.Println("Your grade is:", grade)
```

- 👆 `<T>` Is generic & I think it can also support function return💡

```go
       <T> grade = switch (marks,arg[]) {
    	        case 90 -> func (int arg[]) int grade{
			return arg + 90;
		}
		case 80 -> func (int arg[]) int grade{
			return arg + 80;
		}
		..... 
    	        default -> ....  
	};
	fmt.Println("Your grade is:", grade)
```

### 2.The version of go I currently use is:
 
`go version go1.13 darwin/amd64`

### 3.What did you expect to see?

- I hope this new version has this new feature ~~

### 4.Of course, these are my personal needs and opinions, and do not represent everyone's thoughts. I just have this idea. I hope you can consider it, thank you😄🤝～

- 2019-12-07 18:58:07 UTC + (＋0800) By:BinScholl Email:BinScholl@foxmail.com