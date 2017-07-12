package main

/*
 0   4   8   12  16  20  24  28  32  36  40  44  48  52  56  60  64
-+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+--- > t
4  [============)
1   [================)
6       [==================)
2          [==================)
3                    [==================)
7                                [==================)
5                                  [=============)
8                                          [==================)
-+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+--- > t
 0   4   8   12  16  20  24  28  32  36  40  44  48  52  56  60  64
*/

import "fmt"

type job struct {
	id    int
	start int
	end   int
}

func (j job) String() string {
	return fmt.Sprintf("%d %d %d", j.id, j.start, j.end)
}

func (j job) overlapsWith(o job) bool {
	if (j.start >= o.start && j.start < o.end) || (o.start >= j.start && o.start < j.end) {
		return true
	}

	return false
}

type server struct {
	id           int
	assignedJobs []job
}

func (s server) String() string {
	return fmt.Sprintf("%d %s", s.id, s.assignedJobs)
}

func (s *server) addJob(j job) {
	s.assignedJobs = append(s.assignedJobs, j)
}

func (s server) isAvailable(j job) bool {
	for _, v := range s.assignedJobs {
		if v.overlapsWith(j) {
			return false
		}
	}

	return true
}

func main() {
	// assume jobs are sorted via ascending start time
	jobs := []job{
		job{4, 2, 15},
		job{1, 3, 20},
		job{6, 7, 26},
		job{2, 10, 29},
		job{3, 20, 39},
		job{7, 32, 51},
		job{5, 34, 48},
		job{8, 42, 61},
	}

	servers := getServers(jobs)
	fmt.Printf("# of servers %d\n", len(servers))

	for _, s := range servers {
		fmt.Println(s)
	}
}

func getServers(jobs []job) []*server {
	servers := []*server{}

LOOP:
	// solution is O(jobs*servers) or O(n^2)
	for _, j := range jobs {
		for _, s := range servers {
			if s.isAvailable(j) {
				s.addJob(j)
				continue LOOP
			}
		}

		servers = append(servers,
			&server{id: len(servers) + 1, assignedJobs: []job{j}})
	}

	return servers
}
