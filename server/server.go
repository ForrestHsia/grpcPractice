package main

import (
	"context"
	"fmt"
	pb "grpcPractice/proto"
	"log"
	"net"
	sync "sync"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type ServerService struct{}

var UserData sync.Map

func main() {
	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatal("failed to listen:", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyprotoServiceServer(s, &ServerService{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatal("grpc.Serve Error: ", err)
		return
	}
}

func (s *ServerService) AddUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	var result string = "error"
	_, ok := UserData.Load(in.UserName)
	if !ok {
		UserData.Store(in.UserName, in.UserPwd)
		result = "ok"
	}
	return &pb.UserResponse{
		Result: result,
	}, nil
}

func (s *ServerService) LoginUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	var result string = "error"
	pwd, ok := UserData.Load(in.UserName)
	if pwd == in.UserPwd {
		if ok {
			result = "ok"
		}
	}
	return &pb.UserResponse{
		Result: result,
	}, nil
}

func (s *ServerService) UserList(ctx context.Context, in *pb.UserListRequest) (*pb.UserListResponse, error) {
	var users []string
	f := func(k, v interface{}) bool {
		var name string
		name = k.(string)
		fmt.Println(k, v)
		users = append(users, name)
		return true
	}
	UserData.Range(f)
	return &pb.UserListResponse{
		Result:   "ok",
		UserName: users,
	}, nil
}

func (s *ServerService) PingTest(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	result := "ping test OK"
	return &pb.PingResponse{
		ResultString: result,
	}, nil
}

// func main() {
// 	stringChannel := make(chan interface{}, 4)

// 	stringChannel <- strconv.Itoa(100)
// 	stringChannel <- 100
// 	stringChannel <- player{"cpbl", 250}
// 	stringChannel <- player{"mlb", 896}

// 	fmt.Println(<-stringChannel)
// 	fmt.Println(<-stringChannel)
// 	fmt.Println(<-stringChannel)
//     fmt.Println(<-stringChannel)
//     fmt.Println(<-stringChannel)
// 	fmt.Println(stringChannel)

// }

// func main() {

// 	testI := 1000
// 	start := time.Now().Nanosecond()
// 	wg := &sync.WaitGroup{}
// 	wg.Add(testI)
// 	for i := 1; i <= testI; i++ {
// 		go func() {
// 			test := cpblGetPlayer(wg)
// 			fmt.Println("main:", test)
// 		}()
// 	}
// 	wg.Wait()

// 	// for i := 1; i <= testI; i++ {
// 	// 	test := cpblGetPlayer(wg)
// 	// 	fmt.Println("main:", test)
// 	// }
// 	end := time.Now().Nanosecond()
// 	fmt.Println(end-start, "nano-s")
// 	fmt.Println("main finish")
// }

// func cpblGetPlayer(wg *sync.WaitGroup) int {
// 	defer func() {
// 		fmt.Println("branch finish")
// 		wg.Done()
// 	}()
// 	test := rand.Intn(10) + 1
// 	return test
// }

// func cpblGetPlayer(wg *sync.WaitGroup) int {
// 	test := rand.Intn(10) + 1
// 	fmt.Println("branch finish")
// 	return test
// }

// func main() {
//     // A goroutine-safe console printer.
//     logger := log.New(os.Stdout, "", 0)

//     // Sync between goroutines.
//     var wg sync.WaitGroup

//     // Add goroutine 1.
//     wg.Add(1)
//     go func() {
//         defer wg.Done()
//         logger.Println("Print from goroutine 1")
//     }()

//     // Add goroutine 2.
//     wg.Add(1)
//     go func() {
//         defer wg.Done()
//         logger.Println("Print from goroutine 2")
//     }()

//     // Add goroutine 3.
//     wg.Add(1)
//     go func() {
//         defer wg.Done()
//         logger.Println("Print from goroutine 3")
//     }()

//     logger.Println("Print from main")

//     // Wait all goroutines.
//     wg.Wait()
// }
