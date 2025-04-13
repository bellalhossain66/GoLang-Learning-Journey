package main;
import "fmt";

func main() {
	var a [5]int;
	fmt.Println("array empty: ",a);

	a[4] = 100;
	fmt.Println("Set: ", a);
	fmt.Println("Get: ",a);

	fmt.Println("Length: ", len(a));

	b := [5]int{1,2,3,4,5};
	fmt.Println("Decimal: ", b);

	b = [...]int{1,2,3,4,5};
	fmt.Println("Decimal: ", b);

	b = [...]int{100, 3:200, 300};
	fmt.Println("Index: ", b);

	var twoDimentional [2][3]int;
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDimentional[i][j] = i + j;
		}
	}
	fmt.Println("2d: ", twoDimentional);

	twoDimentional = [2][3]int {
		{1,2,3},
		{1,2,3},
	}
	fmt.Println("2d: ", twoDimentional);
}