package hash

import (
	"fmt"
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

func Print(name string, m map[uint32][]string, isPrint bool) {
	s := fmt.Sprintf("%10s->", name)
	cnt := 0
	collideCnt := 0
	collideSum := 0
	collideAvg := 0
	i := 0
	for k, vs := range m {
		i++
		n := len(m[k])
		if isPrint {
			s += fmt.Sprintf("%2d,", n)
		}
		if i > 0 && i%50 == 0 {
			if isPrint {
				fmt.Println(s)
			}
			s = s[:12]
		}
		if n > 10 {
			collideCnt++
			collideSum += n
			if isPrint {
				fmt.Printf("%s=>%s\n", name, vs)
			}
		}
		cnt += n
	}
	if collideCnt > 0 {
		collideAvg = collideSum / collideCnt
	}
	fmt.Printf("\t==>%10s: Size:%5d, count:%5d, collideCnt:%5d ,collideSum:%5d, collideAvg:%5d\n",
		name, len(m), cnt, collideCnt, collideSum, collideAvg)
}

func testHashMothed(name string, f func(s string) uint32, size uint32, isPrint bool) {
	m := make(map[uint32][]string)
	for i := uint32(0); i < size; i++ {
		key := fmt.Sprintf("foo.%d", i)
		hash := f(key)
		m[hash%size] = append(m[hash%size], key)
	}
	Print(name, m, isPrint)
}

func TestBernsteinDistribution(t *testing.T) {
	testHashMothed("BernsteinS", BernsteinS, 10000, false)
}

func TestMurmur3Distribution(t *testing.T) {
	testHashMothed("Murmur3S", func(s string) uint32 {
		return Murmur3S(s, M3Seed)
	}, 10000, false)
}

func TestFNV1ADistribution(t *testing.T) {
	testHashMothed("FNV1AS", FNV1AS, 10000, false)
}

func TestWukehongDistribution(t *testing.T) {
	testHashMothed("WukehongS", WukehongS, 10000, false)
}

func TestMeiyanDistribution(t *testing.T) {
	testHashMothed("MeiyanS", MeiyanS, 10000, false)
}

func TestJesteressDistribution(t *testing.T) {
	testHashMothed("JesteressS", JesteressS, 10000, false)
}

func TestYorikkeDistribution(t *testing.T) {
	testHashMothed("YorikkeS", YorikkeS, 10000, false)
}

var ch = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@$#%^&*()")

func sizedBytes(minSize, maxSize int) [][]byte {
	sz := maxSize - minSize
	b := make([][]byte, sz)
	for i := range b {
		b[i] = make([]byte, minSize+i)
		for j := range b[i] {
			b[i][j] = ch[rand.Intn(len(ch))]
		}
	}
	return b
}

var smlKey = sizedBytes(1, 16)
var medKey = sizedBytes(32, 47)
var lrgKey = sizedBytes(256, 271)

func Benchmark_Bernstein_SmallKey(b *testing.B) {
	for i := 0; i < b.N; i += len(smlKey) {
		for j := range smlKey {
			Bernstein(smlKey[j])
		}
	}
}

func Benchmark_Murmur3___SmallKey(b *testing.B) {
	for i := 0; i < b.N; i += len(smlKey) {
		for j := range smlKey {
			Murmur3(smlKey[j], M3Seed)
		}
	}
}

func Benchmark_FNV1A_____SmallKey(b *testing.B) {
	for i := 0; i < b.N; i += len(smlKey) {
		for j := range smlKey {
			FNV1A(smlKey[j])
		}
	}
}

func Benchmark_Wukehong__SmallKey(b *testing.B) {
	for i := 0; i < b.N; i += len(smlKey) {
		for j := range smlKey {
			Wukehong(smlKey[j])
		}
	}
}

func Benchmark_Meiyan____SmallKey(b *testing.B) {
	for i := 0; i < b.N; i += len(smlKey) {
		for j := range smlKey {
			Meiyan(smlKey[j])
		}
	}
}

func Benchmark_Jesteress_SmallKey(b *testing.B) {
	for i := 0; i < b.N; i += len(smlKey) {
		for j := range smlKey {
			Jesteress(smlKey[j])
		}
	}
}

func Benchmark_Yorikke___SmallKey(b *testing.B) {
	for i := 0; i < b.N; i += len(smlKey) {
		for j := range smlKey {
			Yorikke(smlKey[j])
		}
	}
}

func Benchmark_Bernstein___MedKey(b *testing.B) {
	for i := 0; i < b.N; i += len(medKey) {
		for j := range smlKey {
			Bernstein(medKey[j])
		}
	}
}

func Benchmark_Murmur3_____MedKey(b *testing.B) {
	for i := 0; i < b.N; i += len(medKey) {
		for j := range medKey {
			Murmur3(medKey[j], M3Seed)
		}
	}
}

func Benchmark_FNV1A_______MedKey(b *testing.B) {
	for i := 0; i < b.N; i += len(medKey) {
		for j := range medKey {
			FNV1A(medKey[j])
		}
	}
}

func Benchmark_Wukehong____MedKey(b *testing.B) {
	for i := 0; i < b.N; i += len(medKey) {
		for j := range medKey {
			Wukehong(medKey[j])
		}
	}
}

func Benchmark_Meiyan______MedKey(b *testing.B) {
	for i := 0; i < b.N; i += len(medKey) {
		for j := range medKey {
			Meiyan(medKey[j])
		}
	}
}

func Benchmark_Jesteress___MedKey(b *testing.B) {
	for i := 0; i < b.N; i += len(medKey) {
		for j := range medKey {
			Jesteress(medKey[j])
		}
	}
}

func Benchmark_Yorikke_____MedKey(b *testing.B) {
	for i := 0; i < b.N; i += len(medKey) {
		for j := range medKey {
			Yorikke(medKey[j])
		}
	}
}

func Benchmark_Bernstein___LrgKey(b *testing.B) {
	for i := 0; i < b.N; i += len(lrgKey) {
		for j := range lrgKey {
			Bernstein(lrgKey[j])
		}
	}
}

func Benchmark_Murmur3_____LrgKey(b *testing.B) {
	for i := 0; i < b.N; i += len(lrgKey) {
		for j := range lrgKey {
			Murmur3(lrgKey[j], M3Seed)
		}
	}
}

func Benchmark_FNV1A_______LrgKey(b *testing.B) {
	for i := 0; i < b.N; i += len(lrgKey) {
		for j := range lrgKey {
			FNV1A(lrgKey[j])
		}
	}
}

func Benchmark_Wukehong____LrgKey(b *testing.B) {
	for i := 0; i < b.N; i += len(lrgKey) {
		for j := range lrgKey {
			Wukehong(lrgKey[j])
		}
	}
}

func Benchmark_Meiyan______LrgKey(b *testing.B) {
	for i := 0; i < b.N; i += len(lrgKey) {
		for j := range lrgKey {
			Meiyan(lrgKey[j])
		}
	}
}

func Benchmark_Jesteress___LrgKey(b *testing.B) {
	for i := 0; i < b.N; i += len(lrgKey) {
		for j := range lrgKey {
			Jesteress(lrgKey[j])
		}
	}
}

func Benchmark_Yorikke_____LrgKey(b *testing.B) {
	for i := 0; i < b.N; i += len(lrgKey) {
		for j := range lrgKey {
			Yorikke(lrgKey[j])
		}
	}

}

func Benchmark_ToSlice(b *testing.B) {
	s := "12234567890"
	for i := 0; i < b.N; i++ {
		toSlice(s)
	}
}

func Benchmark_ToString(b *testing.B) {
	bs := []byte("12234567890")
	for i := 0; i < b.N; i++ {
		toString(bs)
	}
}
func ToString(b [][]byte) []string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = toString(b[i])
	}
	return s
}

