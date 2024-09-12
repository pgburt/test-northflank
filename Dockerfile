FROM cgr.dev/chainguard/go:latest as build

WORKDIR /work

COPY <<EOF go.mod
module hello
go 1.19
EOF

COPY <<EOF main.gopackage main
import (
	"math/big"
	"time"
	"fmt"
)
func calcDenominator(i int) int64 {
	// Calc the denominator
	// The value should be the nth odd number
	// For ex: if n = 4, then return 7.
	return int64(i * 2 - 1)
}
func findSign(i int, d int64) *big.Rat {
	if i % 2 == 0 {
		return big.NewRat(-1, d)
	} else {
		return big.NewRat(1, d)
	}	
}
func sleepyTime() {
	time.Sleep(10 * time.Millisecond)
}
func main() {
	// Rat is for Rational Numbers of arbitrary length
	sum := new(big.Rat)
	r := new(big.Rat)
	// The Leibniz caculation for Pi follows
	// (pi/4) = 1 - (1/3) + (1/5) - (1/7) ...
	for i := 1; i <= 1000; i++ {
		d := calcDenominator(i)	
		r = findSign(i, d)
		sum = sum.Add(sum, r) 
		sleepyTime()
	}
	// Change pi/4 to just pi by multiplying answer by 4
	// FloatString asks for an amount of precision.
	fmt.Println(sum.FloatString(32))
}
EOF
RUN go build -o hello .

FROM cgr.dev/chainguard/static:latest

COPY --from=build /work/hello /hello
CMD ["/hello"]
