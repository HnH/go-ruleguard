package extra

import "fmt"

func testFormatInt() {
	{
		x16 := int16(342)
		_ = fmt.Sprintf("%d", x16) // want `use strconv.FormatInt\(int64\(x16\), 10\)`
	}
	{
		x64 := int64(32)
		_ = fmt.Sprintf("%d", x64) // want `use strconv.FormatInt\(x64, 10\)`
	}
	{
		// Check that convertibleTo(int64) condition works and rejects this.
		s := struct{}{}
		_ = fmt.Sprintf("%d", s)
	}
}

func testFormatBool() {
	{
		i := int64(4)
		_ = fmt.Sprintf("%t", (i+i)&1 == 0) // want `use strconv.FormatBool\(\(i \+ i\)&1 == 0\)`
	}
}

func testBlankAssign() {
	x := foo()
	_ = x // want `please remove the assignment to _`

	// This is OK, could be for side-effects.
	_ = foo()
}

func nilErrCheck() {
	if mightFail() == nil { // want `assign mightFail\(\) to err and then do a nil check`
	}
	if mightFail() != nil { // want `assign mightFail\(\) to err and then do a nil check`
	}

	// Good.
	if err := mightFail(); err != nil {
	}
	err := mightFail()
	if err == nil {
	}

	// Not error-typed LHS.
	if newInt() == nil {
	}
}