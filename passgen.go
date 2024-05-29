package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/atotto/clipboard"
)

var passwords = make(map[string]string)

const dataFile = "passwords.gob"

func main() {
	loadData()
	removePtr := flag.Bool("r", false, "boolean for removal")
	flag.Parse()

	//removing
	if *removePtr {
		fmt.Println("removing")
		args := os.Args[2:]
		key := args[0]
		delete(passwords, key)
		saveData()
		return
	}
	//no flag, generate password & store
	args := os.Args[1:]
	key := args[0]

	if password, ok := passwords[key]; ok {
		fmt.Println("Password for " + key + ": " + password)

		err := clipboard.WriteAll(password)
		if err != nil {
			fmt.Println("Error copying password to clipboard:", err)
		} else {
			fmt.Println("Password copied to clipboard.")
		}
	} else {
		newPass := generateString()
		passwords[key] = newPass
		fmt.Println("New Password: " + passwords[key])
		err := clipboard.WriteAll(newPass)
		if err != nil {
			fmt.Println("Error copying password to clipboard:", err)
		} else {
			fmt.Println("Password copied to clipboard.")
		}
		saveData()
	}
}

func loadData() {
	file, err := os.Open(dataFile)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(&passwords)
		if err != nil {
			fmt.Println("Error decoding data:", err)
		}
		file.Close()
	}
}

func saveData() {
	file, err := os.Create(dataFile)
	if err == nil {
		encoder := gob.NewEncoder(file)
		err = encoder.Encode(passwords)
		if err != nil {
			fmt.Println("Error encoding data:", err)
		}
		file.Close()
	}
}

//Password must be 12 character long
//must include upper, lower, special character, number
func generateString() string {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	upper := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	lower := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	nums := [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	special := [10]string{"!", "@", "#", "4", "%", "^", "&", "*", "(", ")"}

	password := ""

	for i := 0; i < 5; i++ {
		password = password + upper[r1.Intn(25)]
		password = password + lower[r1.Intn(25)]
		password = password + nums[r1.Intn(9)]
		password = password + special[r1.Intn(9)]
	}

	return password
}
