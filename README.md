# Hash Encryption

Hash Encryption is a Go package that implements a secure hash algorithm for encrypting data and generating a message digest. The package provides a simple and efficient way to secure sensitive information and ensure data integrity.

## Features

- Generates a 160-bit message digest for input data.
- Generates 40 digits hexadecimal code of password
- Implements a secure hash algorithm for encryption.
- Supports encryption of various data types.

![SHA-1_state_model](https://user-images.githubusercontent.com/62175306/94359128-d14be700-00c2-11eb-9e9d-8dc78d1fb41a.png)

## Functionality
Hash Encryption provides the following functionality:

 - Encryption of messages, passwords, or any other sensitive data
 - Support for popular hash encryption algorithms such as MD5, SHA-1, SHA-256, etc.
 - Customizable options for adjusting encryption parameters
 - Protection against unauthorized access and data tampering
 
## Go Implementation

Run The `main.go` file Using Command:
```
go run main.go
````
### Output
![SHA1_GO_IMPLEMENTATION](https://github.com/Sahil-4555/Secure_Hash_Algorithm-SHA1/blob/master/assets/sha1-go.png)

## Installation

To use Hash Encryption in your Go project, you need to have Go programming language installed. Then, follow these steps:

 - Install the package using the `go get` command:
```
go get github.com/Sahil-4555/Secure_Hash_Algorithm-SHA1
```
 - Import the package in your Go code:
```
import "github.com/Sahil-4555/Secure_Hash_Algorithm-SHA1"
```
## Usage
To encrypt data and generate a message digest, follow these steps:

 - Invoke the Encrypt function, passing your input data as an argument:
```
encryptedData := hashencryption.Encrypt("input data")
```
 - Use the encryptedData as desired in your application.
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
