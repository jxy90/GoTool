package anyX

import (
	"fmt"
	"testing"
)

func TestAnyX(t *testing.T) {
	fmt.Println(MatchAny("123.456", 123.456))
	fmt.Println(MatchAny("123", 123))
	fmt.Println(MatchAny("123", int8(123)))
	fmt.Println(MatchAny("test", int8(123)))
	fmt.Println(MatchAny("test", "test"))
	fmt.Println(MatchAny("test", "test\n"))
	fmt.Println(MatchAny("test", "test "))
	fmt.Println(MatchAny("true", true))
	fmt.Println(MatchAny("false", true))
	fmt.Println(MatchAny("false", false))
}
