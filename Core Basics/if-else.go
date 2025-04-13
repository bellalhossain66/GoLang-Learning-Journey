package main;
import "fmt";

func main() {
	if(8%2 == 0){
		fmt.Println("8 is even");
	}else{
		fmt.Println("8 is odd");
	}

	if(8%2 == 0 || 9%2 == 0){
		fmt.Println("Either 8 or 9 is even");
	}

	if number := 8; number < 0 {
		fmt.Println("Number is negetive");
	}else if number < 10 {
		fmt.Println("Number is 1 digit");
	}else {
		fmt.Println("Number is multiple digit");
	}
}