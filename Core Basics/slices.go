package main;
import (
	"fmt"
	"slices"
);

func main() {
	var s []string;
	fmt.Println("Un INIT: ", s, s == nil, len(s) == 0);

	s = make([]string, 3);
	fmt.Println("Empty: ", s, "len: ", len(s), "cap: ", cap(s));

	s[0] = "a";
	s[1] = "b";
	s[2] = "c";
	fmt.Println("Set: ", s);
	fmt.Println("Get: ", s[1]);

	fmt.Println("Length: ", len(s));

	s = append(s, "d");
	s = append(s, "e", "f");
	fmt.Println("Append: ", s);

	c := make([]string, len(s));
	copy(c, s);
	fmt.Println("Copy: ", c);

	l := s[2:5];
	fmt.Println("Slice1: ", l);

	l = s[:5];
	fmt.Println("Slice2: ", l);

	l = s[2:];
	fmt.Println("Slice3: ", l);

	t := []string{"g", "h", "i"}
	fmt.Println("DCL: ", t);

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2){
		fmt.Println("t == t2");
	}

	twoDimentional := make([][]int, 3);
	for i := 0; i < 3; i++ {
		innerLength := i + 1;
		twoDimentional[i] = make([]int, innerLength);
		for j := 0; j < innerLength; j++ {
			twoDimentional[i][j] = i + j;
		}
	}
	fmt.Println("2d: ", twoDimentional);
}