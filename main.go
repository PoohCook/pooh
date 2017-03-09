package main

import (
	"fmt"
	"time"
	"./intervals"

)


///////////Derivations of Interval base/////////////////////////////////


type PoohInterval struct {
	intervals.Interval
	Name  string
}

func makePoohInterval( start time.Time, end time.Time, name string) PoohInterval {
	in, _ := intervals.MakeInterval(start, end)
	return PoohInterval{in, name}
}

func (in PoohInterval ) String() string {
	return fmt.Sprintf( "PoohInterval||%s||%s", in.Name, in.Interval.Start().String() )
}

func (in PoohInterval ) Join( in2 PoohInterval ) (PoohInterval, error){
	iResult, err := in.Interval.Join(in2.Interval)
	if err == nil{
		return PoohInterval{iResult, in.Name}, nil
	}else {
		return PoohInterval{}, err
	}

}


//////////  Main ////////////////

func main() {
	t := time.Now()
	i1 := makePoohInterval(t, t.Add(time.Hour * 2), "Pooh Bear Time")
	i2 := makePoohInterval(t.Add(time.Hour * 1), t.Add(time.Hour * 4), "Pooh Stick Time")
	fmt.Println( i1)
	fmt.Println( i2)

	i3, err := i1.Join(i2)

	if(err == nil) {
		fmt.Println( i3)
	}else{
		fmt.Println(err)
	}
}
