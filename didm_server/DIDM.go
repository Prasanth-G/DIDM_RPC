package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"sync"
	"math/rand"
	"math"
	"strconv"
	"net"
	"net/http"
	"net/url"
	"os"
	"io/ioutil"
	"io"
	"encoding/json"
	"path"
	"github.com/Prasanth-G/split-downloader"
	"strings"

	"google.golang.org/grpc"
	pb "github.com/Prasanth-G/DIDM_RPC/didm"
)

const (
	port = ":50051"
	NO_OF_PARTS = 8
)

type bytesrange struct{
	order int
	start uint64
	end uint64
}

type job struct{
	workerid int
	work bytesrange
	worker string
}

type server struct{
	Link string
	NeighbourDevices map[int]string
	mutex *sync.Mutex
	jobsChannel chan job
	completed chan int
}


func (s *server)LinkHasRangeSupport() bool {
	client := &http.Client{}
	request, err := http.NewRequest("HEAD", s.Link, nil)
	if err != nil{
		return false
	}
	response, err := client.Do(request)
	if err != nil{
		return false
	}
	if response.Header.Get("Accept-Ranges") == "bytes"{
		request.Method = "GET"
		request.Header.Set("Range", "bytes="+strconv.FormatInt(0, 10)+"-"+strconv.FormatInt(1, 10))
		response, err = client.Do(request)
		switch{
		case err != nil:
			return false
		case response.ContentLength != 2:
			return false
		}
		return true
	}
	return false
}

func (s *server)AssignJobToWorker(work bytesrange){
	s.mutex.Lock()
	for len(s.NeighbourDevices) <= 0 {
		time.Sleep(time.Second)
	}
	randNumber := rand.Intn(len(s.NeighbourDevices))
	s.jobsChannel <- job{randNumber, work, s.NeighbourDevices[randNumber]}
	delete(s.NeighbourDevices, randNumber)
	s.mutex.Unlock()
}

func (s *server) DistributeDownload (ctx context.Context, in *pb.DistributedDownloadRequest) (*pb.DownloadResponse, error) {

	log.Printf("DDRequest Received : %v", in)
	s.Link = in.Link
	finalResponse := &pb.DownloadResponse{
		Data : nil,
		RequestReceived : &pb.DownloadResponse_DDRequest{DDRequest : in},
	}

	//Assign filename if not given
	if len(strings.TrimSpace(in.Saveas)) == 0 {
		parsedurl, _ := url.Parse(s.Link)
		str := strings.Split(parsedurl.Path, "/")

		if len(strings.TrimSpace(str[len(str)-1])) > 0 {
			in.Saveas = str[len(str)-1]
		} else {
			number := 1
			for _, err := os.Stat(fmt.Sprintf("DIDM_Unnamed_File_%d", number)); ! os.IsNotExist(err) ; number++ {}
			in.Saveas = fmt.Sprintf("DIDM_Unnamed_File_%d", number)
		}
	}

	//Check Whether the server supports downloading the files part by part
	splittable := s.LinkHasRangeSupport()
	if ! splittable {
		final, err := os.OpenFile(path.Join(in.Saveto, in.Saveas), os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0600)
		if err != nil {
			log.Printf("Error opening File: %v", err)
			return finalResponse, err
		}
		defer final.Close()
		response, err := http.Get(s.Link)
		if err != nil{
			log.Printf("Error while downloading: %v", err)
			return finalResponse, err
		}
		defer response.Body.Close()
		io.Copy(final, response.Body)
		return finalResponse, nil
	}
	if len(in.PeerIPAddr) == 1 {
		i := splitdownload.SDR{NO_OF_PARTS, in.Link}
		i.CompleteDownload(in.Saveas, in.Saveto)
	}


	//Distribute Download
	s.NeighbourDevices = make(map[int]string)
	s.jobsChannel = make(chan job, 32)
	s.completed = make(chan int, 32)
	s.mutex = &sync.Mutex{}

	rand.Seed(time.Now().UTC().UnixNano())
	response, err := http.Head(in.Link)
	if err != nil{
		log.Printf("Error sending HEAD Request: %v", err)
		return finalResponse, err
	}

	//Populate neighbours (necessary ???)
	for index := 0; index < len(in.PeerIPAddr); index++ {
		s.NeighbourDevices[index] = in.PeerIPAddr[index]
	}
	
	//divide the file into 8 parts; update - split file depends on DOWNLOAD SPEED, and NUMBER OF NEIGHBOUR DEVICE
	chunkSize := uint64(math.Ceil(float64(response.ContentLength) / float64(NO_OF_PARTS)))
	NoOfWorkers := len(s.NeighbourDevices)

	for w := 0; w < NoOfWorkers; w++ {
		go s.worker()
	}

	var jobs = make([]bytesrange, NO_OF_PARTS)
	var start uint64
	index := 0
	uContentLength := uint64(response.ContentLength)
	for start = 0; start < uContentLength; start += chunkSize{
		end := start + chunkSize - 1
		if end > uContentLength + 1{
			end = uContentLength
		}
		jobs[index] = bytesrange{index, start, end}
		index++
	}

	for _, j := range jobs{
		s.AssignJobToWorker(j)
	}
	for i := 0; i < NO_OF_PARTS; i++{
		log.Println(<-s.completed)
	}
	defer close(s.jobsChannel)

	final, err := os.OpenFile(path.Join(in.Saveto, in.Saveas), os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0600)
	if err != nil{
		log.Printf("Error opening File: %v", err)
		return finalResponse, err
	}

	defer final.Close()
	for i:= 0;i < NO_OF_PARTS; i++{
		file := "Part_"+strconv.Itoa(i)
		out, _ := ioutil.ReadFile(file)
		defer os.Remove(file)
		final.Write(out)
	}

	fmt.Printf("Distributed Download Called : %v", in)
	return finalResponse, nil
}

