package hello
import (
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)
func Hello() string {
//	return "Hello"
	return quote.Hello()
}
func Proverb() string {
	return quoteV3.Concurrency()
}
