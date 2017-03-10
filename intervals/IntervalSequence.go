package intervals

import (
	"fmt"
	"time"
	"errors"
)

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


	switch in2.(type) {
	case IntervalSequence:
		return seq.insertSequence(in2.(IntervalSequence) )
	default:
		return seq.insertInterval(in2)
	}


}

func( seq IntervalSequence) insertSequence ( seq2 IntervalSequence) (IntervalSequence, error) {

	err := errors.New("O length sequence insert");

	for i := range seq2 {
		seq, err = seq.insertInterval(seq2[i])
		if( err != nil ) {
			return IntervalSequence{}, err
		}
	}

	return seq, err;


}



func( seq IntervalSequence) insertInterval ( in2 Interval) (IntervalSequence, error) {


	for i := range seq {
		thisIntv := seq[i]
		switch{

		case !in2.End().After(thisIntv.Start()):
			return  seq.insertIntervalAtPos(i+1, thisIntv)

		case in2.Start().Before(thisIntv.End()):
			return seq, errors.New("Merge not implemented")


		}

	}

	return seq.insertIntervalAtPos(len(seq), in2)



}


func( seq IntervalSequence) insertIntervalAtPos ( pos int, in2 Interval) (IntervalSequence, error) {

	if pos < 0 || pos > len(seq){
		return seq, fmt.Errorf("Insert pos[%d] is out of bounds[0:%d]", pos, len(seq))
	}
	seq= append(seq, in2)
	copy(seq[pos+1:], seq[pos:])
	seq[pos] = in2

	return seq, nil
}


func (seq IntervalSequence) Split( spTime time.Time ) (  IntervalSequence,  error) {

	for i := range seq {
		interval := seq[i]
		if(interval.Start().Before(spTime) && interval.End().After(spTime)){
			splitSeq, err := interval.Split(spTime)
			if err != nil {
				return seq, err
			}
			seq[i] = splitSeq[0]
			seq, err = seq.insertIntervalAtPos(i+1, splitSeq[1])

			return seq, err
		}

	}

	return seq, nil
}

func(seq IntervalSequence) Merge( in2 Interval )( IntervalSequence, error){
	return IntervalSequence{}, errors.New("Not Implemented")
}



func(seq IntervalSequence) Details() string {
	retString := seq.String() + "\n"

	for i := range seq {
		retString += fmt.Sprintln(seq[i])
	}

	return retString
}