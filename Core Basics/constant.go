package main;
import (
    "fmt"
    "math"
)


const s string = "okay";
func main() {
	fmt.Println(s);

	const n = 50000000;
	const d = 3e20;
	fmt.Println(n,d,d/n);
	fmt.Println(int64(d/n));

	fmt.Println(math.Sin(n));
}