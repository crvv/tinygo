package main

/*
int add(int x, int y) {
    return x + y;
}
*/
import "C"

func main() {
    sum := C.add(1, 2)
    println(sum)
}
