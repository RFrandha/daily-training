package hackerrank

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'largestRectangle' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts INTEGER_ARRAY h as parameter.
 */

func largestRectangle(h []int32) int64 {
	finalMax := int32(0)
	for i := 0; i < len(h); i++ {
		curArea := h[i]
		for j := i + 1; j < len(h); j++ {
			if h[j] >= h[i] {
				curArea += h[i]
			} else {
				break
			}
		}
		for k := i - 1; k >= 0; k-- {
			if h[k] >= h[i] {
				curArea += h[i]
			} else {
				break
			}
		}
		if curArea > finalMax {
			finalMax = curArea
		}
	}
	return int64(finalMax)
}

func HRLargestRect() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	hTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var h []int32

	for i := 0; i < int(n); i++ {
		hItemTemp, err := strconv.ParseInt(hTemp[i], 10, 64)
		checkError(err)
		hItem := int32(hItemTemp)
		h = append(h, hItem)
	}

	result := largestRectangle(h)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
