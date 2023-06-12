package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// The leftRotate function performs a left rotation operation on the given slice A by n positions.
func leftRotate(A []int, n int) []int {
	// temp := make([]int, len(A))
	// copy(temp, A[n:])
	// copy(temp[len(A)-n:], A[:n])
	// return temp
	return append(A[n:], A[:n]...)
}

// The xorOfFour function performs the XOR operation on four input slices A, B, C, and D element-wise.
func xorOfFour(A, B, C, D []int) []int {
	temp := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		temp[i] = (A[i] + B[i] + C[i] + D[i]) % 2
	}
	return temp
}

// The andOp function performs the logical AND operation on two input slices A and B element-wise.
func andOp(A, B []int) []int {
	temp := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		temp[i] = A[i] & B[i]
	}
	return temp
}

// The orOfThree function performs the logical OR operation on three input slices A, B, and C element-wise.
func orOfThree(A, B, C []int) []int {
	temp := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		temp[i] = A[i] | B[i] | C[i]
	}
	return temp
}

// The notOp function performs the logical NOT operation on the input slice A element-wise.
func notOp(A []int) []int {
	temp := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		temp[i] = 1 - A[i]
	}
	return temp
}

// The addTwo function adds two input slices A and B as binary numbers.
func addTwo(A, B []int) []int {
	temp := make([]int, len(A))
	carry := 0
	for i := len(A) - 1; i >= 0; i-- {
		sum := carry + A[i] + B[i]
		temp[i] = sum % 2
		carry = sum / 2
	}
	return temp
}

// This function adds five input slices A, B, C, D, and E as binary numbers.
func addFive(A, B, C, D, E []int) []int {
	temp := make([]int, len(A))
	carry := 0
	for i := len(A) - 1; i >= 0; i-- {
		sum := carry + A[i] + B[i] + C[i] + D[i] + E[i]
		temp[i] = sum % 2
		carry = sum / 2
	}
	return temp
}

// This function converts the binary representation stored in the input slice A into a hexadecimal string.
func getDigest(A []int) string {
	s := ""
	for i := 0; i < 160; i += 4 {
		an := A[i]*8 + A[i+1]*4 + A[i+2]*2 + A[i+3]
		// if an <= 9 {
		// 	s += strconv.Itoa(an)
		// } else if an == 10 {
		// 	s += "a"
		// } else if an == 11 {
		// 	s += "b"
		// } else if an == 12 {
		// 	s += "c"
		// } else if an == 13 {
		// 	s += "d"
		// } else if an == 14 {
		// 	s += "e"
		// } else {
		// 	s += "f"
		// }
		s += strconv.FormatInt(int64(an), 16)
	}
	return s
}

// This function takes an input slice of integers (inputArray), which represents the input message, and performs various operations to prepare the message for hashing.
func Encrypt(inputArray []int) []int {

	// Convert the all string into binary
	length := len(inputArray)
	messageBit := []int{}

	// convert the char to binary
	for i := 0; i < length; i++ {
		a := inputArray[i]
		anss := []int{}
		// Convert the character to binary representation
		for a != 0 {
			h := a % 2
			anss = append(anss, h)
			a = a / 2
		}

		// Append leading zeros until the binary representation becomes 8 bits
		tmp := len(anss)
		for j := 0; j < 8-tmp; j++ {
			anss = append(anss, 0)
		}
		reverse(anss)

		// Append the binary representation of the character to the message bits
		for j := 0; j < 8; j++ {
			messageBit = append(messageBit, anss[j])
		}
	}

	// Append a '1' to the message bits
	messageBit = append(messageBit, 1)

	// Append zeros until the length of message bits is a multiple of 512 and has a remainder of 448
	for len(messageBit)%512 != 448 {
		messageBit = append(messageBit, 0)
	}

	// Calculate the length of the original message in bits
	length = length * 8

	// Create a slice of size 64 to store the length of the message in binary representation
	ml := make([]int, 64)
	p := 63

	// Convert the length to binary representation
	for length != 0 {
		h := length % 2
		ml[p] = h
		length = length / 2
		p = p - 1
	}

	// Append the binary representation of the length to the message bits
	for j := 0; j < 64; j++ {
		messageBit = append(messageBit, ml[j])
	}

	// Return the prepared message bits
	return messageBit
}

