# Vigenere Cipher
This is a simple command line implementation of a Vigenere cipher in Go.

## Flags
- `-i` string: the path to the input file (text to encode or decode)
- `-o` string: the path to the desired output file (where to store encoded or decoded message)
- `-k` string: the keyword used to encrypt the message
- `-e` boolean: flag for encoding
- `-d` boolean: flag for decoding

Encoding is the default, with a default keyword "a", which keeps all letters the same. The default input file is cleartext.txt and the default output file is ciphertext.txt.

## Example
```
vigenere -i cleartext.txt -k forget -e -o ciphertext.txt
vigenere -i ciphertext.txt -k forget -d -o message.txt
```
