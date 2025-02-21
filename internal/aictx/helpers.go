package aictx

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	// KB is number of bytes in Kilobyte.
	KB = 1024
	// MB is number of bytes in Megabyte.
	MB = 1024 * KB
)

// fileHeader renders a header for each file.
// fileHeader renders a header for each file.
// It now includes a file counter (e.g. "[1/6]" or "[01/12]") inserted into a 60-char line.
func fileHeader(node *TreeNode, fileNum, totalFiles int) []byte {
	const totalLen = 60 // total characters (without the newline)
	var buf bytes.Buffer

	// Determine padding width: if totalFiles is single digit then width=1,
	// if <100 then width=2, if <1000 then width=3, etc.
	width := len(strconv.Itoa(totalFiles))
	// Create the number info string with padded file number.
	numInfo := fmt.Sprintf("[%0*d/%d]", width, fileNum, totalFiles)
	// Calculate remaining space and split evenly on left/right.
	rem := totalLen - len(numInfo)
	left := rem / 2 //nolint: mnd // 2 for half
	right := rem - left
	headerLine := strings.Repeat("=", left) + numInfo + strings.Repeat("=", right) + "\n"

	buf.WriteString(headerLine)
	buf.WriteString(fmt.Sprintf("File: %s\n", node.Path))
	if node.Size > 0 {
		buf.WriteString(fmt.Sprintf("Size: %s\n", formatSize(node.Size)))
	}
	buf.WriteString(strings.Repeat("-", totalLen) + "\n")
	return buf.Bytes()
}

// isHidden returns true if the provided file or folder name starts with a dot.
func isHidden(name string) bool {
	return len(name) > 0 && name[0] == '.'
}

// isBinary returns true if the provided data appears to be binary.
func isBinary(data []byte) bool {
	if len(data) == 0 {
		return false // empty files are considered text
	}
	for _, b := range data {
		if b == 0 {
			return true
		}
	}
	return !utf8.Valid(data)
}

// formatSize converts bytes to a human-friendly string.
func formatSize(bytes int64) string {
	if bytes >= MB {
		return fmt.Sprintf("%.2f MB", float64(bytes)/float64(MB))
	} else if bytes >= KB {
		return fmt.Sprintf("%.2f KB", float64(bytes)/float64(KB))
	}
	return fmt.Sprintf("%d B", bytes)
}

// exceedsThreshold returns true if the file size (in bytes) exceeds the threshold (in MB).
func exceedsThreshold(sizeBytes int64, thresholdMb float64) bool {
	mb := float64(sizeBytes) / (KB * KB)
	return mb > thresholdMb
}
