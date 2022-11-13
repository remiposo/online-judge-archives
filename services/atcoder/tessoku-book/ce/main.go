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

	n := mio.MustScanInt()
	as := mio.MustScanIntN(n)
	sy := make([]int, n+1)
	sn := make([]int, n+1)
	for idx, a := range as {
		if a == 1 {
			sy[idx+1] = sy[idx] + 1
			sn[idx+1] = sn[idx]
		} else {
			sy[idx+1] = sy[idx]
			sn[idx+1] = sn[idx] + 1
		}
	}
	q := mio.MustScanInt()
	for i := 0; i < q; i++ {
		line := mio.MustScanIntN(2)
		l, r := line[0], line[1]
		cy := sy[r] - sy[l-1]
		cn := sn[r] - sn[l-1]
		if cy > cn {
			mio.Printf("win\n")
		} else if cy < cn {
			mio.Printf("lose\n")
		} else {
			mio.Printf("draw\n")
		}
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
