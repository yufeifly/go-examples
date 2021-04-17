package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func DeleteExtraSpace(s string) string {
	//删除字符串中的多余空格，有多个空格时，仅保留一个空格
	s1 := strings.Replace(s, "  ", " ", -1)      //替换tab为空格
	regstr := "\\s{2,}"                          //两个及两个以上空格的正则表达式
	reg, _ := regexp.Compile(regstr)             //编译正则表达式
	s2 := make([]byte, len(s1))                  //定义字符数组切片
	copy(s2, s1)                                 //将字符串复制到切片
	spc_index := reg.FindStringIndex(string(s2)) //在字符串中搜索
	for len(spc_index) > 0 {                     //找到适配项
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...) //删除多余空格
		spc_index = reg.FindStringIndex(string(s2))            //继续在字符串中搜索
	}
	return string(s2)
}

func adjustNameForGBT(Names string) string {
	names := strings.Split(Names, ",")
	var newNames []string
	for i, p := range names {
		names[i] = strings.TrimSpace(p)
		names[i] = DeleteExtraSpace(names[i])
		pieces := strings.Split(names[i], " ")
		s := 0
		t := len(pieces) - 1
		for s < t {
			pieces[s], pieces[t] = pieces[t], pieces[s]
			s++
			t--
		}
		for ind := range pieces[:len(pieces)-1] {
			pieces[ind] += "."
		}
		var newName string
		for _, piece := range pieces[:len(pieces)-1] {
			newName += piece + " "
		}
		newName += pieces[len(pieces)-1]
		newNames = append(newNames, newName)
	}
	var outputNames string
	for _, newN := range newNames[:len(newNames)-1] {
		outputNames += newN + ", "
	}
	if newNames != nil {
		outputNames += newNames[len(newNames)-1]
	}
	return outputNames
}

func deleteAnd(raw string) string {
	ind := strings.Index(raw, "&")
	if ind != -1 {
		return raw[:ind] + raw[ind+1:]
	}
	return raw
}

func xprint(names []string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func dealName(name string) string {
	var output string
	for _, ch := range name {
		if ch == ' ' {
			continue
		}
		if ch == '.' {
			ch = ','
		}
		output += string(ch)
	}
	return output
}

func adjustNameForAPA(Names string) string {
	Names = deleteAnd(Names)
	names := strings.Split(Names, ".,")
	var newNames []string
	for i, p := range names {
		names[i] = strings.TrimSpace(p)
		names[i] = DeleteExtraSpace(names[i])
		names[i] = dealName(names[i])
		pieces := strings.Split(names[i], ",")
		//xprint(pieces)
		s := 0
		t := len(pieces) - 1
		for s < t {
			pieces[s], pieces[t] = pieces[t], pieces[s]
			s++
			t--
		}
		for ind := range pieces[:len(pieces)-1] {
			pieces[ind] += "."
			//fmt.Println("f: ",pieces[ind])
		}
		var newName string
		for _, piece := range pieces[:len(pieces)-1] {
			newName += piece + " "
		}
		newName += pieces[len(pieces)-1]
		newNames = append(newNames, newName)
	}
	//xprint(newNames)
	var outputNames string
	for _, newN := range newNames[:len(newNames)-1] {
		outputNames += newN + ", "
	}
	if newNames != nil {
		outputNames += newNames[len(newNames)-1]
	}
	return outputNames
}

func main() {
	fi, err := os.Open("names.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(adjustNameForGBT(string(a)))
	}
}
