package util
//栈数据结构
type Stack struct{
	data []interface{}//元素为空接口型的切片
}

func NewStackBySlice(sli []interface{}) *Stack{
	return &Stack{data:sli}
}

func NewStack() *Stack{ //初始化栈
	return &Stack{data: []interface{}{}}
}
func (this *Stack) ToSlice() []interface{}{
	res:=make([]interface{},len(this.data))
	copy(res,this.data)
	return res
}
func (this *Stack) Peek() interface{}{
	return this.data[len(this.data)-1]
}
func (this *Stack) PeekFloat64() float64{
	return this.Peek().(float64)
}
func (this *Stack) PeekChar() byte{
	return this.Peek().(byte)
}
func (this *Stack) Push(in interface{}){
	this.data=append(this.data,in)
}
func (this *Stack) Pop() interface{}{
	res:=this.Peek()
	this.data=this.data[:len(this.data)-1]
	return res
}
func (this *Stack) PopFloat64() float64{
	return this.Pop().(float64)
}

func (this *Stack) PopChar() byte{
	return this.Pop().(byte)
}
func (this *Stack) IsEmpty() bool{
	return len(this.data)==0
}