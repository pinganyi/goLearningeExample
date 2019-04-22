package main

import "errors"
import "fmt"

func f1(arg int) (int , error){
	if arg == 42 {
		return -1, errors.New("can work")
	}
	return arg + 3,nil
}

type argError struct {
    arg  int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
func f2(arg int)(int,error){
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3,nil
}

func main() {

    // The two loops below test out each of our
    // error-returning functions. Note that the use of an
    // inline error check on the `if` line is a common
    // idiom in Go code.
    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

    // If you want to programmatically use the data in
    // a custom error, you'll need to get the error as an
    // instance of the custom error type via type
    // assertion.
    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}