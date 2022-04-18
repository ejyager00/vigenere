package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var infile string
	var outfile string
	var keyword string
	var to_encrypt bool
	var to_decrypt bool
	flag.StringVar(&infile, "i", "cleartext.txt", "Specify input file. Default is cleartext.txt.")
	flag.StringVar(&outfile, "o", "ciphertext.txt", "Specify output file. Default is ciphertext.txt.")
	flag.StringVar(&keyword, "k", "a", "Specify keyword for encryption. Default is 'a'.")
	flag.BoolVar(&to_encrypt, "e", false, "Specify encryption. Default is true.")
	flag.BoolVar(&to_decrypt, "d", false, "Specify decryption. Default is false.")

	flag.Parse()
	if !to_encrypt && !to_decrypt {
		to_encrypt = true
	}

	if to_encrypt {
		cleartext, err := os.ReadFile(infile)
		fileReadError(err, infile)
		err = os.WriteFile(outfile, []byte(Encrypt(string(cleartext), keyword)), 0644)
		fileReadError(err, outfile)
	} else {
		ciphertext, err := os.ReadFile(infile)
		fileReadError(err, infile)
		err = os.WriteFile(outfile, []byte(Decrypt(string(ciphertext), keyword)), 0644)
		fileReadError(err, outfile)
	}
}

func fileReadError(e error, file string) {
	if e != nil {
		fmt.Println(e)
		fmt.Printf("The file %s could not be used.\n", file)
		os.Exit(1)
	}
}

func Encrypt(cleartext string, keyword string) string {
	text := allCaps(removePunctuation(cleartext))
	b := []byte(text)
	k := []byte(allCaps(keyword))
	for i, x := range b {
		b[i] = ((x-65)-(k[i%len(k)]-65)+26)%26 + 65
	}
	return string(b)
}

func Decrypt(ciphertext string, keyword string) string {
	b := []byte(ciphertext)
	k := []byte(allCaps(keyword))
	for i, x := range b {
		b[i] = ((x-65)+(k[i%len(k)]-65))%26 + 65
	}
	return string(b)
}

func removePunctuation(text string) string {
	b := []byte(text)
	out := make([]byte, 0)
	for _, x := range b {
		if (x >= 97 && x <= 122) || (x >= 65 && x <= 90) {
			out = append(out, x)
		}
	}
	return string(out)
}

func allCaps(text string) string {
	b := []byte(text)
	out := make([]byte, len(text))
	for i, x := range b {
		if x >= 97 && x <= 122 {
			out[i] = x - 32
		} else {
			out[i] = x
		}
	}
	return string(out)
}
