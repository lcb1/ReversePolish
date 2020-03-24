package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"util"
)

var opm=map[byte]byte{ //运算符的优先级映射表
	'(':0,
	')':0,
	'+':1,
	'-':1,
	'*':2,
	'/':2,
}

func testPaseEp(){
	ep:="1+((2+3)*4)-5"
	ins := parseEp(ep)
	for i:=0;i<len(ins);i++{
		if v,ok:=ins[i].(float64);ok{
			fmt.Print(v)
		}else {
			v:=ins[i].(byte)
			fmt.Print(string(v))
		}
	}
}

func testStack(){
	s:=util.NewStack()
	ep:="((1 2 3 .1 + 1))+((1+2)*4)"
	ins := parseEp(ep)
	for i:=0;i<len(ins);i++{
		s.Push(ins[i])
	}
	for !s.IsEmpty(){

		if _ ,ok:= s.Peek().(byte);ok{
			fmt.Println(string(s.PopChar()))
		}else {
			fmt.Println(s.PopFloat64())
		}
	}
}

func testConvert(){
	ep:="1+((2+3)*4)-5"
	mid:=parseEp(ep)
	fmt.Println(mid)
	suffix:=convert(mid)
	for i:=0;i<len(suffix);i++{
		if b,ok:=suffix[i].(byte);ok{
			fmt.Println(string(b))
		}else {
			fmt.Println(suffix[i])
		}
	}

}
func testCalcSuffix(){
	//ep:="1+((2+3)*4)-5"
	ep:="(1*2)/2+1-1"
	mid:=parseEp(ep)
	suffix:=convert(mid)
	//fmt.Println(suffix)
	res:=calcSuffix(suffix)
	fmt.Println(res)
}

func main() {
	//testPaseEp()
	//testConvert()
	//testCalcSuffix()

	fmt.Println(calcSuffix(convert(parseEp(getAllArgs()))))
}

func getAllArgs() string{   //获取命令行所有参数
	res:=""
	for i:=1;i<len(os.Args);i++{
		res+=os.Args[i]
	}
	return res
}


func calcSuffix(suffix []interface{}) float64{  //根据逆波兰表达式计算结果
	s:=util.NewStack()
	for i:=0;i<len(suffix);i++{
		if op,ok:=suffix[i].(byte);ok{
			num1:=s.PopFloat64()
			num2:=s.PopFloat64()
			res:=calc(num2,num1,op)
			s.Push(res)
		}else {
			s.Push(suffix[i])
		}

	}
	return s.PopFloat64()
}


func calc(num1,num2 float64,op byte) float64{ // 根据操作符计算
	res:=num1
	switch op {
		case '+':{
			res+=num2
		}
		case '-':{
			res-=num2
		}
		case '*':{
			res*=num2
		}
		case '/':{
			res/=num2
		}
	}
	return res
}





func convert(mid []interface{}) []interface{}{ //中缀表达式转逆波兰表达式
	suffix:=util.NewStack()
	chars:=util.NewStack()
	for i:=0;i<len(mid);i++{
		if op,ok:=mid[i].(byte);ok{
			if op=='('{
				chars.Push(op)
			}else if op==')'{
				for chars.PeekChar()!='('{
					suffix.Push(chars.PopChar())
				}
				chars.PopChar()
			}else {
				if chars.IsEmpty(){
					chars.Push(op)
				}else{
					if chars.PeekChar()==')'||opm[chars.PeekChar()]<opm[op]{
						chars.Push(op)
					}else {
						for !chars.IsEmpty()&&opm[chars.PeekChar()]>=opm[op]{
							suffix.Push(chars.PopChar())
						}
						chars.Push(op)
					}
				}
			}
		}else {
			suffix.Push(mid[i])
		}
	}
	for !chars.IsEmpty(){
		suffix.Push(chars.PopChar())
	}

	return suffix.ToSlice()
}






func readLine() string{
	s:=bufio.NewScanner(os.Stdin)
	s.Scan()
	return s.Text()
}


func parseEp(ep string)  []interface{}{  //将字符串转化为中缀表达式
	pre:=0
	res:=util.NewStack()
	for i:=0;i<len(ep);i++{
		if isOp(ep[i]){
			if pre<i && ep[i]!='('{
				//fmt.Println(pre,i)
				num:=ep[pre:i]
				f,err:=strconv.ParseFloat(trim(num),64)
				check(err)
				res.Push(f)

			}
			pre=i+1
			res.Push(ep[i])
		}
	}
	if pre<len(ep){
		num:=ep[pre:]
		f,err:=strconv.ParseFloat(trim(num),64)
		check(err)
		res.Push(f)
	}
	return res.ToSlice()
}

func trim(str string) string{ //去掉字符串首尾及中间的空格
	var res []byte
	for i:=0;i<len(str);i++{
		if str[i]!=' '{
			res=append(res,str[i])
		}
	}
	return string(res)


}


func isOp(ch byte) bool{ //判断是否为操作字符
	_,ok:=opm[ch]
	return ok
}
func check(err error){ //检查错误
	if err!=nil{
		fmt.Println(err)
	}
}