var smlKeyS = ToString(smlKey)
var medKeyS = ToString(medKey)
var lrgKeyS = ToString(lrgKey)

func Benchmark_BernsteinS_SmallKey(b *testing.B) {
	for i := 0; i < b.N; i += len(smlKeyS) {
		for j := range smlKeyS {
			BernsteinS(smlKeyS[j])
		}
	}
}

func Benchmark_Murmur3S___SmallKey(b *testing.B) {
	for i := 0; i < b.N; i += len(smlKeyS) {
		for j := range smlKeyS {
			Murmur3S(smlKeyS[j], M3Seed)
		}
	}
}

func Benchmark_FNV1AS_____SmallKey(b *testing.B) {
	for i := 0; i < b.N; i += len(smlKeyS) {
		for j := range smlKeyS {
			FNV1AS(smlKeyS[j])
		}
	}
}

func Benchmark_WukehongS__SmallKey(b *testing.B) {
	for i := 0; i < b.N; i += len(smlKeyS) {
		for j := range smlKeyS {
			WukehongS(smlKeyS[j])
		}
	}

}

func Benchmark_MeiyanS____SmallKey(b *testing.B) {
	for i := 0; i < b.N; i += len(smlKeyS) {
		for j := range smlKeyS {
			MeiyanS(smlKeyS[j])
		}
	}
}

