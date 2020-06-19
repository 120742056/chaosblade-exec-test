/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"
	"fmt"
	"io/ioutil" //io 工具包
	"github.com/120742056/chaosblade-exec-test/exec/bin"
)

var (
	testStart, testStop 		bool
	info                            string
)

func main() {
	flag.BoolVar(&testStart, "start", false, "start burn cpu")
	flag.BoolVar(&testStop, "stop", false, "stop burn cpu")
	flag.StringVar(&info, "info", "", "the info to show")
	bin.ParseFlagAndInitLog()

	if testStart {
		startTestShow()
	} else if testStop {
		if success, errs := stopTestShow(); !success {
			bin.PrintErrAndExit(errs)
		}
	} else {
		bin.PrintErrAndExit("less --start or --stop flag")
	}
}

func check(e error) {
        if e != nil {
                panic(e)
        }
}

func wfile(s string) {
	d1 := []byte(s)
	err2 := ioutil.WriteFile("./output.txt", d1, 0666) //写入文件(字节数组)
	check(err2)
}

func startTestShow() {
	if info != "" {
		fmt.Println(info)
		wfile(info)
	} else {
		fmt.Println("no info")
		wfile("no info")
	}
}


// stopBurnCpu
func stopTestShow() (success bool, errs string) {
	fmt.Println("stop testshow")
	wfile("stop testshow")
	return true, errs
}
