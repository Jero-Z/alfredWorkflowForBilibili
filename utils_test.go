package main

import (
	"fmt"
	"testing"
)

type args struct {
	args []string
}

type test struct {
	name          string
	args          args
	wantQueryType string
	wantKeyWord   string
	wantSort      string
}

func Test_parseArgs(t *testing.T) {
	a := make([]string, 0)

	var tests []struct {
		name          string
		args          args
		wantQueryType string
		wantKeyWord   string
		wantSort      string
	}

	a = append(a, "123")
	a = append(a, "t=v")
	a = append(a, "keyword")
	a = append(a, "s=p")
	fmt.Println(a)
	//args{
	//	args: a,
	//}
	//
	//test{
	//	name: "t1",
	//	args: args{
	//		args: args,
	//	},
	//}
	//tests = append(tests,{
	//	name:"zhangsan"
	//	args:""
	//	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQueryType, gotKeyWord, gotSort := parseArgs(tt.args.args)
			if gotQueryType != tt.wantQueryType {
				t.Errorf("parseArgs() gotQueryType = %v, want %v", gotQueryType, tt.wantQueryType)
			}
			if gotKeyWord != tt.wantKeyWord {
				t.Errorf("parseArgs() gotKeyWord = %v, want %v", gotKeyWord, tt.wantKeyWord)
			}
			if gotSort != tt.wantSort {
				t.Errorf("parseArgs() gotSort = %v, want %v", gotSort, tt.wantSort)
			}
		})
	}
}
