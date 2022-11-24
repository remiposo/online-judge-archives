package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	mio := NewMyIO(os.Stdin, os.Stdout)
	defer mio.Flush()

	l1 := mio.MustScanIntN(2)
	h, w := l1[0], l1[1]
	x := make([][]int, h+1)
	for i := 0; i < h+1; i++ {
		x[i] = make([]int, w+1)
	}

	for i := 1; i <= h; i++ {
		r := mio.MustScanIntN(w)
		for j := 1; j <= w; j++ {
			x[i][j] = x[i][j-1] + r[j-1]
		}
	}

	for i := 1; i <= w; i++ {
		for j := 1; j <= h; j++ {
			x[j][i] = x[j-1][i] + x[j][i]
		}
	}

	q := mio.MustScanInt()
	for i := 0; i < q; i++ {
		l2 := mio.MustScanIntN(4)
		a, b, c, d := l2[0], l2[1], l2[2], l2[3]
		s := x[c][d] + x[a-1][b-1] - x[c][b-1] - x[a-1][d]
		mio.Printf("%d\n", s)
	}
}

//#--------------------------------------------------#
//# io                                               #
//#--------------------------------------------------#

type MyIO struct {
	Scanner   *bufio.Scanner
	BufWriter *bufio.Writer
}

func NewMyIO(r io.Reader, w io.Writer) *MyIO {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	// extend maximum size of buffer to scan long token
	sc.Buffer(make([]byte, 0, 4096), math.MaxInt64)

	bw := bufio.NewWriter(os.Stdout)
	return &MyIO{
		Scanner:   sc,
		BufWriter: bw,
	}
}

func (mio *MyIO) Flush() error {
	return mio.BufWriter.Flush()
}

func (mio *MyIO) ScanStr() (string, error) {
	if !mio.Scanner.Scan() {
		return "", errors.New("EOF or an error")
	}
	return mio.Scanner.Text(), nil
}

func (mio *MyIO) MustScanStr() string {
	s, err := mio.ScanStr()
	if err != nil {
		panic(err)
	}
	return s
}

func (mio *MyIO) ScanStrN(n int) ([]string, error) {
	ss := make([]string, n)
	for idx := 0; idx < n; idx++ {
		s, err := mio.ScanStr()
		if err != nil {
			return nil, err
		}
		ss[idx] = s
	}
	return ss, nil
}

func (mio *MyIO) MustScanStrN(n int) []string {
	ss, err := mio.ScanStrN(n)
	if err != nil {
		panic(err)
	}
	return ss
}

func (mio *MyIO) ScanBytes() ([]byte, error) {
	s, err := mio.ScanStr()
	return []byte(s), err
}

func (mio *MyIO) MustScanBytes() []byte {
	b, err := mio.ScanBytes()
	if err != nil {
		panic(err)
	}
	return b
}

func (mio *MyIO) ScanBytesN(n int) ([][]byte, error) {
	bs := make([][]byte, n)
	for idx := 0; idx < n; idx++ {
		b, err := mio.ScanBytes()
		if err != nil {
			return nil, err
		}
		bs[idx] = b
	}
	return bs, nil
}

func (mio *MyIO) MustScanBytesN(n int) [][]byte {
	bs, err := mio.ScanBytesN(n)
	if err != nil {
		panic(err)
	}
	return bs
}

func (mio *MyIO) ScanInt() (int, error) {
	s, err := mio.ScanStr()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(s)
}

func (mio *MyIO) MustScanInt() int {
	i, err := mio.ScanInt()
	if err != nil {
		panic(err)
	}
	return i
}

func (mio *MyIO) ScanIntN(n int) ([]int, error) {
	result := make([]int, n)
	for idx := 0; idx < n; idx++ {
		i, err := mio.ScanInt()
		if err != nil {
			return nil, err
		}
		result[idx] = i
	}
	return result, nil
}

func (mio *MyIO) MustScanIntN(n int) []int {
	is, err := mio.ScanIntN(n)
	if err != nil {
		panic(err)
	}
	return is
}

func (mio *MyIO) Printf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(mio.BufWriter, format, a...)
}
