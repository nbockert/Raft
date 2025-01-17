package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import (
	"os"
	"strconv"
)

// example to show how to declare the arguments
// and reply for an RPC.

type ExampleArgs struct {
	X int
}
type ExampleReply struct {
	Y int
}
type MapArg struct {
	Id         int
	Status     string
	Boolval    bool
	Complete   chan bool
	ReduceDone bool
}

type MapReply struct {
	Id         int
	Status     string
	Tasktype   string
	Title      string
	Nummaptask int
}

type Arg struct {
	Id int
}
type Bucket struct {
	TaskID int        // ID of the reduce task
	Files  []KeyValue // Files belonging to this bucket
}
type Reply struct {
	Id     int
	Type   string
	Finish bool
}
type IsDone struct {
	Finished bool
}

// Add your RPC definitions here.

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/824-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
