package intervals

import (
	"testing"
	"time"

	"fmt"
)


func TestMakeInterval(tst *testing.T) {

	t := time.Now()

	switch i1, err :=  MakeInterval(t, t.Add(time.Hour * 2));  {

	case err != nil:
		tst.Error(err)

	case !i1.Start().Equal(t) :
		tst.Errorf("unexpected start. Expected: %v  got:%v", t, i1.Start() );

	case !i1.End().Equal(t.Add(time.Hour*2)) :
		tst.Errorf("unexpected end. Expected: %v  got:%v", t.Add(time.Hour*2), i1.End() );

	}


	switch _, err :=  MakeInterval(t.Add(time.Hour * 2), t);  {

	case err == nil:
		tst.Error("expected invalid interval errror")

	case err.Error() != "Invalid Interval":
		tst.Errorf("unexpected Invalid Interval error message:%v", err.Error() );

	}
}

func TestIntervalContiguousJoin(tst *testing.T){

	t := time.Now()
	i1, _ := MakeInterval(t, t.Add(time.Hour * 2))
	i2, _ := MakeInterval(t.Add(time.Hour * 2), t.Add(time.Hour * 4))

	switch i3, err := i1.Join(i2);  {
	case err != nil:
		tst.Error(err)

	case !i3.Start().Equal(t) :
		tst.Errorf("unexpected start. Expected: %v  got:%v", t, i3.Start() );

	case !i3.End().Equal(t.Add(time.Hour*4)) :
		tst.Errorf("unexpected end. Expected: %v  got:%v", t.Add(time.Hour*4), i3.End() );

	}


	switch i3, err := i2.Join(i1);  {
	case err != nil:
		tst.Error(err)

	case !i3.Start().Equal(t) :
		tst.Errorf("unexpected start. Expected: %v  got:%v", t, i3.Start() );

	case !i3.End().Equal(t.Add(time.Hour*4)) :
		tst.Errorf("unexpected end. Expected: %v  got:%v", t.Add(time.Hour*4), i3.End() );

	}


}

func TestIntervalOverlapJoin(tst *testing.T){

	t := time.Now()
	i1, _ := MakeInterval(t, t.Add(time.Hour * 2))
	i2, _ := MakeInterval(t.Add(time.Hour * 1), t.Add(time.Hour * 4))

	switch i3, err := i1.Join(i2);  {
	case err != nil:
		tst.Error(err)

	case !i3.Start().Equal(t) :
		tst.Errorf("unexpected start. Expected: %v  got:%v", t, i3.Start() );

	case !i3.End().Equal(t.Add(time.Hour*4)) :
		tst.Errorf("unexpected end. Expected: %v  got:%v", t.Add(time.Hour*4), i3.End() );

	}


	switch i3, err := i2.Join(i1);  {
	case err != nil:
		tst.Error(err)

	case !i3.Start().Equal(t) :
		tst.Errorf("unexpected start. Expected: %v  got:%v", t, i3.Start() );

	case !i3.End().Equal(t.Add(time.Hour*4)) :
		tst.Errorf("unexpected end. Expected: %v  got:%v", t.Add(time.Hour*4), i3.End() );

	}

}

