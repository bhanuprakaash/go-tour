package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var thumbnail = ThumbnailMock{}

// it will execute the thumbnails sequentially
func makeThumbnails1(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// it will run parallely but it will close the main so routines will be finished.
// we wont see the result
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f)
	}
}

// It will solve the above problem. wait until routines finishes.
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			log.Println(err)
			ch <- struct{}{}
		}(f)
	}
	for range filenames {
		<-ch
	}
}

// limitation with unbuffered channel
//
//	recevier is gone. so routine will be blocked
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}

	return nil
}


// solved the above problem using buffered channels
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbnail string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbnail, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbnail)
	}

	return thumbfiles, nil
}

// using wait groups It returns the number of bytes occupied by the files it creates.
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}
	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

func main() {
	inputFiles := make(chan string, 10)

	// 2. Feed it dummy data
	files := []string{
		"vacation.jpg", "profile.png", "dog.jpg", "cat.jpg",
		"document.pdf", "nature.bmp", "car.jpg", "food.png",
	}

	go func() {
		for _, f := range files {
			inputFiles <- f
		}
		close(inputFiles)
	}()

	fmt.Println("Starting thumbnail generation...")

	// makeThumbnails1(files)
	// makeThumbnails2(files)
	// makeThumbnails3(files)
	// fmt.Println(makeThumbnails4(files))
	// fmt.Println(makeThumbnails5(files))
_:
	makeThumbnails6(inputFiles)
}
