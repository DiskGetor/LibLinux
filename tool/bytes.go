package tool

import (
	"bytes"
	"io/ioutil"
	"sort"
)

func (p *ObjTool) JoinBytes(Bytes ...[]byte) []byte {
	Size := len(Bytes)
	BaseAddr := make([][]byte, Size)
	for i := 0; i < Size; i++ {
		BaseAddr[i] = Bytes[i]
	}
	return bytes.Join(BaseAddr, []byte(``)) //?0x00合适？测试下
}

func (p *ObjTool) CompareBufFromFile(buf []byte, fileName string) {
	W.LogInfo("Compare Buffer from ", fileName)
	W.LogHexDump("input", buf)
	want, err := ioutil.ReadFile(fileName)
	if !W.CheckErr(err) {
		return
	}
	W.LogHexDump("want", want)
	if bytes.Equal(buf, want) {
		W.LogSuccess("pass")
		println()
		return
	}
	W.LogError("not pass")
	println()
}

//数组去重
func (p *ObjTool) RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	sort.Strings(arr)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func (p *ObjTool) RemoveRepeatedElementV2(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