func TestIntervalNonContiguousJoin(tst *testing.T){

	t := time.Now()
	i1, _ := MakeInterval(t, t.Add(time.Hour * 2))
	i2, _ := MakeInterval(t.Add(time.Hour * 3), t.Add(time.Hour * 4))

	switch seq, err := i1.Join(i2);  {
	case err != nil:
		tst.Error(err)

	case len(seq) != 2:
		tst.Errorf("unexpected sequence size:%t    %t", len(seq), 2 );

	case !seq[0].Start().Equal(t) :
		tst.Errorf("unexpected start of first interval. Expected: %v  got:%v", t, seq[0].Start() );

	case !seq[0].End().Equal(t.Add(time.Hour*2)) :
		tst.Errorf("unexpected end of first interval. Expected: %v  got:%v", t.Add(time.Hour*2), seq[0].End() );

	case !seq[1].Start().Equal(t.Add(time.Hour*3)) :
		tst.Errorf("unexpected start of second interval. Expected: %v  got:%v", t.Add(time.Hour*3), seq[1].Start() );

	case !seq[1].End().Equal(t.Add(time.Hour*4)) :
		tst.Errorf("unexpected end of second interval. Expected: %v  got:%v", t.Add(time.Hour*4), seq[1].End() );

	case !seq.Start().Equal(t):
		tst.Errorf("unexpected start of sequence. Expected: %v  got:%v", t, seq[0].Start() );

	case !seq.End().Equal(t.Add(time.Hour*4)):
		tst.Errorf("unexpected start of sequence. Expected: %v  got:%v", t.Add(time.Hour*4), seq[0].End() );


	}


	switch seq, err := i2.Join(i1);  {
	case err != nil:
		tst.Error(err)

	case len(seq) != 2:
		tst.Errorf("unexpected sequence size:%t    %t", len(seq), 2 );

	case !seq[0].Start().Equal(t) :
		tst.Errorf("unexpected start of first interval. Expected: %v  got:%v", t, seq[0].Start() );

	case !seq[0].End().Equal(t.Add(time.Hour*2)) :
		tst.Errorf("unexpected end of first interval. Expected: %v  got:%v", t.Add(time.Hour*2), seq[0].End() );

	case !seq[1].Start().Equal(t.Add(time.Hour*3)) :
		tst.Errorf("unexpected start of second interval. Expected: %v  got:%v", t.Add(time.Hour*3), seq[1].Start() );

	case !seq[1].End().Equal(t.Add(time.Hour*4)) :
		tst.Errorf("unexpected end of second interval. Expected: %v  got:%v", t.Add(time.Hour*4), seq[1].End() );

	case !seq.Start().Equal(t):
		tst.Errorf("unexpected start of sequence. Expected: %v  got:%v", t, seq[0].Start() );

	case !seq.End().Equal(t.Add(time.Hour*4)):
		tst.Errorf("unexpected start of sequence. Expected: %v  got:%v", t.Add(time.Hour*4), seq[0].End() );

	}

}

func TestIntervalSplit(tst *testing.T){

	t := time.Now()
	i1, _ := MakeInterval(t, t.Add(time.Hour * 2))

	switch seq, err := i1.Split(t.Add(time.Hour * 1));  {


	case err != nil:
		tst.Error(err)

	case len(seq) != 2:
		tst.Errorf("unexpected sequence size:%t    %t", len(seq), 2 );

	case !seq[0].Start().Equal(t) :
		tst.Errorf("unexpected start of first split. Expected: %v  got:%v", t, seq[0].Start() );

	case !seq[0].End().Equal(t.Add(time.Hour*1)) :
		tst.Errorf("unexpected end of first split. Expected: %v  got:%v", t.Add(time.Hour*1), seq[0].End() );

	case !seq[1].Start().Equal(t.Add(time.Hour*1)) :
		tst.Errorf("unexpected start of second split. Expected: %v  got:%v", t.Add(time.Hour*1), seq[1].Start() );

	case !seq[1].End().Equal(t.Add(time.Hour*2)) :
		tst.Errorf("unexpected end of second split. Expected: %v  got:%v", t.Add(time.Hour*2), seq[1].End() );

	case !seq.Start().Equal(t):
		tst.Errorf("unexpected start of sequence. Expected: %v  got:%v", t, seq[0].Start() );

	case !seq.End().Equal(t.Add(time.Hour*2)):
		tst.Errorf("unexpected start of sequence. Expected: %v  got:%v", t, seq[0].Start() );



	}


}

func TestIntervalSplitOutOfBounds(tst *testing.T){

	t := time.Now()
	i1 := intervalBase{t, t.Add(time.Hour * 2)}

	switch _, err := i1.Split(t.Add(time.Hour * 4));  {


	case err == nil:
		tst.Error("Split Out Of bounds should have failed" );

	case err.Error() != "Split Out of Bounds" :
		tst.Errorf("unexpected error message for Split:%v", err );


	}

}

