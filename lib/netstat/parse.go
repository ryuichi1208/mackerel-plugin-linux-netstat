package netstat

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type parser func(io.Reader) map[string]float64

type source struct {
	path   string
	parser parser
}

func (s *source) read() (map[string]float64, error) {
	values := make(map[string]float64)
	file, err := os.Open(s.path)
	if err != nil {
		return values, err
	}

	values = s.parser(file)
	return values, nil
}

var sources = [...]source{
	{path: "/proc/net/netstat", parser: parseCompact},
	{path: "/proc/net/snmp", parser: parseCompact},
	{path: "/proc/net/sockstat", parser: parseSocket},
}

func splitKeyValues(s, kvSep string) (string, []string, error) {
	var key string
	var vals []string
	if kv := strings.SplitN(s, kvSep, 2); len(kv) == 2 {
		key = kv[0]
		vals = strings.Fields(kv[1])
	} else {
		err := fmt.Errorf("No key-value separator found in input string.")
		if err != nil {
			return key, vals, err
		}
	}
	return key, vals, nil
}

func parseSocket(file io.Reader) map[string]float64 {
	result := make(map[string]float64)
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		txt := scanner.Text()
		if strings.Contains(txt, "TCP") {
			arr := strings.Split(txt, " ")
			v, err := strconv.ParseFloat(arr[len(arr)-1], 64)
			if err != nil {
				panic(err)
			}
			result["tcp_memory_size"] = v * float64(os.Getpagesize())
		}
	}
	return result
}

func parseCompact(file io.Reader) map[string]float64 {
	result := make(map[string]float64)

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		key, names, err := splitKeyValues(scanner.Text(), ": ")

		if err != nil || !scanner.Scan() {
			break
		}

		key2, values, err := splitKeyValues(scanner.Text(), ": ")

		if err != nil || key != key2 || len(names) != len(values) {
			break
		}

		for i := 0; i < len(names); i++ {
			result[key+names[i]], err = strconv.ParseFloat(values[i], 64)
			if err != nil {
				panic(err)
			}
		}
	}

	return result
}

func merge(m ...map[string]float64) map[string]float64 {
	ans := make(map[string]float64, 0)

	for _, c := range m {
		for k, v := range c {
			ans[k] = v
		}
	}
	return ans
}

func get() map[string]float64 {
	r := make(map[string]float64)
	for _, source := range sources {
		stats, err := source.read()
		if err != nil {
			log.Printf("Error reading stats file: %v", err)
			continue
		}
		for k, v := range stats {
			r[k] = v
		}
	}
	return r
}
