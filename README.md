# Hash Encryption

Hash Encryption is a Go package that implements a secure hash algorithm for encrypting data and generating a message digest. The package provides a simple and efficient way to secure sensitive information and ensure data integrity.

## Features

- Generates a 160-bit message digest for input data.
- Implements a secure hash algorithm for encryption.
- Supports encryption of various data types.

![SHA-1_state_model](https://user-images.githubusercontent.com/62175306/94359128-d14be700-00c2-11eb-9e9d-8dc78d1fb41a.png)

## ## Installation

To use Hash Encryption in your Go project, you need to have Go programming language (version X.X.X or higher) installed. Then, follow these steps:

1. Install the package using the `go get` command:
```
go get github.com/Sahil-4555/Secure_Hash_Algorithm-SHA1
```
2. Import the package in your Go code:
```
import "github.com/Sahil-4555/Secure_Hash_Algorithm-SHA1"
```
## Usage
To encrypt data and generate a message digest, follow these steps:

1. Invoke the Encrypt function, passing your input data as an argument:
```
encryptedData := hashencryption.Encrypt("input data")
```
2. Use the encryptedData as desired in your application.
Here's an example of encrypting a string using the Hash Encryption package:
```
package main

import (
	"fmt"
	"github.com/Sahil-4555/Secure_Hash_Algorithm-SHA1"
)

func main() {
	encryptedData := hashencryption.Encrypt("Hello, World!")
	fmt.Println("Encrypted Data:", encryptedData)
}
```
This will output the encrypted message digest for the input data.
