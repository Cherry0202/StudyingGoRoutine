package main

import(
	"time"
	"log"
)

func main() {
	log.Print("started")

//	チャネル
	finished := make(chan bool)

	// 配列宣言
	funcs := []func(){
		func() {
			// 1秒かかるコマンド
			log.Print("sleep1 started.")
			time.Sleep(1 * time.Second)
			log.Print("sleep1 finished.")
			finished <- true
		},
		func() {
			// 2秒かかるコマンド
			log.Print("sleep2 started.")
			time.Sleep(2 * time.Second)
			log.Print("sleep2 finished.")
			finished <- true
		},
		func() {
			// 3秒かかるコマンド
			log.Print("sleep3 started.")
			time.Sleep(3 * time.Second)
			log.Print("sleep3 finished.")
			finished <- true
		},
	}

	// 並行化する _,はindex省略
	for _, sleep := range funcs {
		go sleep()
	}

	// 終わるまで待つ
	for i := 0; i < len(funcs); i++ {
		<-finished
	}

	log.Print("finished")


}
