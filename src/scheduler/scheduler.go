package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Job struct {
	weight int
	length int
}

type JobSlice []Job

func (jobs JobSlice) SumWeightedCompletionTimes() int64 {
	sum := int64(0)
	time := 0

	for _, job := range jobs {
		sum += int64(job.weight * (job.length + time))
		time += job.length
	}

	return sum
}

type Strategy func(JobSlice) JobSlice

func Schedule(jobs JobSlice, strat Strategy) JobSlice {
	return strat(jobs)
}

type DifferenceJobSlice JobSlice

func (m DifferenceJobSlice) Len() int {
	return len(m)
}

func (m DifferenceJobSlice) Less(i, j int) bool {
	diff1, diff2 := m[i].weight-m[i].length, m[j].weight-m[j].length
	if diff1 == diff2 {
		return m[i].weight > m[j].weight
	}

	return diff1 > diff2
}

func (m DifferenceJobSlice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func DifferenceStrategy(jobs JobSlice) JobSlice {
	sort.Sort(DifferenceJobSlice(jobs))

	return jobs
}

type RatioJobSlice JobSlice

func (m RatioJobSlice) Len() int {
	return len(m)
}

func (m RatioJobSlice) Less(i, j int) bool {
	return float32(m[i].weight)/float32(m[i].length) >
		float32(m[j].weight)/float32(m[j].length)
}

func (m RatioJobSlice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func RatioStrategy(jobs JobSlice) JobSlice {
	sort.Sort(RatioJobSlice(jobs))

	return jobs
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	numJobsStr, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	var numJobs int
	fmt.Sscanf(numJobsStr, "%d", &numJobs) // Screw error checking.

	jobs := make(JobSlice, numJobs)
	for i := 0; i < numJobs; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		var weight, length int
		fmt.Sscanf(line, "%d %d", &weight, &length)

		jobs[i] = Job{weight, length}
	}

	differenceJobs := Schedule(jobs, DifferenceStrategy)
	fmt.Println(differenceJobs.SumWeightedCompletionTimes())

	ratioJobs := Schedule(jobs, RatioStrategy)
	fmt.Println(ratioJobs.SumWeightedCompletionTimes())
}
