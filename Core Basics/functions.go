package main;
import "fmt";

func plus(a int, b int) int {
	return a + b;
}

func plusPlus(a, b, c int) int {
	return a + b + c;
}

func main() {
	responce := plus(5, 10);
	fmt.Println("Responce: ", responce);


	responce = plusPlus(5, 10, 20);
	fmt.Println("Responce: ", responce);
}
