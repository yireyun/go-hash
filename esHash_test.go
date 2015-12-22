package hash

import (
	"math/rand"
	"testing"
)

var keys = [][]byte{
	[]byte("foo"),
	[]byte("bar"),
	[]byte("apcera.continuum.router.foo.bar"),
	[]byte("apcera.continuum.router.foo.bar.baz"),
}

var keyS = []string{
	"foo",
	"bar",
	"apcera.continuum.router.foo.bar",
	"apcera.continuum.router.foo.bar.baz",
}

func TestBernstein(t *testing.T) {
	results := []uint32{193491849, 193487034, 2487287557, 3139297488}
	for i, key := range keys {
		h := Bernstein(key)
		if h != results[i] {
			t.Fatalf("hash is incorrect, expected %d, got %d\n",
				results[i], h)
		}
	}
}

func TestBernsteinS(t *testing.T) {
	results := []uint32{193491849, 193487034, 2487287557, 3139297488}
	for i, key := range keyS {
		h := BernsteinS(key)
		if h != results[i] {
			t.Fatalf("hash is incorrect, expected %d, got %d\n",
				results[i], h)
		}
	}
}

func TestFNV1A(t *testing.T) {
	results := []uint32{2851307223, 1991736602, 1990810707, 1244015104}
	for i, key := range keys {
		h := FNV1A(key)
		if h != results[i] {
			t.Fatalf("hash is incorrect, expected %d, got %d\n",
				results[i], h)
		}
	}
}

func TestFNV1AS(t *testing.T) {
	results := []uint32{2851307223, 1991736602, 1990810707, 1244015104}
	for i, key := range keyS {
		h := FNV1AS(key)
		if h != results[i] {
			t.Fatalf("hash is incorrect, expected %d, got %d\n",
				results[i], h)
		}
	}
}
func TestJesteress(t *testing.T) {
	results := []uint32{1058908168, 1061739001, 4242539713, 3332038527}
	for i, key := range keys {
		h := Jesteress(key)
		if h != results[i] {
			t.Fatalf("hash is incorrect, expected %d, got %d\n",
				results[i], h)
		}
	}
}

func TestJesteressS(t *testing.T) {
	results := []uint32{1058908168, 1061739001, 4242539713, 3332038527}
	for i, key := range keyS {
		h := JesteressS(key)
		if h != results[i] {
			t.Fatalf("hash is incorrect, expected %d, got %d\n",
				results[i], h)
		}
	}
}
func TestMeiyan(t *testing.T) {
	results := []uint32{1058908168, 1061739001, 2891236487, 3332038527}
	for i, key := range keys {
		h := Meiyan(key)
		if h != results[i] {
			t.Fatalf("hash is incorrect, expected %d, got %d\n",
				results[i], h)
		}
	}
}

func TestMeiyanS(t *testing.T) {
	results := []uint32{1058908168, 1061739001, 2891236487, 3332038527}
	for i, key := range keyS {
		h := MeiyanS(key)
		if h != results[i] {
			t.Fatalf("hash is incorrect, expected %d, got %d\n",
				results[i], h)
		}
	}
}
func TestYorikke(t *testing.T) {
	results := []uint32{3523423968, 2222334353, 407908456, 359111667}
	for i, key := range keys {
		h := Yorikke(key)
		if h != results[i] {
			t.Fatalf("hash is incorrect, expected %d, got %d\n",
				results[i], h)
		}
	}
}

func TestYorikkeS(t *testing.T) {
	results := []uint32{3523423968, 2222334353, 407908456, 359111667}
	for i, key := range keyS {
		h := YorikkeS(key)
		if h != results[i] {
			t.Fatalf("hash is incorrect, expected %d, got %d\n",
				results[i], h)
		}
	}
}

func TestMurmur3(t *testing.T) {
	results := []uint32{659908353, 522989004, 135963738, 990328005}
	for i, key := range keys {
		h := Murmur3(key, M3Seed)
		if h != results[i] {
			t.Fatalf("hash is incorrect, expected %d, got %d\n",
				results[i], h)
		}
	}
}

func TestMurmur3S(t *testing.T) {
	results := []uint32{659908353, 522989004, 135963738, 990328005}
	for i, key := range keyS {
		h := Murmur3S(key, M3Seed)
		if h != results[i] {
			t.Fatalf("hash is incorrect, expected %d, got %d\n",
				results[i], h)
		}
	}
}

var ch = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@$#%^&*()")

