package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

type Response struct {
	Data []interface{} `json:"data"`
}

type API struct{}

func (r API) getJSON(url string) (*Response, error) {

	body, err := makeRequest(url)
	checkError("Error in making request", err)

	// remove UTF-8 BOM
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

	s, err := parseJSON(body)
	return s, err
}

func makeRequest(url string) ([]byte, error) {

	res, err := http.Get(url)
	checkError("Error in HTTP GET request", err)

	body, err := ioutil.ReadAll(res.Body)
	checkError("Error reading response body", err)

	return []byte(body), err
}

func parseJSON(body []byte) (*Response, error) {
	var s = new(Response)
	err := json.Unmarshal(body, &s)
	checkError("Error parsing", err)
	return s, err
}

func savedataToCsv(kota string, data []interface{}, folder string) {
	if _, err := os.Stat(folder + "/" + kota + ".csv"); os.IsNotExist(err) {
		_, err := os.Create(folder + "/" + kota + ".csv")
		checkError("Cannot create file", err)
	}
	file, err := os.OpenFile(folder+"/"+kota+".csv", os.O_APPEND|os.O_WRONLY, 0600)
	checkError("Cannot open file", err)
	defer file.Close()

	writer := csv.NewWriter(file)

	var row []string
	for _, d := range data {

		m, _ := d.(map[string]interface{})

		for _, v := range m {
			if v == nil {
				row = append(row, "null")
			} else {
				row = append(row, v.(string))
			}
		}
		err = writer.Write(row)
		checkError("Cannot write to file", err)
	}

	defer writer.Flush()
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func (r API) getMuseum(kode string) ([]interface{}, error) {
	url := "http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET?kode_kab_kota="

	m, err := r.getJSON(url + kode)
	checkError("Retrieving museum list error", err)

	return m.Data, err
}

func getKota() []interface{} {
	api := new(API)

	// GET wilayah
	url := "http://jendela.data.kemdikbud.go.id/api/index.php/CWilayah/wilayahGET"
	wilayahList, err := api.getJSON(url)
	checkError("Retrieving wilayah list error", err)

	var kotaList []interface{}

	// GET kota
	for _, wilayah := range wilayahList.Data {
		w, _ := wilayah.(map[string]interface{})
		kode := strings.Trim(w["kode_wilayah"].(string), " ")

		url = "http://jendela.data.kemdikbud.go.id/api/index.php/CWilayah/wilayahGET?mst_kode_wilayah=" + kode
		kl, err := api.getJSON(url)
		checkError("Retrieving kota list error", err)

		for _, kota := range kl.Data {
			kotaList = append(kotaList, kota)
		}
	}
	return kotaList
}

func dothread(kota interface{}, folder string) {
	api := new(API)
	k, _ := kota.(map[string]interface{})

	kode := strings.Trim(k["kode_wilayah"].(string), " ")
	nama := strings.Trim(k["nama"].(string), " ")

	museum, err := api.getMuseum(kode)
	checkError("Retrieving kota list error", err)

	if len(museum) > 0 {
		savedataToCsv(nama, museum, folder)
	}
}

func main() {
	fmt.Println("Please Input the folder name : ")
	reader := bufio.NewReader(os.Stdin)
	read, _ := reader.ReadString('\r')
	read = strings.Replace(read, "\r", "", -1)
	err := os.Mkdir(read, 0700)
	checkError("Error creating directory", err)

	kotaList := getKota()

	totalJob := len(kotaList)
	totalWorker := 10

	idle := make(chan int, totalWorker)
	defer close(idle)

	// sleep all
	for i := 0; i < totalWorker; i++ {
		idle <- i + 1
	}

	// cek status thread
	var wg sync.WaitGroup
	wg.Add(totalJob)

	// save
	for _, kota := range kotaList {
		k, _ := kota.(map[string]interface{})
		nama := strings.Trim(k["nama"].(string), " ")

		w := <-idle
		go func(w int) {
			fmt.Printf("thread %d handle job %s \n", w, nama)
			defer func() {
				// sleep
				idle <- w
				wg.Done()
			}()

			dothread(kota, read)
		}(w)
	}

	wg.Wait()
}
