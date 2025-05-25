package main

import (
	"fmt"
	"math"
)

type Interval struct {
	Start int
	End   int
}

func (interval *Interval) IsIntersect(other Interval) bool {
	return interval.Start <= other.End && interval.End >= other.Start
}

type Films map[string]Interval

func (films *Films) Print() {
	for name, interval := range *films {
		fmt.Printf("Film: %s [%d, %d]\n", name, interval.Start, interval.End)
	}
}

func (films *Films) EarliestEndFilm() (earlistEndFilmName string) {

	end := math.MaxInt
	for name, interval := range *films {
		if interval.End < end {
			end = interval.End
			earlistEndFilmName = name
		}
	}
	return
}

func (films *Films) RemoveIntersections(filmName string) {

	jInterval := (*films)[filmName]

	for name, interval := range *films {
		if jInterval.IsIntersect(interval) && filmName != name {
			delete(*films, name)
		}
	}
}

func (films *Films) Optimize() {

	for len(*films) != 0 {
		name := films.EarliestEndFilm()
		films.RemoveIntersections(name)
	}
}

func main() {

	films := Films{
		"Film1": Interval{1, 5},
		"Film2": Interval{2, 6},
		"Film3": Interval{4, 8},
		"Film4": Interval{7, 10},
	}

	films.Optimize()
	films.Print()
}
