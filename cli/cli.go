package cli

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/hachi-n/passwd_gen/lib/config"
	"log"
	"os"
)

var Option *Flags

type Flags struct {
	Server       bool
	Port         int
	ServiceName  string
	CategoryName string
	Length       int
	NumberLength int
	SymbolLength int
	List         bool
}

func Load() {
	f := new(Flags)
	flag.StringVar(&f.ServiceName, "s", "", "service name")
	flag.StringVar(&f.CategoryName, "c", config.Config.CategoryName, "category name")
	flag.BoolVar(&f.Server, "S", false, "server mode default port 8888")
	flag.IntVar(&f.Port, "P", config.Config.Port, "server mode default port 8888")
	flag.IntVar(&f.Length, "l", config.Config.Length, "password length")
	flag.IntVar(&f.NumberLength, "n", config.Config.NumberLength, "password numberType length")
	flag.IntVar(&f.SymbolLength, "m", config.Config.SymbolLength, "password symbolType length")
	flag.BoolVar(&f.List, "L", false, "saved password list")

	flag.Parse()

	if !f.List {
		paramsValidate(f)
	}

	Option = f
}

func paramsValidate(f *Flags) {
	validPasswordLength := 10
	if f.Length <= validPasswordLength {
		log.Fatalf("password is too short. password is more than %s charactors.\n", validPasswordLength)
	}

	if f.Length < f.NumberLength+f.SymbolLength {
		log.Fatalln("password length is more than (number length + symbol length).")
	}

	if !f.Server && f.ServiceName == "" {
		f.ServiceName = interactive("サービス名")

	}

	if !f.Server && f.CategoryName == "" {
		f.CategoryName = interactive("カテゴリ名")
	}
}

func interactive(message string) string {
	scanner := bufio.NewScanner(os.Stdin)
	var result string
	fmt.Printf("%sを入力してください。 : ", message)
	var failureCount int
	for scanner.Scan() {
		result = scanner.Text()
		if result == "" {
			fmt.Printf("%sを入力してください。 : ", message)
			failureCount++
		} else {
			break
		}

		if failureCount >= 5 {
			fmt.Println()
			log.Fatalf("%sは必須項目です。", message)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}
