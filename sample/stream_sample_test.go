package sample

import (
	"strconv"
	"testing"
)

func TestAdd(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{
			Str: "",
			Int: i,
		})
		if stream.Last().Int != i {
			t.Fatal("Unexpect Value: Add ", stream)
		}
	}
	if stream.Last().Int != 9999 || stream.First().Int != 0 || stream.Len() != 10000 {
		t.Fatal("Unexpect Value: Add ", stream)
	}
}
func TestAddAll(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.AddAll(Sample{
			Str: "",
			Int: i,
		}, Sample{
			Str: "",
			Int: i,
		}, Sample{
			Str: "",
			Int: i,
		}, Sample{
			Str: "",
			Int: i,
		}, Sample{
			Str: "",
			Int: i,
		})
		if stream.Last().Int != i {
			t.Fatal("Unexpect Value: AddAll ", stream)
		}
	}
	stream.ForEach(func(sample Sample, i int) {
		if sample.Int != i/5 {
			t.Fatal("Unexpect Value: AddAll ", stream)
		}
	})
	if stream.Last().Int != 9999 || stream.First().Int != 0 || stream.Len() != 10000*5 {
		t.Fatal("Unexpect Value: AddAll ", stream)
	}
}
func TestAddSafe(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.AddSafe(nil)
	}
	if stream.IsPreset() {
		t.Fatal("Unexpect Value: AddSafe ", stream)
	}
	for i := 0; i < 10000; i++ {
		stream.AddSafe(&Sample{
			Str: "",
			Int: i,
		})
	}
	stream.ForEach(func(sample Sample, i int) {
		if sample.Int != i {
			t.Fatal("Unexpect Value: AddSafe ", stream)
		}
	})
}
func TestAllMatch(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	if stream.AllMatch(func(sample Sample, i int) bool {
		return sample.Int%2 == 0
	}) {
		t.Fatal("Unexpect Value: AllMatch ", stream)
	}
	if !stream.AllMatch(func(sample Sample, i int) bool {
		return true
	}) {
		t.Fatal("Unexpect Value: AllMatch ", stream)
	}

}
func TestAnyMatch(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	if !stream.AnyMatch(func(sample Sample, i int) bool {
		return sample.Int%2 == 0
	}) {
		t.Fatal("Unexpect Value: AllMatch ", stream)
	}
	if !stream.AnyMatch(func(sample Sample, i int) bool {
		return true
	}) {
		t.Fatal("Unexpect Value: AllMatch ", stream)
	}
}
func TestClone(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	copied := stream.Clone()
	copied.Set(3, Sample{
		Str: "",
		Int: 99,
	})
	if copied.Get(3).Int != 99 || stream.Get(3).Int == 99 {
		t.Fatal("Unexpect Value: Clone ", stream)
	}
	copied.Clean()
	if copied.IsPreset() || stream.IsEmpty() {
		t.Fatal("Unexpect Value: Clone ", stream)
	}
	if copied == &stream {
		t.Fatal("Unexpect Value: Clone ", stream)
	}
}
func TestCopy(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	if tmp := stream.Copy(); tmp == &stream {
		t.Fatal("Unexpect Value: Copy ", stream)
	}
}
func TestConcat(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		tmp := SampleStreamOf(Sample{
			Str: "",
			Int: i,
		}, Sample{
			Str: "",
			Int: i,
		}, Sample{
			Str: "",
			Int: i,
		}, Sample{
			Str: "",
			Int: i,
		}, Sample{
			Str: "",
			Int: i,
		})
		stream.Concat(tmp)
		if stream.Last().Int != i {
			t.Fatal("Unexpect Value: Concat ", stream)
		}
	}
	stream.ForEach(func(sample Sample, i int) {
		if sample.Int != i/5 {
			t.Fatal("Unexpect Value: Concat ", stream)
		}
	})
	if stream.Last().Int != 9999 || stream.First().Int != 0 || stream.Len() != 10000*5 {
		t.Fatal("Unexpect Value: Concat ", stream)
	}
}
func TestContains(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	if !stream.Contains(Sample{Int: 100}) ||
		stream.Contains(Sample{Int: -1}) ||
		stream.Contains(Sample{Int: 10000}) ||
		stream.Contains(Sample{Str: "1000"}) ||
		stream.Contains(Sample{Str: "test"}) {
		t.Fatal("Unexpect Value: Contains ", stream)
	}
}
func TestDelete(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	if stream.Delete(100); stream.Get(99).Int != 99 ||
		stream.Get(100).Int != 101 ||
		stream.Len() != 9999 {
		t.Fatal("Unexpect Value: Delete ", stream)
	}
	stream.ForEach(func(sample Sample, i int) {
		if i < 100 {
			if sample.Int != i {
				t.Fatal("Unexpect Value: Delete ", stream)
			}
		} else if sample.Int != i+1 {
			t.Fatal("Unexpect Value: Delete ", stream)
		}
	})
	stream.Delete(100)
	stream.ForEach(func(sample Sample, i int) {
		if i < 100 {
			if sample.Int != i {
				t.Fatal("Unexpect Value: Delete ", stream)
			}
		} else if sample.Int != i+2 {
			t.Fatal("Unexpect Value: Delete ", stream)
		}
	})
}
func TestDeleteRange(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	stream.DeleteRange(10, 100)
	stream.ForEach(func(sample Sample, i int) {
		if i < 10 {
			if sample.Int != i {
				t.Fatal("Unexpect Value: Delete ", stream)
			}
		} else {
			if sample.Int != i+91 {
				t.Fatal("Unexpect Value: Delete ", stream)
			}
		}
	})
}
func TestDistinct(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	stream.Concat(stream.Copy().Val())
	stream.Concat(stream.Copy().Val())
	stream.Concat(stream.Copy().Val())
	if stream.Len() < 10000*4 {
		t.Fatal("Unexpect Value: Distinct ", stream)
	}
	if stream.Distinct(); stream.Len() != 10000 {
		t.Fatal("Unexpect Value: Distinct ", stream.Len())
	}
	stream.ForEach(func(sample1 Sample, i int) {
		stream.ForEach(func(sample2 Sample, j int) {
			if i != j && (sample1.Int == sample2.Int || &sample1 == &sample2) {
				t.Fatal("Unexpect Value: Distinct ", i, sample1, j, sample2)
			}
		})
	})
}
func TestEach(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	count := 0
	stream.Each(func(sample Sample) {
		if sample.Int != stream[count].Int || &sample == &stream[count] {
			t.Fatal("Unexpect Value: Each ", stream)
		}
		count++
	})
}
func TestEachRight(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	count := stream.Len() - 1
	stream.EachRight(func(sample Sample) {
		if sample.Int != stream[count].Int || &sample == &stream[count] {
			t.Fatal("Unexpect Value: Each ", sample, stream[count])
		}
		count--
	})
}
func TestEquals(t *testing.T) {
	stream := SampleStreamOf(Sample{
		Str: "test1",
		Int: 999,
	}, Sample{
		Str: "test2",
		Int: -999,
	})

	if !stream.Equals([]Sample{
		{
			Str: "test1",
			Int: 999,
		},
		{
			Str: "test2",
			Int: -999,
		},
	}) || stream.Equals([]Sample{
		{
			Str: "test1",
			Int: 9199,
		},
		{
			Str: "test4",
			Int: -999,
		},
	}) {
		t.Fatal("Unexpect Value: Equals ", stream)
	}
}
func TestFilter(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	stream.
		Filter(func(sample Sample, i int) bool { return sample.Int%2 == 0 }).
		Each(func(sample Sample) {
			if sample.Int%2 != 0 {
				t.Fatal("Unexpect Value: Equals ", sample)
			}
		})
}
func TestFilterSlim(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	stream.
		FilterSlim(func(sample Sample, i int) bool { return sample.Int%2 == 0 }).
		Each(func(sample Sample) {
			if sample.Int%2 != 0 {
				t.Fatal("Unexpect Value: Equals ", sample)
			}
		})
}
func TestFind(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	if v := stream.Find(func(sample Sample, i int) bool {
		return sample.Int == 999
	}); v == nil || v.Int != 999 {
		t.Fatal("Unexpect Value: Find ", stream)
	}
	if v := stream.Find(func(sample Sample, i int) bool {
		return sample.Int == -1
	}); v != nil {
		t.Fatal("Unexpect Value: Find ", stream)
	}
}
func TestFindOr(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	if v := stream.FindOr(func(sample Sample, i int) bool {
		return sample.Int == 999
	}, Sample{
		Str: "",
		Int: 0,
	}); v.Int != 999 {
		t.Fatal("Unexpect Value: Find ", stream)
	}
	if v := stream.FindOr(func(sample Sample, i int) bool {
		return sample.Int == -1
	}, Sample{
		Str: "",
		Int: -99,
	}); v.Int != -99 {
		t.Fatal("Unexpect Value: Find ", stream)
	}
}
func TestFindIndex(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	if v := stream.FindIndex(func(sample Sample, i int) bool {
		return sample.Int == 999
	}); v != 999 {
		t.Fatal("Unexpect Value: Find ", v)
	}
	if v := stream.FindIndex(func(sample Sample, i int) bool {
		return sample.Int == -1
	}); v != -1 {
		t.Fatal("Unexpect Value: Find ", stream)
	}
}
func TestFirst(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	if v := stream.First(); stream.Get(0).Int != v.Int {
		t.Fatal("Unexpect Value: First ", v)
	}
	stream.ForEach(func(sample Sample, i int) {
		if i != 0 && stream.Get(i).Int == stream.First().Int {
			t.Fatal("Unexpect Value: First ", stream)
		}
	})

}
func TestFirstOr(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	if v := stream.FirstOr(Sample{}); stream.Get(0).Int != v.Int {
		t.Fatal("Unexpect Value: First ", v)
	}
	stream.ForEach(func(sample Sample, i int) {
		if i != 0 && stream.Get(i).Int == stream.FirstOr(Sample{}).Int {
			t.Fatal("Unexpect Value: First ", stream)
		}
	})
	stream = SampleStreamOf()
	if v := stream.FirstOr(Sample{Int: -999}); v.Int != -999 {
		t.Fatal("Unexpect Value: FirstOr ", stream)
	}

}
func TestForEach(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	count := 0
	stream.ForEach(func(sample Sample, i int) {
		if sample.Int != stream[count].Int || &sample == &stream[count] {
			t.Fatal("Unexpect Value: Each ", stream)
		}
		count++
	})
}
func TestForEachRight(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	count := stream.Len() - 1
	stream.ForEachRight(func(sample Sample, i int) {
		if sample.Int != stream[count].Int || &sample == &stream[count] {
			t.Fatal("Unexpect Value: Each ", sample, stream[count])
		}
		count--
	})
}
func TestGroupBy(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	for k, v := range stream.GroupBy(func(sample Sample, i int) string {
		return strconv.Itoa(sample.Int % 5)
	}) {
		if i, _ := strconv.Atoi(k); v[0].Int%5 != i {
			t.Fatal("Unexpect Value: GroupBy ", k, v)
		}
	}
}
func TestGroupByValues(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	for _, v := range stream.GroupByValues(func(sample Sample, i int) string {
		return strconv.Itoa(sample.Int % 5)
	}) {
		tmp := v[0].Int % 5
		for _, _v := range v {
			if _v.Int%5 != tmp {
				t.Fatal("Unexpect Value: GroupBy ", v, _v)
			}
		}
	}
}
func TestIndexOf(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	if v := stream.IndexOf(Sample{
		Str: "",
		Int: 999,
	}); v != 999 {
		t.Fatal("Unexpect Value: IndexOf ", v)
	}
	if v := stream.IndexOf(Sample{
		Str: "",
		Int: 22,
	}); v != 22 {
		t.Fatal("Unexpect Value: IndexOf ", v)
	}
}
func TestIsEmpty(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	if stream.IsEmpty() {
		t.Fatal("Unexpect Value: IsEmpty ", stream)
	}
	stream = SampleStreamOf()
	if !stream.IsEmpty() {
		t.Fatal("Unexpect Value: IsEmpty ", stream)
	}
}
func TestIsPreset(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	if !stream.IsPreset() {
		t.Fatal("Unexpect Value: IsPreset ", stream)
	}
	stream = SampleStreamOf()
	if stream.IsPreset() {
		t.Fatal("Unexpect Value: IsPreset ", stream)
	}
}
func TestLast(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	v := stream.Last()
	if v.Int != 9999 {
		t.Fatal("Unexpect Value: Last ", v)
	}
	if v.Int = 81; stream.Last().Int != 9999 {
		t.Fatal("Unexpect Value: Last ", v)
	}
}
func TestLastOr(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	v := stream.LastOr(Sample{})
	if v.Int != 9999 {
		t.Fatal("Unexpect Value: LastOr ", v)
	}
	if v.Int = 81; stream.LastOr(Sample{}).Int != 9999 {
		t.Fatal("Unexpect Value: LastOr ", v)
	}
	stream = SampleStreamOf()
	if last := stream.LastOr(Sample{Int: 99}); last.Int != 99 {
		t.Fatal("Unexpect Value: LastOr ", v)
	}
}
func TestLen(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	if stream.Len() != len(stream) {
		t.Fatal("Unexpect Value: Len ", stream)
	}
}
func TestLimit(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	if stream.Limit(100); stream.Len() != 100 {
		t.Fatal("Unexpect Value: Limit ", stream.Len())
	}
	stream.ForEach(func(sample Sample, i int) {
		if sample.Int > 100 {
			t.Fatal("Unexpect Value: Limit ", sample)
		}
	})
	if stream.Limit(-1); stream.Len() != 0 {
		t.Fatal("Unexpect Value: Limit ", stream)
	}
}
func TestNoneMatch(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	if stream.NoneMatch(func(sample Sample, i int) bool {
		return true
	}) {
		t.Fatal("Unexpect Value: NoneMatch ", stream)
	}
	if !stream.NoneMatch(func(sample Sample, i int) bool {
		return false
	}) {
		t.Fatal("Unexpect Value: NoneMatch ", stream)
	}
}
func TestGet(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	if v := stream.Get(-1); v != nil {
		t.Fatal("Unexpect Value: Get ", stream)
	}
	if v := stream.Get(0); v.Int != stream.First().Int {
		t.Fatal("Unexpect Value: Get ", stream)
	}
	v := stream.Get(0)
	v.Int += 100
	if v := stream.Get(0); v.Int != stream.First().Int {
		t.Fatal("Unexpect Value: Get ", stream)
	}
}
func TestGetOr(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	if v := stream.GetOr(-1, Sample{Int: 99}); v.Int != 99 {
		t.Fatal("Unexpect Value: Get ", stream)
	}
	if v := stream.GetOr(0, Sample{}); v.Int != stream.First().Int {
		t.Fatal("Unexpect Value: Get ", stream)
	}
	v := stream.GetOr(0, Sample{})
	v.Int += 100
	if v := stream.GetOr(0, Sample{}); v.Int != stream.First().Int {
		t.Fatal("Unexpect Value: Get ", stream)
	}
}
func TestReverse(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	c := stream.Clone().Reverse()
	stream.ForEach(func(sample Sample, i int) {
		if sample.Int != c.Get(c.Len()-i-1).Int {
			t.Fatal("Unexpect Value: Get ", sample.Int, c.Get(c.Len()-i-1))
		}
	})
}
func TestSet(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	stream.Set(91, Sample{
		Str: "",
		Int: -999,
	})
	if stream.Get(91).Int != -999 {
		t.Fatal("Unexpect Value: Get ", stream.Get(91))
	}
}
func TestSkip(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	if stream.Skip(5000); stream.Len() != 5000 {
		t.Fatal("Unexpect Value: Get ", stream.Len())
	}
	stream.ForEach(func(sample Sample, i int) {
		if sample.Int < 5000 {
			t.Fatal("Unexpect Value: Get ", sample)
		}
	})
	if stream.Skip(5000000); stream.Len() != 0 {
		t.Fatal("Unexpect Value: Get ", stream)
	}
	if stream.Skip(-5000000); stream.Len() != 0 {
		t.Fatal("Unexpect Value: Get ", stream)
	}
}
func TestSlice(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	c := stream.Clone()
	if stream.Slice(100, 100); stream.Len() != 100 {
		t.Fatal("Unexpect Value: Slice ", stream.Len())
	}
	for i := 100; i < 200; i++ {
		if c.Get(i).Int != stream.Get(i-100).Int {
			t.Fatal("Unexpect Value: Get ", c.Get(i), stream.Get(i-100))
		}
	}
}
func TestSort(t *testing.T) {
	stream := SampleStreamOf()
	for i := 0; i < 10000; i++ {
		stream.Add(Sample{Int: i})
	}
	c := stream.Clone().Sort(func(i, j int) bool {
		return stream.Get(i).Int > stream.Get(j).Int
	})
	stream.Reverse().ForEach(func(sample Sample, i int) {
		if c.Get(i).Int != sample.Int {
			t.Fatal("Unexpect Value: Slice ", c.Get(i), sample)
		}
	})
}

