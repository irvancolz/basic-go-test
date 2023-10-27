package packageusage

import "testing"

func TestPenyewaanBola(t *testing.T) {
	PenyewaanBola()
}

func TestBelanja(t *testing.T) {
	rinabelanja()
}

func TestUrutan(t *testing.T) {
	urutkanangka()
}

func TestString(t *testing.T) {
	merubahstring()
}

func TestWaktu(t *testing.T) {
	waktu()
}

func TestFactorial(t *testing.T) {
	t.Log(getFactorial(0))
}
