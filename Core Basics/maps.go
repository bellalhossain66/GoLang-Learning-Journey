package main;
import (
	"fmt"
	"maps"
);

func main() {
	a := make(map[string]int);
	fmt.Println("Empty: ", a);

	a["k1"] = 10;
	a["k2"] = 20;
	fmt.Println("Map: ", a);

	v1 := a["k1"];
	fmt.Println("v1: ", v1);

	v3 := a["k3"];
	fmt.Println("v3: ", v3);

	fmt.Println("Length: ", len(a));

	delete(a, "k2");
	fmt.Println("Map: ", a);

	clear(a);
	fmt.Println("Map: ", a);

	_,exist := a["k2"];  //value,exist
	fmt.Println("Exist: ", exist);

	x := map[string]int{ "v1": 1, "v2": 2 };
	fmt.Println("Map: ", x);

	y := map[string]int{ "v1": 1, "v2": 2, "v3": 3 };
	fmt.Println("Map: ", y);

	if maps.Equal(x,y) {
		fmt.Println("x == y");
	}else{
		fmt.Println("x != y");
	}
}