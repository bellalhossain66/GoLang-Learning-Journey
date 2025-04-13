package main;
import "fmt";

func summation(nums ...int){
	fmt.Print(nums, " ");
	total := 0;
	for _,nums:= range nums {
		total += nums;
	}
	fmt.Println(total);
}

func main() {
	summation(2, 2);
	summation(1,2,3,4,5);

	list := []int{1,5,9,18};
	summation(list...);
}