func sizedBytes(sz int) []byte {
	b := make([]byte, sz)
	for i := range b {
		b[i] = ch[rand.Intn(len(ch))]
	}
	return b
}

var smlKey = sizedBytes(8)
var medKey = sizedBytes(32)
var lrgKey = sizedBytes(256)

func Benchmark_Bernstein_SmallKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bernstein(smlKey)
	}
}

func Benchmark_Murmur3___SmallKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Murmur3(smlKey, M3Seed)
	}
}

func Benchmark_FNV1A_____SmallKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FNV1A(smlKey)
	}
}

func Benchmark_Wukehong__SmallKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Wukehong(smlKey)
	}
}

func Benchmark_Meiyan____SmallKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Meiyan(smlKey)
	}
}

func Benchmark_Jesteress_SmallKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Jesteress(smlKey)
	}
}

func Benchmark_Yorikke___SmallKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Yorikke(smlKey)
	}
}

func Benchmark_Bernstein___MedKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bernstein(medKey)
	}
}

func Benchmark_Murmur3_____MedKey(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Murmur3(medKey, M3Seed)
	}

}

func Benchmark_FNV1A_______MedKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FNV1A(medKey)
	}
}

func Benchmark_Wukehong____MedKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Wukehong(medKey)
	}
}

func Benchmark_Meiyan______MedKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Meiyan(medKey)
	}
}

func Benchmark_Jesteress___MedKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Jesteress(medKey)
	}
}

func Benchmark_Yorikke_____MedKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Yorikke(medKey)
	}
}

func Benchmark_Bernstein___LrgKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bernstein(lrgKey)
	}
}

func Benchmark_Murmur3_____LrgKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Murmur3(lrgKey, M3Seed)
	}
}

func Benchmark_FNV1A_______LrgKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FNV1A(lrgKey)
	}
}

func Benchmark_Wukehong____LrgKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Wukehong(lrgKey)
	}
}

func Benchmark_Meiyan______LrgKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Meiyan(lrgKey)
	}
}

func Benchmark_Jesteress___LrgKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Jesteress(lrgKey)
	}
}

func Benchmark_Yorikke_____LrgKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Yorikke(lrgKey)
	}
}

func Benchmark_ToSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toSlice("12234567890")
	}
}

func Benchmark_ToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toString(lrgKey)
	}
}

var smlKeyS = toString(smlKey)
var medKeyS = toString(medKey)
var lrgKeyS = toString(lrgKey)

func Benchmark_BernsteinS_SmallKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BernsteinS(smlKeyS)
	}
}

func Benchmark_Murmur3S___SmallKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Murmur3S(smlKeyS, M3Seed)
	}
}

func Benchmark_FNV1AS_____SmallKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FNV1AS(smlKeyS)
	}
}

func Benchmark_WukehongS__SmallKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WukehongS(smlKeyS)
	}
}

func Benchmark_MeiyanS____SmallKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MeiyanS(smlKeyS)
	}
}

func Benchmark_JesteressS_SmallKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JesteressS(smlKeyS)
	}
}

func Benchmark_YorikkeS___SmallKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YorikkeS(smlKeyS)
	}
}

func Benchmark_BernsteinS___MedKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BernsteinS(medKeyS)
	}
}

func Benchmark_Murmur3S_____MedKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Murmur3S(medKeyS, M3Seed)
	}
}

func Benchmark_FNV1AS_______MedKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FNV1AS(medKeyS)
	}
}

func Benchmark_WukehongS____MedKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WukehongS(medKeyS)
	}
}

func Benchmark_MeiyanS______MedKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MeiyanS(medKeyS)
	}
}

func Benchmark_JesteressS___MedKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JesteressS(medKeyS)
	}
}

func Benchmark_YorikkeS_____MedKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YorikkeS(medKeyS)
	}
}

func Benchmark_BernsteinS___LrgKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BernsteinS(lrgKeyS)
	}
}

func Benchmark_Murmur3S_____LrgKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Murmur3S(lrgKeyS, M3Seed)
	}
}

func Benchmark_FNV1AS_______LrgKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FNV1AS(lrgKeyS)
	}
}

func Benchmark_WukehongS____LrgKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WukehongS(lrgKeyS)
	}
}

func Benchmark_MeiyanS______LrgKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MeiyanS(lrgKeyS)
	}
}

func Benchmark_JesteressS___LrgKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JesteressS(lrgKeyS)
	}
}

func Benchmark_YorikkeS_____LrgKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YorikkeS(lrgKeyS)
	}
}
