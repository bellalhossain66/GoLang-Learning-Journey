package main;
import "fmt";

func mutliReturn() (int, int) {
	return 5, 10;
}

func main() {
	a, b := mutliReturn();
	fmt.Println(a);
	fmt.Println(b);

	_,c := mutliReturn();
	fmt.Println(c);
}