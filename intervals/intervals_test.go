package intervals

import (
	"testing"
	"time"

	"fmt"
)




func TestIntervalContiguousJoin(tst *testing.T){

	t := time.Now()
	i1 := IntervalBase{t, t.Add(time.Hour * 2)}
	i2 := IntervalBase{t.Add(time.Hour * 2), t.Add(time.Hour * 4)}

	switch i3, err := i1.Join(i2);  {
	case err != nil:
		tst.Error(err)

	case !i3.Start().Equal(t) :
		tst.Errorf("unexpected start. Expected: %v  got:%v", t, i3.Start() );
		fallthrough

	case !i3.End().Equal(t.Add(time.Hour*4)) :
		tst.Errorf("unexpected end. Expected: %v  got:%v", t.Add(time.Hour*4), i3.End() );

	}


	switch i3, err := i2.Join(i1);  {
	case err != nil:
		tst.Error(err)

	case !i3.Start().Equal(t) :
		tst.Errorf("unexpected start. Expected: %v  got:%v", t, i3.Start() );
		fallthrough

	case !i3.End().Equal(t.Add(time.Hour*4)) :
		tst.Errorf("unexpected end. Expected: %v  got:%v", t.Add(time.Hour*4), i3.End() );

	}


}


func TestIntervalOverlapJoin(tst *testing.T){

	t := time.Now()
	i1 := IntervalBase{t, t.Add(time.Hour * 2)}
	i2 := IntervalBase{t.Add(time.Hour * 1), t.Add(time.Hour * 4)}

	switch i3, err := i1.Join(i2);  {
	case err != nil:
		tst.Error(err)

	case !i3.Start().Equal(t) :
		tst.Errorf("unexpected start. Expected: %v  got:%v", t, i3.Start() );
		fallthrough

	case !i3.End().Equal(t.Add(time.Hour*4)) :
		tst.Errorf("unexpected end. Expected: %v  got:%v", t.Add(time.Hour*4), i3.End() );

	}


	switch i3, err := i2.Join(i1);  {
	case err != nil:
		tst.Error(err)

	case !i3.Start().Equal(t) :
		tst.Errorf("unexpected start. Expected: %v  got:%v", t, i3.Start() );
		fallthrough

	case !i3.End().Equal(t.Add(time.Hour*4)) :
		tst.Errorf("unexpected end. Expected: %v  got:%v", t.Add(time.Hour*4), i3.End() );

	}

}



func TestIntervalNonContiguousJoin(tst *testing.T){

	t := time.Now()
	i1 := IntervalBase{t, t.Add(time.Hour * 2)}
	i2 := IntervalBase{t.Add(time.Hour * 3), t.Add(time.Hour * 4)}

	switch _, err := i1.Join(i2);  {
	case err == nil:
		tst.Error("expected excpetion for non contiguous interval but no eception was returned")

	case fmt.Sprint(err) != "intervals are not contiguous and do not overlap" :
		tst.Errorf("unexpected error message:%v",  );

	}


	switch _, err := i2.Join(i1);  {
	case err == nil:
		tst.Error("expected excpetion for non contiguous interval but no eception was returned")

	case fmt.Sprint(err) != "intervals are not contiguous and do not overlap" :
		tst.Errorf("unexpected error message:%v",  );

	}

}



func TestIntervalSplit(tst *testing.T){

	t := time.Now()
	i1 := IntervalBase{t, t.Add(time.Hour * 2)}

	switch seq, err := i1.Split(t.Add(time.Hour * 1));  {
	case err != nil:
		tst.Error(err)

	case len(seq) != 2:
		tst.Errorf("unexpected sequence size:%v", len(seq) );

	case !seq[0].Start().Equal(t) :
		tst.Errorf("unexpected start of first split. Expected: %v  got:%v", t, seq[0].Start() );
		fallthrough

	case !seq[0].End().Equal(t.Add(time.Hour*1)) :
		tst.Errorf("unexpected end of first split. Expected: %v  got:%v", t.Add(time.Hour*1), seq[0].End() );
		fallthrough

	case !seq[1].Start().Equal(t.Add(time.Hour*1)) :
		tst.Errorf("unexpected start of second split. Expected: %v  got:%v", t.Add(time.Hour*1), seq[1].Start() );
		fallthrough

	case !seq[1].End().Equal(t.Add(time.Hour*2)) :
		tst.Errorf("unexpected end of second split. Expected: %v  got:%v", t.Add(time.Hour*2), seq[1].End() );

	}


}