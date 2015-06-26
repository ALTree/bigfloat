package floatutils

import "testing"

var want string = "3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706798"

func TestPi(t *testing.T) {
	for _, test := range []struct {
		prec           uint
		decimal_digits int
	}{
		{50, 15},
		{100, 30},
		{200, 60},
		{300, 90},
	} {
		pi := Pi(test.prec)

		if pi.Text('f', test.decimal_digits) != want[:test.decimal_digits+2] {
			t.Errorf("pi(%d)\ngot  %s\nwant %s",
				test.prec,
				pi.Text('f', test.decimal_digits),
				want[:test.decimal_digits+2])
		}
	}

}
