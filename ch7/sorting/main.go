package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (c customSort) Len() int           { return len(c.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func byArtist(x, y *Track) bool {
	return x.Artist < y.Artist
}

func byYear(x, y *Track) bool {
	return x.Year < y.Year
}

// type byArtist []*Track
// type byYear []*Track

// func (x byArtist) Len() int           { return len(x) }
// func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
// func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// func (x byYear) Len() int           { return len(x) }
// func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
// func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type IntSlice []int
type IntSlice2 []int

func (x IntSlice) Len() int           { return len(x) }
func (x IntSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x IntSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func isPalindrome(s sort.Interface) bool {
	n := s.Len()
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

func main() {
	// sort.Sort(byArtist(tracks))
	// printTracks(tracks)
	// sort.Sort(sort.Reverse(byArtist(tracks)))
	// printTracks(tracks)

	// sort.Sort(byYear(tracks))

	// printTracks(tracks)

	sort.Sort(customSort{tracks, byArtist})
	printTracks(tracks)
	sort.Sort(customSort{tracks, byYear})
	printTracks(tracks)

	fmt.Println(isPalindrome(IntSlice{1, 2, 2, 1}))
}
