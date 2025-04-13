package main;
import (
	"fmt"
	"time"
);

func main() {
	i := 2;
	fmt.Print("Write ", i, " as ");
	switch i {
		case 1:
			fmt.Println("one");
		case 2:
			fmt.Println("two");
		case 3:
			fmt.Println("three");
	}

	switch time.Now().Weekday() {
		case time.Saturday, time.Friday:
			fmt.Println("Its weekend");
		default:
			fmt.Println("Its weekday");
	}

	timeNow := time.Now();
	switch {
		case timeNow.Hour() < 12 :
			fmt.Println("Its befor noon");
		default :
			fmt.Println("Its after noon");
	}

	whoAmI := func( i interface{}) {
		switch t := i.(type) {
			case bool:
				fmt.Println("Its boolean");
			case int:
				fmt.Println("Its interger");
			default:
				fmt.Println("Dont know the type %T", t);
		}
	}

	whoAmI(!true);
	whoAmI(100);
	whoAmI("asdfghj");
}