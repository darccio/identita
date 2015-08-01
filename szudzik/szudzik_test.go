package szudzik

import (
	"testing"
)

func TestSzudzik(t *testing.T) {
	const x, y, z = 32767.0, 32767.0, 1073741823.0
	out := ElegantPair(x, y)
	if z != out {
		t.Errorf("ElegantPair(%v, %v) = %v, want %v", x, y, out, z)
	}
	if outX, outY := ElegantUnpair(out); outX != x || outY != y {
		t.Errorf("ElegantPair(%v) = %v, %v; want %v, %v", out, outX, outY, x, y)
	}
}

func TestMonthDaySzudzik(t *testing.T) {
	const month, day, z = 12, 31, 973
	out := ElegantPair(month, day)
	if z != out {
		t.Errorf("ElegantPair(%v, %v) = %v, want %v", month, day, out, z)
	}
	if outMonth, outDay := ElegantUnpair(out); outMonth != month || outDay != day {
		t.Errorf("ElegantPair(%v) = %v, %v; want %v, %v", out, outMonth, outDay, month, day)
	}
}

func TestChainedSzudzik(t *testing.T) {
	const c, y, m = 3, 99, 12
	p1 := ElegantPair(c, y)
	t.Log("First pair:", p1)
	p2 := ElegantPair(m, p1)
	t.Log("Second pair:", p2)
	om, op := ElegantUnpair(p2)
	if om != m || op != p1 {
		t.Errorf("ElegantUnpair(%v) = %v, %v; want %v, %v", p2, om, op, m, p1)
	}
	oc, oy := ElegantUnpair(op)
	if oc != c || oy != y {
		t.Errorf("ElegantUnpair(%v) = %v, %v; want %v, %v", op, oc, oy, c, y)
	}
}
