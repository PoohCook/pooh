package main

import (
	"fmt"
	"time"
	"./intervals"

)


///////////Derivations of Interval base/////////////////////////////////


type PoohInterval struct {
	intervals.IntervalBase
	Name  string
}

func (in PoohInterval ) String() string {
	return fmt.Sprintf( "PoohInterval||%s||%s", in.Name, in.IntervalBase.String() )
}

func (in PoohInterval ) Join( in2 intervals.Interval ) (PoohInterval, error){
	iResult, err := in.IntervalBase.Join(in2)
	if err == nil{
		return PoohInterval{iResult, in.Name}, nil
	}else {
		return PoohInterval{}, err
	}

}


//////////  Main ////////////////

func main() {
	t := time.Now()
	i1 := PoohInterval{intervals.IntervalBase{t, t.Add(time.Hour * 2)}, "Pooh Bear Time"}
	i2 := PoohInterval{intervals.IntervalBase{t.Add(time.Hour * 1), t.Add(time.Hour * 4)}, "Pooh Stick Time"}
	fmt.Println( i1)
	fmt.Println( i2)

	i3, err := i1.Join(i2)

	if(err == nil) {
		fmt.Println( i3)
	}else{
		fmt.Println(err)
	}
}
