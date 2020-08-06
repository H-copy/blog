package convert

import "testing"


func TestString(t *testing.T){

	msg := "msg"
	str := StrTo(msg)
	
	if  str.String() != msg{
		t.Fatal("String fail ", str.String())
	}

}

func TestInt(t *testing.T){

	str := StrTo("0")
	
	if v, err := str.Int(); v != 0{
		t.Fatal("Int fail", v, err)
	}

	str = StrTo("msg")

	if v, err := str.Int(); err == nil{
		t.Fatal("Int err fail", v, err)
	}
	
}

func TestMustInt(t *testing.T){

	str := StrTo("10")

	if num := str.MustInt(); num != 10{
		t.Fatal("MustInt fail", num)
	}
	
}

func TestUInt32(t *testing.T){

	var a uint32 = 10
	str := StrTo("10")

	if v, err := str.UInt32(); v != a{
		t.Fatal("uint32 fail", v, err)
	}

}

func TestMustUInt32(t *testing.T){

	var a uint32 = 10
	str := StrTo("10")

	if str.MustUInt32() != a{
		t.Fatal("MustUInt32 fail", a)
	}
	
}