//func Test(t *testing.T) {
//	stream := SampleStreamOf()
//	stream.AddAll(
//		Sample{
//			Str: "1",
//			Int: 1,
//		},
//		Sample{
//			Str: "2",
//			Int: 2,
//		},
//		Sample{
//			Str: "3",
//			Int: 3,
//		},
//		Sample{
//			Str: "4",
//			Int: 4,
//		},
//		Sample{
//			Str: "5",
//			Int: 5,
//		},
//	)
//	cloned1 := stream.Clone()
//	cloned2 := CreateSampleStream(stream.Clone().Val()...)
//	cloned3 := GenerateSampleStream(stream.Clone().Val())
//
//	stream.ForEach(func(arg Sample, index int) {
//		if arg.Int != index+1 {
//			t.Fatal("Unexpect Value: AddAll ", index)
//		}
//	})
//	tmp := cloned2.Val()
//	if tmp[0].Str = "aaa"; cloned2.Get(0).Str == "aaa" {
//		t.Fatal("Unexpect Value stream Val.", cloned2)
//	}
//	if stream.Add(Sample{Str: "6", Int: 6}); stream.Len() != 6 {
//		t.Fatal("Unexpect Value stream length.", stream)
//	}
//	if stream.AnyMatch(func(sample Sample, _ int) bool {
//		sample.Int = 99999
//		return true
//	}); stream.Contains(Sample{Str: "", Int: 99999}) {
//		t.Fatal("Unexpect Value stream AllMatch.", stream)
//	}
//	if !stream.AllMatch(func(_ Sample, _ int) bool { return true }) ||
//		stream.AllMatch(func(_ Sample, _ int) bool { return false }) {
//		t.Fatal("Unexpect Value stream AllMatch.", stream)
//	}
//	if stream.Concat([]Sample{{Str: "7", Int: 7}, {Str: "8", Int: 8}}); stream.Get(6).Str != "7" || stream.Get(6).Int != 7 {
//		t.Fatal("Unexpect Value stream Concat.", stream)
//	}
//	if cloned1.Delete(0); cloned1.Len() == 6 || stream.Len() == cloned1.Len() || cloned1.Get(0).Str != stream.Get(1).Str {
//		t.Fatal("Unexpect Value stream Delete.", cloned1)
//	}
//	if cloned1.DeleteRange(0, 2); cloned1.Get(0).Int != 5 {
//		t.Fatal("Unexpect Value stream DeleteRange.", cloned1)
//	}
//	if cloned1 = stream.Copy(); !reflect.DeepEqual(cloned1.Last(), stream.Last()) {
//		t.Fatal("Unexpect Value stream Copy.", cloned1)
//	}
//	if cloned1.Filter(func(arg Sample, _ int) bool { return arg.Int%2 == 0 }); cloned1.Len() != 4 || cloned1.Get(0).Str != "2" {
//		t.Fatal("Unexpect Value stream Filter.", cloned1)
//	}
//	if val := stream.Find(func(arg Sample, _ int) bool { return arg.Str == "5" }); cloned1.Find(func(arg Sample, _ int) bool { return arg.Str == "5" }) != nil || val.Str != "5" {
//		t.Fatal("Unexpect Value stream Find.", cloned1)
//	}
//	if index := stream.FindIndex(func(arg Sample, _ int) bool { return arg.Int == 8 }); index != stream.Len()-1 {
//		t.Fatal("Unexpect Value stream FindIndex.", stream)
//	}
//	if !reflect.DeepEqual(stream.First(), stream.Get(0)) {
//		t.Fatal("Unexpect Value stream First.", stream)
//	}
//	if m := stream.GroupBy(func(arg Sample, _ int) string { return strconv.Itoa(arg.Int % 4) }); len(m["1"]) != 2 || m["1"][0].Int != 1 || m["2"][0].Int != 2 || m["3"][0].Int != 3 {
//		t.Fatal("Unexpect Value stream GroupBy.", m)
//	}
//	if v := stream.GroupByValues(func(arg Sample, _ int) string { return strconv.Itoa(arg.Int % 4) }); len(v) != 4 {
//		t.Fatal("Unexpect Value stream GroupByValues.", v)
//	}
//	if tmp := SampleStreamOf(); stream.IsEmpty() || !tmp.IsEmpty() {
//		t.Fatal("Unexpect Value stream IsEmpty.", tmp)
//	}
//	if tmp := SampleStreamOf(); !stream.IsPreset() || tmp.IsPreset() {
//		t.Fatal("Unexpect Value stream IsPreset.", tmp)
//	}
//	//if cloned2.Map(func(arg Sample, index int) Sample { return Sample{Str: "test", Int: index} }); cloned2.First().Str != "test" && cloned2.First().Int != 0 && cloned2.Last().Str != "test" && cloned2.Last().Int != 4 {
//	//	t.Fatal("Unexpect Value stream Map.", cloned2)
//	//}
//	if cloned2.Add(Sample{Str: "last", Int: 999}); cloned2.Last().Str != "last" || cloned2.Last().Int != 999 {
//		t.Fatal("Unexpect Value stream make slice.", cloned2)
//	}
//	if !cloned2.NoneMatch(func(_ Sample, _ int) bool { return false }) || cloned1.NoneMatch(func(_ Sample, _ int) bool { return true }) {
//		t.Fatal("Unexpect Value stream NoneMatch.", cloned2)
//	}
//	if stream.Get(8888) != nil || stream.Get(0) == nil || stream.Get(stream.Len()-1) == nil || stream.Get(-1) != nil {
//		t.Fatal("Unexpect Value stream Get.", stream)
//	}
//	if cloned3.ReduceInit(func(result, current Sample, index int) Sample { current.Int += result.Int; return current }, Sample{Int: 0}); cloned3.Last().Int != 15 {
//		t.Fatal("Unexpect Value stream ReduceInit.", cloned3)
//	}
//	if cloned3.Replace(func(arg Sample, index int) Sample { return Sample{Str: "test", Int: 5} }); cloned2.First().Str != "test" && cloned2.First().Int != 5 && cloned2.Last().Str != "test" && cloned2.Last().Int != 5 {
//		t.Fatal("Unexpect Value stream Map.", cloned2)
//	}
//	if cloned3.Reduce(func(result, current Sample, index int) Sample { current.Int += result.Int; return current }); cloned3.Last().Int != 25 {
//		t.Fatal("Unexpect Value stream Reduce.", cloned3)
//	}
//	if cloned2.Reverse(); cloned2.First().Int != 999 || cloned2.Get(1).Int != 4 || cloned2.Get(2).Int != 3 || cloned2.Get(3).Int != 2 || cloned2.Get(4).Int != 1 || cloned2.Last().Int != 0 {
//		t.Fatal("Unexpect Value stream Reverse.", cloned3)
//	}
//	if cloned2 = stream.Copy(); !stream.Equals(cloned2.Val()) || stream.Equals(cloned3.Val()) {
//		t.Fatal("Unexpect Value stream Copy.", stream)
//	}
//	if index := stream.IndexOf(*stream.Get(3)); index != 3 {
//		t.Fatal("Unexpect Value stream IndexOf.", stream)
//	}
//	if index := stream.IndexOf(*stream.Get(2)); index != 2 {
//		t.Fatal("Unexpect Value stream IndexOf.", stream)
//	}
//	if index := stream.IndexOf(*stream.Get(5)); index != 5 {
//		t.Fatal("Unexpect Value stream IndexOf.", stream)
//	}
//	if v := stream.Max(func(arg Sample, index int) float64 { return float64(arg.Int) }); v.Int != 8 {
//		t.Fatal("Unexpect Value stream Max.", stream)
//	}
//	if v := stream.Min(func(arg Sample, index int) float64 { return float64(arg.Int) }); v.Int != 1 {
//		t.Fatal("Unexpect Value stream Max.", stream)
//	}
//	if v := cloned2.Peek(func(arg *Sample, index int) { arg.Int = -1 }); v.Get(0).Int != -1 {
//		t.Fatal("Unexpect Value stream Peek.", cloned2)
//	}
//	if v := cloned2.ForEach(func(arg Sample, index int) { arg.Int = 2 }); v.Get(0).Int == 2 {
//		t.Fatal("Unexpect Value stream ForEach.", cloned2)
//	}
//	count := 0
//	if stream.While(func(arg Sample, index int) bool { arg.Int = -1; count++; return index != 3 }); count != 4 && stream.Get(0).Int == -1 {
//		t.Fatal("Unexpect Value stream While.", stream)
//	}
//	if v := stream.Get(0); v.Str != "1" {
//		t.Fatal("Unexpect Value stream Get.", stream)
//	} else {
//		v.Str = ""
//		(*v).Str = ""
//		if stream.Get(0).Str != "1" {
//			t.Fatal("Unexpect Value stream Get.", stream)
//		}
//	}
//	if stream.Set(stream.Len()-1, Sample{Str: "last", Int: 0}); stream.Last().Str != "last" {
//		t.Fatal("Unexpect Value stream Set.", stream)
//	}
//	if tmp := stream.Copy().Skip(2).Limit(2); tmp.Get(0).Int != 3 || tmp.Len() != 2 {
//		t.Fatal("Unexpect Value stream Limit and Skip.", stream)
//	}
//	if !stream.Contains(Sample{Str: "1", Int: 1}) || stream.Contains(Sample{Str: "1", Int: 9}) {
//		t.Fatal("Unexpect Value stream Contains.", stream)
//	}
//	if stream.Add(*stream.Get(0)).Distinct(); stream.Last().Str != "last" {
//		t.Fatal("Unexpect Value stream Distinct.", stream)
//	}
//	sum := 0
//	if stream.SkippingEach(func(sample Sample, _ int) int {
//		sum += sample.Int
//		return 1
//	}); sum != 16 {
//		t.Fatal("Unexpect Value stream SkippingEach.", stream)
//	}
//
//	if created := CreateSampleStream().Add(Sample{Str: "1", Int: 2}).Clean(); created.IsPreset() {
//		t.Fatal("Unexpect Value stream CreateSampleStream.", created)
//	}
//	if created := GenerateSampleStream(SampleStream{}).Add(Sample{Str: "1", Int: 2}).Clean(); created.IsPreset() {
//		t.Fatal("Unexpect Value stream CreateSampleStream.", created)
//	}
//	if created := CreateSampleStream().Add(Sample{Str: "1", Int: 1}).Add(Sample{Str: "1", Int: 1}).Distinct(); created.Len() != 1 || created.Last().Int != 1 {
//		t.Fatal("Unexpect Value stream CreateSampleStream.", created)
//	}
//	if !CreateSampleStream().Equals(SampleStream{}) {
//		t.Fatal("Unexpect Value stream CreateSampleStream.")
//	}
//
//}
