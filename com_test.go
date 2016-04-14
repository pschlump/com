// (C) Philip Schlump, 2014.

package com

import (
	"testing"
)

//type TestCase struct {
//	format string
//	expected string
//	value float64
//}
//
//var testCases = []TestCase{
//	{ "##,##0.00",     "  -123.46",       -123.456 },		// round to .47???
//	{ "##,##0.00",     "   123.46",        123.456 },		// 1
//	{ "##,##0.00",     "##,##0.00",     123123.456 },		// 2
//
//	// Escape
//	// &TestCase{"YYYY/MM/DD hh:mm:ss", "2014/01/10 11:31:32"},
//	// &TestCase{"YYYY-MM-DD hh:mm:ss", "2014-01-10 11:31:32"},
//	// In a string
//	// &TestCase{"/aaaa/YYYY/mm/bbbb", "/aaaa/2014/31/bbbb"},
//
//	// No Format - get rid of value
//	{ "",              "",                 123.456 },		// 3
//	{ ".",             ".",                123.456 },		// 4
//	{ "0",             "3",                  3.456 },		// 5
//	{ "#",             "3",                  3.456 },		// 6
//
//	{ "##,##0.00",     "     0.00",          0.0   },		// 7
//	{ "##,##0.00",     "     0.00",         -0.0   },		// 1
//}

func TestPathToRelativeInverse(t *testing.T) {
	//if false {
	//	fmt.Printf ( "keep compiler happy when we are not using fmt.\n" )
	//}
	rv := PathToRelativeInverse("./x")
	ex := ".."
	if rv != ex {
		t.Fatalf("Error results=[%s] expected=[%s]", rv, ex)
	}
	//for i, v := range testCases {
	//	// fmt.Printf ( "Running %d\n", i )
	//	result := Format(v.format, v.value)
	//	if result != v.expected {
	//		t.Fatalf("Error for %f at [%d] in table: format=[%s]: results=[%s] expected=[%s]", v.value, i, v.format, result, v.expected)
	//	}
	//}
}

/* vim: set noai ts=4 sw=4: */
