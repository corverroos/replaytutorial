package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/j"
)

var hideFlag = flag.Bool("hide", false, "If enabled, tips are hidden")

func main() {
	if err := run(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	var n int
	if len(flag.Args()) > 0 {
		var err error
		n, err = strconv.Atoi(flag.Arg(0))
		if err != nil || len(flag.Args()) != 1 {
			return errors.New("apart from flags, only a single int argument supported: " + fmt.Sprint(flag.Args()))
		}
	}

	file := os.Getenv("GOFILE")
	if *hideFlag {
		return hide(pwd, file)
	} else {
		return show(pwd, file, n)
	}
}

func show(pwd string, file string, n int) error {
	if n == 0 {
		return nil
	}

	fmt.Printf("Showing %d solution(s) for %s/%s\n", n, pwd, file)

	m, err := loadJson(pwd)
	if err != nil {
		return err
	}

	f, err := os.Open(path.Join(pwd, file))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var (
		num    int
		output []string
	)
	for scanner.Scan() {
		num++
		line := strings.TrimSpace(scanner.Text())
		if n > 0 && strings.HasPrefix(line, "//showme:hidden") {
			tip := strings.TrimSpace(strings.TrimPrefix(line, "//showme:hidden"))
			content, ok := m[tip]
			if !ok {
				return errors.New("solution not found in showme.json", j.MKV{"tip": tip, "num": num})
			}

			output = append(output, strings.Replace(scanner.Text(), "hidden", "start", 1))
			output = append(output, strings.Split(string(content), "\n")...)
			output = append(output, strings.Replace(scanner.Text(), "hidden", "end", 1))
			n--
			continue
		}

		output = append(output, scanner.Text())
	}

	return os.WriteFile(path.Join(pwd, file), []byte(strings.Join(output, "\n")), 0644)
}

func hide(pwd string, file string) error {
	fmt.Printf("Hiding solutions for %s/%s\n", pwd, file)

	m, err := loadJson(pwd)
	if err != nil {
		return err
	}

	f, err := os.Open(path.Join(pwd, file))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var (
		content []string
		intip   string
		num     int
		output  []string
	)
	for scanner.Scan() {
		num++
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "//showme:start") {
			if intip != "" {
				return errors.New("start while in tip", j.MKV{"line": line, "intip": intip, "num": num})
			}
			intip = strings.TrimSpace(strings.TrimPrefix(line, "//showme:start"))
			output = append(output, strings.Replace(scanner.Text(), "start", "hidden", 1))
		} else if strings.HasPrefix(line, "//showme:end") {
			if intip == "" {
				return errors.New("end while in not tip", j.MKV{"line": line, "num": num})
			}
			m[intip] = []byte(strings.Join(content, "\n"))
			content = nil
			intip = ""
		} else if intip != "" {
			content = append(content, scanner.Text())
		} else {
			output = append(output, scanner.Text())
		}
	}

	if intip != "" {
		return errors.New("EOF while in tip", j.MKV{"intip": intip, "num": num})
	}

	err = storeJson(pwd, m)
	if err != nil {
		return err
	}

	src := []byte(strings.Join(output, "\n"))

	src, err = format.Source(src)
	if err != nil {
		return err
	}

	return os.WriteFile(path.Join(pwd, file), src, 0644)
}

func loadJson(pwd string) (map[string][]byte, error) {
	res := make(map[string][]byte)

	b, err := os.ReadFile(path.Join(pwd, "showme.json"))
	if os.IsNotExist(err) {
		return res, nil
	} else if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func storeJson(pwd string, m map[string][]byte) error {
	b, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(path.Join(pwd, "showme.json"), b, 0644)
}
