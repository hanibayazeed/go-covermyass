package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

/**
 * ThrowError	Throws an error and exit the tool.
 */
func ThrowError(err string) {
	fmt.Println(err)
	os.Exit(1)
}

/**
 * ReadInput	Function that reads the user input.
 */
func ReadInput(separator string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(separator)
	input, _ := reader.ReadString('\n')

	return strings.TrimRight(input, "\n")
}

/**
 * isWritable	Function that tells you if a
 * file is writable or not.
 */
func isWritable(path string) bool {
	return unix.Access(path, unix.W_OK) == nil
}
