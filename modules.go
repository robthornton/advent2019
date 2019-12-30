package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func modulesFromReader(r io.Reader) ([]int64, error) {
	modules := make([]int64, 0)

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return modules, err
	}

	buf = bytes.TrimRight(buf, "\n")
	masses := strings.Split(string(buf), "\n")
	for _, mass := range masses {
		n, err := strconv.ParseInt(mass, 10, 64)
		if err != nil {
			return modules, err
		}
		modules = append(modules, n)
	}

	return modules, nil
}
