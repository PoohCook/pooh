package intervals


import (
	"fmt"
	"time"

)

type Interval interface{
	Start() time.Time
	End() time.Time
}

type IntervalBase struct {
	Istart time.Time
	Iend   time.Time
}

func (in IntervalBase) Start() time.Time{
	return in.Istart
}
func (in IntervalBase) End() time.Time{
	return in.Iend
}

func (in IntervalBase ) String() string {
	nm, _ := in.Start().Zone()
	format := "Jan 2, 2006 3:04:05 PM"
	return fmt.Sprintf("Interval||%s -> %s: %s dur[%s]", in.Start().Format(format), in.End().Format(format), nm, in.End().Sub(in.Start()) )
}

func (in IntervalBase ) Join( in2 Interval ) (IntervalBase, error) {
	switch {
	case  in.Start().Before(in2.Start()) && !in.End().Before(in2.Start()):
		return IntervalBase{in.Start(), in2.End()}, nil
	case  in2.Start().Before(in.Start()) && !in2.End().Before(in.Start()):
		return IntervalBase{in2.Start(), in.End()}, nil
	default:
		return IntervalBase{}, fmt.Errorf("intervals are not contiguous and do not overlap")
	}

}

type IntervalSequence []IntervalBase

func (in IntervalBase) Split( time.Time ) (IntervalSequence, error) {


	seq := make([]IntervalBase, 2, 2 )
	seq[0] = IntervalBase{}
	seq[1] = IntervalBase{}
	return seq, fmt.Errorf("not implemented")
}


func (seq IntervalSequence) Start() time.Time{
	return seq[0].Start()
}
func (seq IntervalSequence) End() time.Time{
	return seq[len(seq)-1].End()
}