func (s *server) Download(ctx context.Context, in *pb.DownloadRequest) (*pb.DownloadResponse, error) {

	log.Printf("DRequest Received : %v", in)
	resp := &pb.DownloadResponse{
		Data : nil,
		RequestReceived : &pb.DownloadResponse_DRequest{DRequest : in},
	}

	i := splitdownload.SDR{NO_OF_PARTS, in.Link}
	i.PartialDownload([2]uint64{in.Start, in.End}, "file", "")
	out, err := ioutil.ReadFile("file")
	if err != nil {
		log.Printf("Error Reading temp File: %v", err)
		return resp, err
	}
	resp.Data = out
	return resp, nil
}

func (s *server) worker() {

	for j := range s.jobsChannel {
		log.Println("worker", j.workerid, "started  job", j.work)

		//Make call to `Download` function in other systems
		conn, err := grpc.Dial(j.worker, grpc.WithInsecure())
		if err != nil {
			log.Printf("Can't connect to remote machine: %v", err)
			s.NeighbourDevices[j.workerid] = j.worker
			s.AssignJobToWorker(j.work)
			continue
		}
		defer conn.Close()
		c := pb.NewDownloadServiceClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		resp, err := c.Download(ctx, &pb.DownloadRequest{
			Link	: s.Link,
			Start	: j.work.start,
			End		: j.work.end,
		})
		if err != nil {
			log.Printf("Can't Call the Remote Method: %v", err)
			s.NeighbourDevices[j.workerid] = j.worker
			s.AssignJobToWorker(j.work)
			continue
		}

		final, err := os.OpenFile("Part_"+strconv.Itoa(j.work.order), os.O_CREATE | os.O_TRUNC | os.O_WRONLY, 0600)
		if err != nil{
			log.Printf("Can't Open temp File: %v", err)
			s.NeighbourDevices[j.workerid] = j.worker
			s.AssignJobToWorker(j.work)
			continue
		}
		final.Write(resp.Data)
		final.Close()

		//add device back to queue
		s.NeighbourDevices[j.workerid] = j.worker
		s.completed <- j.work.order
	}
}

func Handler(writer http.ResponseWriter, request *http.Request){
	switch request.Method{
	case "POST":
		var dict map[string]string
		data, _ := ioutil.ReadAll(request.Body)
		json.Unmarshal(data, &dict)
		log.Printf("Post Req data: %v", dict)
		if _, ok := dict["Url"]; !ok{
			writer.WriteHeader(400)
			writer.Write([]byte("Request Body must contain an Url"))
		}
		if _, ok := dict["Saveto"]; !ok{
			dict["Saveto"] = ""
		}
		if _, ok := dict["Saveas"]; !ok{
			dict["Saveas"] = ""
		}
		
		s := server{}
		in := &pb.DistributedDownloadRequest{
			Link : dict["Url"],
			Saveto : dict["Saveto"],
			Saveas : dict["Saveas"],
		}
		var arr []string
		json.Unmarshal([]byte(dict["PeerIPAddr"]), &arr)
		in.PeerIPAddr = arr

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		s.DistributeDownload(ctx, in)
		writer.WriteHeader(202)
		writer.Write([]byte("File Downloaded"))
	}
}

func main(){

	http.HandleFunc("/", Handler)
	go http.ListenAndServe(":8000", nil)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDownloadServiceServer(s, &server{})
	log.Println("Registering Server ... ")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	
	//"https://maggiemcneill.files.wordpress.com/2012/04/the-complete-sherlock-holmes.pdf"
}