// This function reverses the order of elements in the input slice arr.
func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// This is the main hash function. It takes a string input inp, converts it to an array of ASCII values (inpBytes), and performs the SHA-1 hashing algorithm on the input message.
func Hash(inp string) string {

	// converts the individual char of string to its ascii value
	message := []rune(inp)

	inpBytes := make([]int, len(message))

	for i, c := range message {
		inpBytes[i] = int(c)
	}

	messageBit := Encrypt(inpBytes)

	A := []int{0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1} // A=[67452301] <- this is in hexadecimal form
	B := []int{1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1} // B=[efcdab89]
	C := []int{1, 0, 0, 1, 1, 0, 0, 0, 1, 0, 1, 1, 1, 0, 1, 0, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0} // C=[98badcfe]
	D := []int{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0} // D=[10325476]
	E := []int{1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0} // E=[c3d2e1f0]

	size := len(messageBit)
	nn := 0
	for n := 0; n < size; n += 512 {
		var w [][]int

		for i := nn; i < nn+16; i++ {
			var temp []int
			for j := i * 32; j < (i+1)*32; j++ {
				temp = append(temp, messageBit[j])
			}
			w = append(w, temp)
		}
		for i := 16; i < 80; i++ {
			temp := xorOfFour(w[i-3], w[i-8], w[i-14], w[i-16])
			w = append(w, leftRotate(temp, 1))
		}

		a := make([]int, len(A))
		b := make([]int, len(B))
		c := make([]int, len(C))
		d := make([]int, len(D))
		e := make([]int, len(E))
		copy(a, A)
		copy(b, B)
		copy(c, C)
		copy(d, D)
		copy(e, E)
		zeros := make([]int, 32)
		for i := 0; i < 80; i++ {
			var F []int
			var k []int

			if i >= 0 && i < 20 {
				F = orOfThree(andOp(b, c), andOp(notOp(b), d), zeros)
				k = []int{0, 1, 0, 1, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1} // k=[5A827999] <- this is in hexadecimal form
			} else if i >= 20 && i < 40 {
				F = xorOfFour(b, c, d, zeros)
				k = []int{0, 1, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 1} // k=[6ED9EBA1]
			} else if i >= 40 && i < 60 {
				F = orOfThree(andOp(b, c), andOp(b, d), andOp(c, d))
				k = []int{1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 0} // k=[8F1BBCDC]
			} else {
				F = xorOfFour(b, c, d, zeros)
				k = []int{1, 1, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 0, 1, 1, 0} // k=[CA62C1D6]
			}

			temp := addFive(leftRotate(a, 5), F, e, k, w[i])
			copy(e, d)
			copy(d, c)
			copy(c, leftRotate(b, 30))
			copy(b, a)
			copy(a, temp)
		}

		A = addTwo(A, a)
		B = addTwo(B, b)
		C = addTwo(C, c)
		D = addTwo(D, d)
		E = addTwo(E, e)

		nn += 16
	}

	dig := getDigest(append(append(append(append(A, B...), C...), D...), E...))
	return dig
}

// The main function handles user input and calls the Hash function to compute the SHA-1 hash of the input message.
func main() {

	// taking input string.
	var input string
	fmt.Println("Enter Input: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	// if input is empty.
	if strings.TrimSpace(input) == "" {
		fmt.Println("Please Enter Input.")
		os.Exit(1)
	}

	// Output the digest
	fmt.Println("Digest (SHA-1):", Hash(input))
}