func TestIntervalSequence_String(tst *testing.T) {

	t := time.Now()
	t2 := t.Add(time.Hour * 2)
	i1 := intervalBase{t, t2}
	seq, _ := i1.Split(t.Add(time.Hour * 1))

	nm, _ := t.Zone()
	format := "Jan 2, 2006 3:04:05 PM"

	if seq.String() != fmt.Sprintf("Interval Sequence[%d]||%s -> %s: %s dur[%s]", 2, t.Format(format), t2.Format(format), nm, seq.End().Sub(seq.Start()) ){
		tst.Errorf("unexpected Seq String:%v", seq.String() );
	}

}

func TestIntervalSequence_Join(tst *testing.T) {
	t := time.Now()
	seq1, _ := MakeIntervalSequence(t, t.Add(time.Hour))
	seq2, _ := MakeIntervalSequence(t.Add(time.Hour*2), t.Add(time.Hour*3))

	switch seq, err := seq1.Join(seq2); {

	case err != nil:
		tst.Error(err)

	case len(seq) != 2:
		tst.Errorf("unexpected sequence size:%t    %t", len(seq), 2 );

	case !seq[0].Start().Equal(t) :
		tst.Errorf("unexpected start of first split. Expected: %v  got:%v", t, seq[0].Start() );

	case !seq[0].End().Equal(t.Add(time.Hour*1)) :
		tst.Errorf("unexpected end of first split. Expected: %v  got:%v", t.Add(time.Hour*1), seq[0].End() );

	case !seq[1].Start().Equal(t.Add(time.Hour*2)) :
		tst.Errorf("unexpected start of second split. Expected: %v  got:%v", t.Add(time.Hour*2), seq[1].Start() );

	case !seq[1].End().Equal(t.Add(time.Hour*3)) :
		tst.Errorf("unexpected end of second split. Expected: %v  got:%v", t.Add(time.Hour*3), seq[1].End() );

	case !seq.Start().Equal(t):
		tst.Errorf("unexpected start of sequence. Expected: %v  got:%v", t, seq[0].Start() );

	case !seq.End().Equal(t.Add(time.Hour*3)):
		tst.Errorf("unexpected start of sequence. Expected: %v  got:%v", t.Add(time.Hour*3), seq[0].Start() );



	}
}

func TestIntervalSequence_Split(tst *testing.T) {

	t := time.Now()
	seq1, _ := MakeIntervalSequence(t, t.Add(time.Hour *2))
	seq2, _ := MakeIntervalSequence(t.Add(time.Hour*2), t.Add(time.Hour*4))

	seq1, _ = seq1.Split(t.Add(time.Hour *1))
	seq2, _ = seq2.Split(t.Add(time.Hour *3))

	switch seq, err := seq1.Join(seq2); {

	case err != nil:
		tst.Error(err)

	case len(seq) != 4:
		tst.Errorf("unexpected sequence size:%t    %t", len(seq), 2 );

	case !seq[0].Start().Equal(t) :
		tst.Errorf("unexpected start of first split. Expected: %v  got:%v", t, seq[0].Start() );

	case !seq[0].End().Equal(t.Add(time.Hour*1)) :
		tst.Errorf("unexpected end of first split. Expected: %v  got:%v", t.Add(time.Hour*1), seq[0].End() );

	case !seq[1].Start().Equal(t.Add(time.Hour*1)) :
		tst.Errorf("unexpected start of second split. Expected: %v  got:%v", t.Add(time.Hour*1), seq[1].Start() );

	case !seq[1].End().Equal(t.Add(time.Hour*2)) :
		tst.Errorf("unexpected end of second split. Expected: %v  got:%v", t.Add(time.Hour*2), seq[1].End() );

	case !seq[2].Start().Equal(t.Add(time.Hour*2)) :
		tst.Errorf("unexpected start of second split. Expected: %v  got:%v", t.Add(time.Hour*2), seq[1].Start() );

	case !seq[2].End().Equal(t.Add(time.Hour*3)) :
		tst.Errorf("unexpected end of second split. Expected: %v  got:%v", t.Add(time.Hour*3), seq[1].End() );

	case !seq[3].Start().Equal(t.Add(time.Hour*3)) :
		tst.Errorf("unexpected start of second split. Expected: %v  got:%v", t.Add(time.Hour*3), seq[1].Start() );

	case !seq[3].End().Equal(t.Add(time.Hour*4)) :
		tst.Errorf("unexpected end of second split. Expected: %v  got:%v", t.Add(time.Hour*4), seq[1].End() );

	case !seq.Start().Equal(t):
		tst.Errorf("unexpected start of sequence. Expected: %v  got:%v", t, seq[0].Start() );

	case !seq.End().Equal(t.Add(time.Hour*4)):
		tst.Errorf("unexpected start of sequence. Expected: %v  got:%v", t.Add(time.Hour*4), seq[0].Start() );



	}

}




