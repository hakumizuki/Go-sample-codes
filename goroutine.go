package main
import(
	"fmt"
	// "time"
	"sync"
)

func goroutine(s string, wg *sync.WaitGroup) {
	defer wg.Done() // こういうときにdeferが使えるんや
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
}
func normal(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
}

// Main2 is in main()
func Main2() {
	// Goroutine 並列処理 goとつけるだけで並列になる goの並列処理よりも早く他の処理が完了してしまった場合はその時点で処理が終了する
	// WaitGroup
	var wg sync.WaitGroup
	wg.Add(1) // 待つDone()の回数
	go goroutine("Goroutine!", &wg) // 
	normal("normal")
	wg.Wait()
}