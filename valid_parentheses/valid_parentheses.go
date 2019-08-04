package valid_parentheses

//
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。
//
//示例 1:
//
//输入: "()"
//输出: true
//示例 2:
//
//输入: "()[]{}"
//输出: true
//示例 3:
//
//输入: "(]"
//输出: false
//示例 4:
//
//输入: "([)]"
//输出: false
//示例 5:
//
//输入: "{[]}"
//输出: true

type Stack struct {
	Top *Node
}

type Node struct {
	Value interface{}
	Next  *Node
}

func (s *Stack) Push(value interface{}) {
	if s == nil {
		return
	}
	n := Node{Value: value, Next: s.Top}
	s.Top = &n
}

func (s *Stack) Pop() (value interface{}) {
	if s == nil || s.Top == nil {
		return nil
	}
	value = s.Top.Value
	s.Top = s.Top.Next
	return
}

func (s *Stack) IsEmpty() (b bool) {
	b = s.Top == nil
	return
}

func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}

	stack := Stack{}
	for _, c := range s {
		if c == 40 || c == 91 || c == 123 {
			stack.Push(c)
		} else {
			p := stack.Pop()
			if p == nil {
				return false
			}
			val := p.(int32)
			if c == 41 && val != 40 {
				return false
			}
			if c == 93 && val != 91 {
				return false
			}
			if c == 125 && val != 123 {
				return false
			}
		}
	}
	return stack.IsEmpty()
}
