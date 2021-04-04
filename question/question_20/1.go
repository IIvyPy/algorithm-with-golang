package question_20

var m = map[string]string{
	")": "(",
	"}": "{",
	"]": "[",
}

type stack interface{
	Pop() interface{}
	Push(interface{})
	Empty() bool
}

type stackList struct{
	list []interface{}
}

func (s *stackList) Pop() interface{}{
	if s.IsEmpty() {
		return nil
	}
	last := s.list[len(s.list)-1]
	s.list = s.list[0:len(s.list)-1]
	return last
}

func (s *stackList) Push(val interface{}) {
	s.list = append(s.list, val)
}

func (s *stackList) Peek() interface{}{
	if s.IsEmpty(){
		return nil
	}
	return s.list[len(s.list)-1]
}

func (s *stackList) IsEmpty() bool{
	if len(s.list) == 0{
		return true
	}
	return false
}

func NewStack() *stackList{
	s := new(stackList)
	s.list = make([]interface{}, 0)
	return s
}

func isValid(s string) bool {
	stack := NewStack()
	for i := 0; i < len(s); i++{
		v, isOk := m[string(s[i])]
		if isOk{
			value := stack.Peek()
			l, ok := value.(string)
			if !ok{
				return false
			}
			if l != v{
				return false
			}
			stack.Pop()
		}else{
			stack.Push(string(s[i]))
		}
	}
	return stack.IsEmpty()
}