package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"path"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"strconv"
	"sync"
	"time"
)

func main() {

	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to `file`")
	memprofile := flag.String("memprofile", "", "write memory profile to `file`")
	traceprofile := flag.String("traceprofile", "", "write trace profile to `file`")

	output := flag.String("output", ".", "output folder")
	count := flag.Int("count", 1000, "number of files")
	size := flag.Int("size", 1000, "size of array")
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if *traceprofile != "" {
		f, err := os.Create(*traceprofile)
		if err != nil {
			log.Fatal("could not create trace profile: ", err)
		}
		if err := trace.Start(f); err != nil {
			log.Fatal("could not start trace profile: ", err)
		}
		defer trace.Stop()
	}

	var wg sync.WaitGroup

	for i := 0; i < *count; i++ {
		wg.Add(1)
		go func(i int) {
			filepath := path.Join(*output, "arr_"+strconv.Itoa(i)+".txt")
			f, err := os.Create(filepath)
			if err != nil {
				log.Fatal("cannot create %s:%v", filepath, err)
			}
			defer f.Close()
			generateAndWrite(*size, f)
			log.Printf("written %s", filepath)
			wg.Done()
		}(i)
	}
	wg.Wait()

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}

}

func generateAndWrite(size int, f *os.File) {
	rand := rand.New(rand.NewSource((int64)(time.Now().Nanosecond())))
	for i := 0; i < size; i++ {
		v := rand.Int()
		f.WriteString(strconv.Itoa(v) + "\n")
	}
}
