package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := flag.String("path", "", "antd path， eg: node_modules/antd")
	flag.Parse()
	*root = strings.TrimLeft(*root, "./")

	walkANTD(*root)
	fmt.Println("EOF")
}

var (
	variablePaths  []string
	componentPaths []string

	variablePath  string
	componentPath string

	basePath string
)

func walkANTD(path string) {
	basePath = path
	// basic
	variablePaths = []string{basePath + "/lib/style/themes", basePath + "/lib/style/color"}
	for _, v := range variablePaths {
		variablePath = v
		err := filepath.Walk(variablePath, walkfuncVariable)
		if err != nil {
			panic(err)
		}
	}

	// components
	componentPaths = []string{basePath + "/lib"}
	for _, v := range componentPaths {
		componentPath = v
		err := filepath.Walk(componentPath, walkfuncComponent)
		if err != nil {
			panic(err)
		}
	}

	writeLessFile()
}

func walkfuncVariable(path string, info os.FileInfo, err error) error {
	if info != nil && !info.IsDir() && strings.HasSuffix(info.Name(), "less") {
		err := turnLess(path, "variable")
		if err != nil {
			return err
		}
	}

	return nil
}

func walkfuncComponent(path string, info os.FileInfo, err error) error {
	if info != nil && !info.IsDir() && strings.HasSuffix(info.Name(), "less") {
		ignore := false
		for _, p := range variablePaths {
			if strings.Index(path, p) == 0 {
				ignore = true
			}
		}
		if !ignore {
			for _, v := range variablePaths {
				if strings.HasPrefix(path, v) {
					break
				}
			}

			err := turnLess(path, "component")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

var (
	importLine         string
	antdComponentsLine string
	antdVariablesLine  string
)

func turnLess(path string, lessType string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineText := scanner.Text()
		if strings.HasPrefix(lineText, "@import") {
			importLine += (lineText + "\n")
		} else {
			if lessType == "variable" {
				antdVariablesLine += lineText + "\n"
			} else if lessType == "component" {
				antdComponentsLine += lineText + "\n"
			}
		}
	}

	err = scanner.Err()
	if err != nil {
		return err
	}

	return nil
}

func writeLessFile() error {
	if len(antdVariablesLine) == 0 || len(antdComponentsLine) == 0 {
		fmt.Println("Please confirm antd path.")
		return errors.New("nil data")
	}
	err := ioutil.WriteFile("antd-components.less", []byte(antdComponentsLine), 0777)
	if err != nil {
		return err
	}
	fmt.Println("Write antd-components ok.")
	err = ioutil.WriteFile("antd-variables.less", []byte(antdVariablesLine), 0777)
	if err != nil {
		return err
	}
	fmt.Println("Write antd-variables ok.")
	err = ioutil.WriteFile("all.less", []byte(antdVariablesLine+antdComponentsLine), 0777)
	if err != nil {
		return err
	}
	fmt.Println("Write all ok.")
	return nil
}
