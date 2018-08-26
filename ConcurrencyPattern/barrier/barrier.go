package barrier

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var timeoutMilliseconds int = 5000

// originally we should have
// one channel for handling error
// one channel for handling response
// But we can use struct to combile them.
// each Goroutine will send back a value of the barrierResp type.
type barrierResp struct {
	Err  error
	Resp string
}

func captureBarrierOutput(endpoints ...string) string {
	// a pipe allows us to connect io.Writer interface
	// to an io.Reader interface so that the Reader input
	// is the Writer output
	// Writer write something (output) --to--> (input) Reader
	reader, writer, _ := os.Pipe()

	// we define the os.Stdout as writer
	os.Stdout = writer
	// to capture "stdout" output (即捕捉有fmt.Println()等印在console在的字)
	// we will need a different gorouine that listen while we write to the console.
	out := make(chan string)
	// if we write, we don't capture
	// if we capture, we don't write
	// that's why we need concurrent structure
	// 因為while這個詞意思為 有A不能有B 有B不能有A的情況
	// 就必須要用concurrency 的架構來設定
	go func() {
		var buf bytes.Buffer
		// copy the reader input to the buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	barrier(endpoints...)

	writer.Close()
	temp := <-out
	// return the content captured from the console.
	return temp
}

// barrier is reponsible for block all the concurrency
// until it got all the response.
// endpoints is []string type
func barrier(endpoints ...string) {
	requestNumber := len(endpoints)
	// buffer channel with size requestNumber
	in := make(chan barrierResp, requestNumber)
	defer close(in)
	// create a slice : barrierResp which size is the number of request.
	responses := make([]barrierResp, requestNumber)
	// we lauch both request.
	for _, endpoint := range endpoints {
		// makeRequest takes a channel
		// and a endpoint.
		go makeRequest(in, endpoint)
	}

	var hasError bool
	// after we lauch two request
	// we wait for two response.
	// 這裡就是barrier的重要關鍵，透過for 搭配 <-channel 等待所有request的goroutine完成，在打印到console給stdout 捕捉
	for i := 0; i < requestNumber; i++ {
		// we block the execution waiting for data from channel [in]
		// got response from channel [in]
		// 用[<-in]來等候其他尚未完成的makeRequest concurrency action.
		resp := <-in
		// check response if there's error
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err)
			hasError = true
		}
		responses[i] = resp
	}

	if !hasError {
		for _, resp := range responses {
			// 將結果印在console，給reader去做捕捉
			fmt.Println(resp.Resp)
		}
	}
}

// accepts a channel to output 'barrierResp' values to and a URL to request.
// out type: (1) chan<-: only input (2) barrierResp: only accept barrierResp
// 由於在barrier中有[in]這個channel 需要拿出東西來, 故若要透過
// makeRequest取得跟Endpoint URL要到的資料，
// 那就得塞入一個只能"塞進去(input)"的out channel
// 然後barrier 再透過 [in] 將資料取出來
func makeRequest(out chan<- barrierResp, url string) {
	res := barrierResp{}
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMilliseconds) * time.Microsecond),
	}

	resp, err := client.Get(url)
	// if request failed, then append error to response.
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	// if response body is error, then append error to response.
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	// if no error, then throw the response to the [out] channel
	res.Resp = string(byt)
	out <- res
}
