package intervals


import (
	"fmt"
	"time"

	"errors"
)



type Interval interface{
	Start() time.Time
	End() time.Time
	Join( in2 Interval ) (IntervalSequence, error)
	Split( spTime time.Time ) (IntervalSequence, error)

}

type intervalBase struct {
	Istart time.Time
	Iend   time.Time
}

func MakeInterval(start time.Time, end time.Time) (Interval, error) {
	if start.After(end){return intervalBase{}, errors.New("Invalid Interval")}
	return intervalBase{start, end}, nil
}

func (in intervalBase) Start() time.Time{
	return in.Istart
}
func (in intervalBase) End() time.Time{
	return in.Iend
}

func (in intervalBase ) String() string {
	nm, _ := in.Start().Zone()
	format := "Jan 2, 2006 3:04:05 PM"
	return fmt.Sprintf("Interval||%s -> %s: %s dur[%s]", in.Start().Format(format), in.End().Format(format), nm, in.End().Sub(in.Start()) )
}

func (in intervalBase ) Join( in2 Interval ) (IntervalSequence, error) {
	switch {
	case  in.Start().Before(in2.Start()) && !in.End().Before(in2.Start()):
		return IntervalSequence{intervalBase{in.Start(), in2.End()}}, nil
	case  in2.Start().Before(in.Start()) && !in2.End().Before(in.Start()):
		return IntervalSequence{intervalBase{in2.Start(), in.End()}}, nil
	case  in.End().Before(in2.Start()):
		return IntervalSequence{in,in2}, nil
	default:
		return IntervalSequence{in2,in}, nil
	}

}

func (in intervalBase) Split( spTime time.Time ) ( rSeq IntervalSequence, err error) {

	if spTime.After(in.End()) || spTime.Before(in.Start()) {
		err = errors.New("Split Out of Bounds")
		return
	}
	i1, err := MakeInterval(in.Start(),spTime)
	i2, err := MakeInterval(spTime,in.End())

	rSeq = IntervalSequence{i1, i2}

	return
}



type IntervalSequence []Interval

func MakeIntervalSequence( start, end time.Time) ( rSeq IntervalSequence, err error) {
	i1, err := MakeInterval(start, end)

	rSeq = IntervalSequence{i1}
	return
}

func (seq IntervalSequence) Start() time.Time{
	return seq[0].Start()
}

func (seq IntervalSequence) End() time.Time{
	return seq[len(seq)-1].End()
}

func (seq IntervalSequence ) String() string {
	nm, _ := seq.Start().Zone()
	format := "Jan 2, 2006 3:04:05 PM"
	return fmt.Sprintf("Interval Sequence[%d]||%s -> %s: %s dur[%s]", len(seq), seq.Start().Format(format), seq.End().Format(format), nm, seq.End().Sub(seq.Start()) )
}

func (seq IntervalSequence ) Join( in2 Interval ) (IntervalSequence, error) {
	return IntervalSequence{}, errors.New("Not Implmented")
}

func (seq IntervalSequence) Split( spTime time.Time ) (  IntervalSequence,  error) {
	return IntervalSequence{}, errors.New("Not Implmented")
}

