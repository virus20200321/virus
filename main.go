package main

import (
	"github.com/micmonay/keybd_event"
	"runtime"
	"time"
)

func main() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// For linux, it is very important wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	i := 0
	//set keys
	for true {
		time.Sleep(time.Second)
		kb.SetKeys(keybd_event.VK_A, keybd_event.VK_B)

		//set shif is pressed
		//kb.HasSHIFT(true)

		//launch
		err = kb.Launching()
		i += 2
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
		if i > 4 {
			kb.SetKeys(keybd_event.VK_KeypadEnter)
			//kb.HasCTRL()

			//launch
			err = kb.Launching()
			if err != nil {
				panic(err)
			}
			i = 0
		}
	}
}

//Ouput : AB