func Benchmark_JesteressS_SmallKey(b *testing.B) {
	for i := 0; i < b.N; i += len(smlKeyS) {
		for j := range smlKeyS {
			JesteressS(smlKeyS[j])
		}
	}
}

func Benchmark_YorikkeS___SmallKey(b *testing.B) {
	for i := 0; i < b.N; i += len(smlKeyS) {
		for j := range smlKeyS {
			YorikkeS(smlKeyS[j])
		}
	}
}

func Benchmark_BernsteinS___MedKey(b *testing.B) {
	for i := 0; i < b.N; i += len(medKeyS) {
		for j := range medKeyS {
			BernsteinS(medKeyS[j])
		}
	}
}

func Benchmark_Murmur3S_____MedKey(b *testing.B) {
	for i := 0; i < b.N; i += len(medKeyS) {
		for j := range medKeyS {
			Murmur3S(medKeyS[j], M3Seed)
		}
	}
}

func Benchmark_FNV1AS_______MedKey(b *testing.B) {
	for i := 0; i < b.N; i += len(medKeyS) {
		for j := range medKeyS {
			FNV1AS(medKeyS[j])
		}
	}
}

func Benchmark_WukehongS____MedKey(b *testing.B) {
	for i := 0; i < b.N; i += len(medKeyS) {
		for j := range medKeyS {
			WukehongS(medKeyS[j])
		}
	}
}

func Benchmark_MeiyanS______MedKey(b *testing.B) {
	for i := 0; i < b.N; i += len(medKeyS) {
		for j := range medKeyS {
			MeiyanS(medKeyS[j])
		}
	}
}

func Benchmark_JesteressS___MedKey(b *testing.B) {
	for i := 0; i < b.N; i += len(medKeyS) {
		for j := range medKeyS {
			JesteressS(medKeyS[j])
		}
	}
}

func Benchmark_YorikkeS_____MedKey(b *testing.B) {
	for i := 0; i < b.N; i += len(medKeyS) {
		for j := range medKeyS {
			YorikkeS(medKeyS[j])
		}
	}
}

func Benchmark_BernsteinS___LrgKey(b *testing.B) {
	for i := 0; i < b.N; i += len(lrgKeyS) {
		for j := range lrgKeyS {
			BernsteinS(lrgKeyS[j])
		}
	}
}

func Benchmark_Murmur3S_____LrgKey(b *testing.B) {
	for i := 0; i < b.N; i += len(lrgKeyS) {
		for j := range lrgKeyS {
			Murmur3S(lrgKeyS[j], M3Seed)
		}
	}
}

func Benchmark_FNV1AS_______LrgKey(b *testing.B) {
	for i := 0; i < b.N; i += len(lrgKeyS) {
		for j := range lrgKeyS {
			FNV1AS(lrgKeyS[j])
		}
	}
}

func Benchmark_WukehongS____LrgKey(b *testing.B) {
	for i := 0; i < b.N; i += len(lrgKeyS) {
		for j := range lrgKeyS {
			WukehongS(lrgKeyS[j])
		}
	}
}

func Benchmark_MeiyanS______LrgKey(b *testing.B) {
	for i := 0; i < b.N; i += len(lrgKeyS) {
		for j := range lrgKeyS {
			MeiyanS(lrgKeyS[j])
		}
	}
}

func Benchmark_JesteressS___LrgKey(b *testing.B) {
	for i := 0; i < b.N; i += len(lrgKeyS) {
		for j := range lrgKeyS {
			JesteressS(lrgKeyS[j])
		}
	}
}

func Benchmark_YorikkeS_____LrgKey(b *testing.B) {
	for i := 0; i < b.N; i += len(lrgKeyS) {
		for j := range lrgKeyS {
			YorikkeS(lrgKeyS[j])

		}
	}
}
