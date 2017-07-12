package main

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

func main() {
	//   jobs := make([]job, 0, 3)
	//   jobs = append(jobs, job{id: 1, start: 3, end: 20})
	//   jobs = append(jobs, job{id: 2, start: 10, end: 29})
	//   jobs = append(jobs, job{id: 3, start: 20, end: 39})

	//   //fmt.Println(jobs)
	//   for kj, j := range jobs {
	//     for ko, o := range jobs {
	//       if kj != ko {
	//         fmt.Printf("job %d overlaps with job %d: %t\n", j.id, o.id, j.overlapsWith(o))
	//       }
	//     }
	//   }

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
	//fmt.Println("%v", servers)
	fmt.Printf("# of servers %d\n", len(servers))
}

type Server struct {
	id           int
	assignedJobs []job
}

func (s *Server) addJob(j job) {
	s.assignedJobs = append(s.assignedJobs, j)
}

func (s Server) isAvailable(j job) bool {
	for _, v := range s.assignedJobs {
		if v.overlapsWith(j) {
			return false
		}
	}

	return true
}

func getServers(jobs []job) []*Server {
	servers := []*Server{}

	// O(js) or O(n^2)

LOOP:
	for _, j := range jobs {
		for _, s := range servers {
			if s.isAvailable(j) {
				s.addJob(j)
				continue LOOP
			}
		}

		servers = append(servers, &Server{id: len(servers), assignedJobs: []job{j}})
	}

	return servers
}

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

a.overlapsWith(b) -> true
a.overlapsWith(c) -> false

*/
