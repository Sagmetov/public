package main

import (
	"reflect"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func challengeBytes(fn1, fn2 interface{}, args ...interface{}) {
	st1 := z01.Monitor(fn1, args)
	st2 := z01.Monitor(fn2, args)
	if !reflect.DeepEqual(st1.Results, st2.Results) {
		z01.Fatalf("%s(%08b) == %08b instead of %08b\n",
			"SwapBits",
			args[0].(byte),
			st1.Results[0].(byte),
			st2.Results[0].(byte),
		)
	}
}

func main() {
	args := []byte{0x24, 0x14, 0x11, 0x22, 0xd2, 0x15, 0xff, 0x0, 0x35, 0x58, 0x43}

	for i := 0; i < 10; i++ {
		n := z01.RandIntBetween(0, 255)
		args = append(args, byte(n))
	}

	for _, v := range args {
		challengeBytes(student.SwapBits, solutions.SwapBits, v)
	}
}
