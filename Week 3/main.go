package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	retry      = 10
	url        = "https://dummy.restapiexample.com/api/v1/employees"
	max_worker = 5
)

// Gọi API https://dummy.restapiexample.com/api/v1/employees và trả về 1 slice các struct có thông tin trong đấy
// Viết 1 worker pool dựa trên ví dụ ở https://gobyexample.com/worker-pools, nhưng job là struct về người gồm các thông tin ở câu trên, job trả về lương chia cho tuổi của mỗi người
type Employee struct {
	Name         string `json:"employee_name"`
	ProfileImage string `json:"profile_image"`
	Id           int    `json:"id"`
	Salary       int    `json:"employee_salary"`
	Age          int    `json:"employee_age"`
}
type Response struct {
	Succes string     `json:"success"`
	Data   []Employee `json:"data"`
}

func busyWait(n int) {
	for i := 0; i < n*1e6; i++ {
		// Busy-waiting by doing nothing
	}
}
func GetData(url string) []Employee {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Receive error:", err)
	} else {
		defer resp.Body.Close()
		scanner := bufio.NewScanner(resp.Body)
		scanner.Scan()
		var reponse Response
		if e := json.Unmarshal(scanner.Bytes(), &reponse); e != nil {
			fmt.Println("Err when unmarshal ", e)
		} else {
			return reponse.Data
		}
	}

	return []Employee{}
}
func worker(jobs chan Employee, wg *sync.WaitGroup) {
	for range jobs {
		time.Sleep(5 * time.Second)
		e := <-jobs
		fmt.Println("Salary / Age employee id ", e.Id, "= ", (e.Salary / e.Age))
		wg.Done()
	}
}
func main() {
	employee_list := GetData(url)
	var wg sync.WaitGroup
	jobs := make(chan Employee, len(employee_list))
	for i := 0; i < max_worker; i++ {
		go worker(jobs, &wg)
	}
	for i := 0; i < len(employee_list); i++ {
		jobs <- employee_list[i]
		wg.Add(1)
	}
	close(jobs)
	wg.Wait()
}
