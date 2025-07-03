## scanning a values
Note: that you need to pass teh pointer to the variable and not the variable itself
```
	const inflationRate float64 = 6.5
	var investment float64 = 1000
	years := 10
	returns := 5.5
	fmt.Println("Investment Calculator Enter your amount here: ")
	// scanning the function for the input
	fmt.Scan(&investment)
	// This is a pointer to the i
```

use <variable Name>  := fmt.Sprintf() to store the formatted variable into the file

## Functions in go
- Code on demand blocks

```
func returnvalue(a int) (square int, cube int) {
	square = a * a
	cube = a * a * a
	return
}

```
Return various values

use return to brekout
