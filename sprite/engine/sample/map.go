package sample
import "fmt"
func TestMap() {
	m := make(map[string] int)
	for i:=0;i<10;i++ {
		var a=fmt.Sprintf("apple%d",i)
		m[a]=i*100
	}

	for k,v  := range m {
		fmt.Printf("key=%s value=%d\n", k,v)
	}
}