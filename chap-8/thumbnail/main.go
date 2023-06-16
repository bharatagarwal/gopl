package thumbnail

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func Image(src image.Image) image.Image {
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y
	width, height := 128, 128
	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(128 * aspect)
	} else {
		height = int(128 / aspect)
	}
	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcx := int(float64(x) * xscale)
			srcy := int(float64(y) * yscale)
			dst.Set(x, y, src.At(srcx, srcy))
		}
	}
	return dst
}
func ImageStream(w io.Writer, r io.Reader) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	dst := Image(src)
	return jpeg.Encode(w, dst, nil)
}
func ImageFile2(outfile, infile string) (err error) {
	in, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(outfile)
	if err != nil {
		return err
	}

	if err := ImageStream(out, in); err != nil {
		out.Close()
		return fmt.Errorf(
			"scaling %s to %s: %s",
			infile,
			outfile,
			err,
		)
	}
	return out.Close()
}
func ImageFile(infile string) (string, error) {
	ext := filepath.Ext(infile)
	outfile := strings.TrimSuffix(
		infile,
		ext,
	) + ".thumb" + ext
	return outfile, ImageFile2(outfile, infile)
}

// naive implementation
func makeThumbnails(filenames []string) {
	for _, file := range filenames {
		ImageFile(file) // ignore errors
	}
}

func makeThumbnails2(filenames []string) {
	// the program will return before
	// the goroutines are done with their job
	for _, file := range filenames {
		go ImageFile(file)
		// the goroutines will terminate when the
		// parent function returns
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})

	for _, file := range filenames {

		// important because closure maintains
		// the right value of `f`
		go func(f string) {
			ImageFile(f)

			// sending to channel
			ch <- struct{}{}
		}(file)
	}

	for range filenames {
		// wait to receive from all channels
		<-ch
	}
}

// handles errors
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, file := range filenames {
		go func(f string) {
			_, err := ImageFile(f)
			errors <- err // can be nil or error
		}(file)
	}

	for range filenames {
		err := <-errors
		if err != nil {
			return err
		}
	}

	return nil
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))

	for _, file := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = ImageFile(f)
			ch <- it
		}(file)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}

		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

/*
Note the asymmetry in the Add and Done methods. Add, which
increments the counter, must be called before the worker
goroutine starts, not within it; otherwise we would not be
sure that the Add happens before the “closer” goroutine
calls Wait. Also, Add takes a parameter, but Done does not;
it’s equivalent to Add(-1). We use defer to ensure that the
counter is decremented even in the error case. The structure
of the code above is a common and idiomatic pattern for
looping in parallel when we don’t know the number of
iterations.
*/

/*
The sizes channel carries each file size back to the main
goroutine, which receives them using a range loop and
computes the sum. Observe how we create a closer goroutine
that waits for the workers to finish before closing the
sizes channel. These two operations, wait and close, must be
concurrent with the loop over sizes. Consider the
alternatives: if the wait operation were placed in the main
goroutine before the loop, it would never end, and if placed
after the loop, it would be unreachable since with nothing
closing the channel, the loop would never terminate.
*/

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)

	// no of working goroutines
	var wg sync.WaitGroup

	for file := range filenames {
		wg.Add(1)

		go func(f string) {
			defer wg.Done()

			thumb, err := ImageFile(file)

			if err != nil {
				log.Println(err)
				return
			}

			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(file)

		go func() {
			wg.Wait()
			close(sizes)
		}()
	}

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}