package leetcode

import (
	"container/list"
)

func isValid(s string) bool {
	signMap := make(map[int32]int32)
	signMap[')'] = '('
	signMap[']'] = '['
	signMap['}'] = '{'

	bs := []byte(s)
	signLst := new(list.List)
	for i := 0; i < len(bs); i++ {
		switch {
		case bs[i] == '(':
			signLst.PushBack('(')
			break
		case bs[i] == '[':
			signLst.PushBack('[')
			break
		case bs[i] == '{':
			signLst.PushBack('{')
			break
		case bs[i] == ')' || bs[i] == ']' || bs[i] == '}':
			if signLst.Len() <= 0 {
				return false
			}
			sign := signLst.Back()
			if sign.Value.(int32) != signMap[int32(bs[i])] {
				return false
			} else {
				signLst.Remove(sign)
			}
			break
		default:
			break
		}
	}

	if signLst.Len() == 0 {
		return true
	} else {
		return false
	}
}
