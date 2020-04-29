package dogtest

import (
	"io/ioutil"


	"fmt"

)

func Foo() {
	fmt.Println("lint should get me")
	tf, _ := ioutil.TempFile("", "")
	defer tf.Close()
}