func TestIntervalSequence_MergedJoin(tst *testing.T) {

	t := time.Now()
	seq1, _ := MakeIntervalSequence(t, t.Add(time.Hour *2))
	seq2, _ := MakeIntervalSequence(t.Add(time.Minute*90), t.Add(time.Hour*4))

	seq1, _ = seq1.Split(t.Add(time.Hour *1))
	seq2, _ = seq2.Split(t.Add(time.Hour *3))

	switch seq, err := seq1.Join(seq2); {

	case err != nil:
		tst.Error(err)

	case len(seq) != 5:
		tst.Errorf("unexpected sequence size:%d  expected  %d", len(seq), 5 );
		tst.Errorf(seq.Details())

	case !seq[0].Start().Equal(t) :
		tst.Errorf("unexpected start of first split. Expected: %v  got:%v", t, seq[0].Start() );

	case !seq[0].End().Equal(t.Add(time.Hour*1)) :
		tst.Errorf("unexpected end of first split. Expected: %v  got:%v", t.Add(time.Hour*1), seq[0].End() );

	case !seq[1].Start().Equal(t.Add(time.Hour*1)) :
		tst.Errorf("unexpected start of second split. Expected: %v  got:%v", t.Add(time.Hour*1), seq[1].Start() );

	case !seq[1].End().Equal(t.Add(time.Minute*90)) :
		tst.Errorf("unexpected end of second split. Expected: %v  got:%v", t.Add(time.Minute*90), seq[1].End() );

	case !seq[2].Start().Equal(t.Add(time.Minute*90)) :
		tst.Errorf("unexpected start of second split. Expected: %v  got:%v", t.Add(time.Minute*90), seq[2].Start() );

	case !seq[2].End().Equal(t.Add(time.Hour*2)) :
		tst.Errorf("unexpected end of second split. Expected: %v  got:%v", t.Add(time.Hour*2), seq[2].End() );

	case !seq[3].Start().Equal(t.Add(time.Hour*3)) :
		tst.Errorf("unexpected start of second split. Expected: %v  got:%v", t.Add(time.Hour*3), seq[3].Start() );

	case !seq[3].End().Equal(t.Add(time.Hour*3)) :
		tst.Errorf("unexpected end of second split. Expected: %v  got:%v", t.Add(time.Hour*3), seq[3].End() );

	case !seq[4].Start().Equal(t.Add(time.Hour*3)) :
		tst.Errorf("unexpected start of second split. Expected: %v  got:%v", t.Add(time.Hour*3), seq[4].Start() );

	case !seq[4].End().Equal(t.Add(time.Hour*4)) :
		tst.Errorf("unexpected end of second split. Expected: %v  got:%v", t.Add(time.Hour*4), seq[4].End() );

	case !seq.Start().Equal(t):
		tst.Errorf("unexpected start of sequence. Expected: %v  got:%v", t, seq[0].Start() );

	case !seq.End().Equal(t.Add(time.Hour*4)):
		tst.Errorf("unexpected start of sequence. Expected: %v  got:%v", t.Add(time.Hour*4), seq[0].Start() );



	}

}
