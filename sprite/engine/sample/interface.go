package sample
import "fmt"
type Food interface {
	price() int64
}
type Apple struct {
	appleKind int64
}
func (a Apple) price() int64 {
	return a.appleKind*1000
}
type Ramen struct {
	ramenKind int64
}
func (r Ramen) price() int64 {
	return r.ramenKind*20000
}
func TestInterface() {
	fmt.Println("test interface")
	a := Apple {appleKind:1}
	r := Ramen {ramenKind:2}
	fmt.Printf("apple %v\n", a)
	fmt.Printf("ramen %v\n", r)
}