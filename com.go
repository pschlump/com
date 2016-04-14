package com

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pschlump/json" //	"encoding/json"
	"github.com/pschlump/ms"
)

const (
	PathSep = string(os.PathSeparator)
)

func ReadInGlobalConfig(path string) map[string]string {
	var jsonData map[string]string

	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error(12006): Reading in global config %v\n", err)
		return jsonData
	}
	err = json.Unmarshal(file, &jsonData)
	if err != nil {
		fmt.Printf("Error(12007): Reading in global config %v\n", err)
		jsonData["error"] = "true"
		return jsonData
	}

	if s, ok := jsonData["static_dir"]; !ok || s == "" {
		jsonData["static_dir"] = "./static"
	}

	if s, ok := jsonData["listen_at_ip_port"]; ok || s != "" {
		port := "80"
		addr := "127.0.0.1"
		tmp := strings.Split(s, ":")
		if len(tmp) == 2 {
			addr = tmp[0]
			port = tmp[1]
		} else if len(tmp) > 0 {
			addr = tmp[0]
		}
		jsonData["ip_addr"] = addr
		jsonData["ip_port"] = addr + ":" + port
		jsonData["port"] = ":" + port
		// fmt.Printf("addr ->%s<- port ->%s<- \n", addr, port)
		//,"ip_addr":"192.168.0.102"
		//,"ip_port":"192.168.0.102:8080"
		//,"port":":8080"
	}

	return jsonData
}

func ReadInGlobalConfigRaw(path string) map[string]string {
	var jsonData map[string]string

	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error(12006): Reading in global config %v\n", err)
		return jsonData
	}
	err = json.Unmarshal(file, &jsonData)
	if err != nil {
		fmt.Printf("Error(12007): Reading in global config %v\n", err)
		jsonData["error"] = "true"
		return jsonData
	}

	if s, ok := jsonData["static_dir"]; !ok || s == "" {
		jsonData["static_dir"] = "./static"
	}

	return jsonData
}

type FtpUser struct {
	Username   string
	Password   string
	Server     string
	Port       int
	DefaultCwd string
}

var FTPConfig FtpUser

func ReadFTPConfig(fn string) (rv FtpUser) {
	s := GetFile(fn)
	err := json.Unmarshal([]byte(s), &rv)
	if err != nil {
		fmt.Printf("Error(12029): Invalid format for FTP config file - %v\n", err)
	}
	return
}

func AddHomeDir(fn string) string {
	// PathSep = string(os.PathSeparator)
	if fn[0:2] == "~"+PathSep {
		return ms.HomeDir() + PathSep + fn[2:]
	} else {
		return fn
	}
}

func SumVector(w []float64) (r float64) {
	r = 0.0
	for _, v := range w {
		r += v
	}
	return
}

// Size		Width x Height (mm)		Width x Height (in)
// 4A0		1682 x 2378 mm			66.2 x 93.6 in
// 2A0		1189 x 1682 mm			46.8 x 66.2 in
// A0		841 x 1189 mm			33.1 x 46.8 in
// A1		594 x 841 mm			23.4 x 33.1 in
// A2		420 x 594 mm			16.5 x 23.4 in
// A3		297 x 420 mm			11.7 x 16.5 in
// A4		210 x 297 mm			8.3 x 11.7 in
// A5		148 x 210 mm			5.8 x 8.3 in
// A6		105 x 148 mm			4.1 x 5.8 in
// A7		74 x 105 mm				2.9 x 4.1 in
// A8		52 x 74 mm				2.0 x 2.9 in
// A9		37 x 52 mm				1.5 x 2.0 in
// A10		26 x 37 mm				1.0 x 1.5 in

func GetPaperWidth(paper string) float64 {
	widthFromPaper := map[string]float64{
		"4A0": 1682,
		"2A0": 1189,
		"A0":  841,
		"A1":  594,
		"A2":  420,
		"A3":  297,
		"A4":  210,
		"A5":  148,
		"A6":  105,
		"A7":  74,
		"A8":  52,
		"A9":  37,
		"A10": 26,
	}
	if w, ok := widthFromPaper[paper]; ok {
		return w
	} else {
		return 210
	}
}

func FindColNo(ColName string, data0 map[string]interface{}) int {
	ii := 0
	for nm := range data0 {
		if nm == ColName {
			return ii
		}
		ii++
	}
	return -1
}

func ReadData(fn string) (data []map[string]interface{}) {
	s := GetFile(fn)
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		// fmt.Fprintf ( xOut, "Data ->%s<-\n", s )
		fmt.Fprintf(os.Stdout, "Error(12001): Invalid format - %v\n", err)
	}
	return
}

func PathToRelativeInverse(path string) (inv string) {
	if path[0:1] == "/" {
		inv = filepath.Clean(path)
	} else {
		inv = "./" + filepath.Clean(path)
	}
	re1 := regexp.MustCompile("/[^/]*")
	inv = re1.ReplaceAllLiteralString(inv, "/..")
	inv = filepath.Clean(inv)
	return
}

// package main

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func CopyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	//if err = os.Link(src, dst); err == nil {
	//    return
	//}
	err = copyFileContents(src, dst)
	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

// -------------------------------------------------------------------------------------------------
func GetFile(fn string) string {
	file, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Printf("Error(10103): File (%s) missing or unreadable error: %v\n", fn, err)
		return ""
	}
	return string(file)
}
