package main

import (
	"Design-Pattern/BehaviorPattern/strategy/example2/shapes"
	"flag"
	"log"
	"os"
)

var output = flag.String("output", "text", "The output to use between "+
	"'console' and 'image' file")

func main() {
	flag.Parse()

	// *output 表示塞入地址為output的值 (string)
	activeStrategy, err := shapes.Factory(*output)
	if err != nil {
		log.Fatal(err)
	}

	switch *output {
	case shapes.TEXT_STRATEGY:
		// 直接取用activeStrategy (type為Output) 的 SetWriter
		// 因為DrawOutput有實作Output的SetWriter方法...
		activeStrategy.SetWriter(os.Stdout)
	case shapes.IMAGE_STRATEGY:
		w, err := os.Create("/tmp/image.jpg")
		if err != nil {
			log.Fatal("Error opening image")
		}
		defer w.Close()

		activeStrategy.SetWriter(w)
	}
	// 執行 Draw()的動作
	err = activeStrategy.Draw()
	if err != nil {
		log.Fatal(err)
	}
}
