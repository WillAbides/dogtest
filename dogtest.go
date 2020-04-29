package dogtest

import (
	"fmt"
	"io/ioutil"
)

func Foo() {
	fmt.Println("lint should get me")
	tf, _ := ioutil.TempFile("", "")
	defer tf.Close()
}
