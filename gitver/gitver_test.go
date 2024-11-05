package gitver

import "fmt"

func Example_semVer_build_meta() {
	fmt.Println(semVer("v1.5.0+something1"))
	fmt.Println(semVer("v1.5.0"))

	// Output:
	// 1.5.0+something1 <nil>
	// 1.5.0 <nil>
}
