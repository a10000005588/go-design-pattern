package concurrency

var addCh chan bool = make(chan bool)

// 宣告一個channel會傳送 一個接收int的channel
var getCountCh chan chan int = make(chan chan int)
var quitCh chan bool = make(chan bool)

func init() {
	var count int
	go func(addCh <-chan bool, getCountCh <-chan chan int, quitCh <-chan bool) {
		for {
			select {
			case <-addCh:
				count++
			case ch := <-getCountCh:
				// 將收到的channel ，並把目前的count塞給ch這個channel
				ch <- count
			case <-quitCh:
				return
			}
		}
	}(addCh, getCountCh, quitCh)
}

type singleton struct{}

// 透過instance來操作所有只能被該instance呼叫的function type
var instance singleton

func GetInstance() *singleton {
	return &instance
}

func (s *singleton) AddOne() {
	// 對addCh 這個channel塞入 true的值
	addCh <- true
}

func (s *singleton) GetCount() int {
	// 宣告一個channel 來得到目前的加總的數目是多少
	resCh := make(chan int)
	// return完後會將resCh這個channel關閉
	defer close(resCh)
	// 將 resCh這個channel塞進到getCountCh這個channel
	getCountCh <- resCh
	return <-resCh
}

func (s *singleton) Stop() {
	quitCh <- true
	// 將宣告的channel關閉掉
	close(addCh)
	close(getCountCh)
	close(quitCh)
}
