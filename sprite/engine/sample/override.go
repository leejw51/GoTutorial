package sample

import "fmt"

type Getter interface {
	Get() string
}

type Drink struct {
	Getter
}

func (base *Drink) Get() string {
	return "drink"
}

func (base *Drink) GetName() string {
	return base.Getter.Get()
}

//////////////////////////////////////////////////////////////////
type AppleCider struct {
	Drink
}

func (t *AppleCider) Get() string {
	return "AppleCider"
}
func New() *AppleCider {
	userType := &AppleCider{}
	//userType.Getter = interface{}(userType).(Getter)
	userType.Getter = userType
	return userType
}

func TestOverride() {
	userType := New()
	fmt.Println(userType.GetName()) // user string
}
