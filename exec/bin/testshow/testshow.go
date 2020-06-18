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
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/chaosblade-io/chaosblade-spec-go/channel"
	"github.com/chaosblade-io/chaosblade-spec-go/util"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"

	"github.com/chaosblade-io/chaosblade-exec-test/exec/bin"
)

var (
	testStart, testStop 		bool
	info                            string
)

func main() {
	flag.BoolVar(&burnCpuStart, "start", false, "start burn cpu")
	flag.BoolVar(&burnCpuStop, "stop", false, "stop burn cpu")
	flag.StringVar(&info, "info", "", "the info to show")
	bin.ParseFlagAndInitLog()

	if burnCpuStart {
		startTestShow()
	} else if burnCpuStop {
		if success, errs := stopTestShowFunc(); !success {
			bin.PrintErrAndExit(errs)
		}
	} else {
		bin.PrintErrAndExit("less --start or --stop flag")
	}
}

func startTestShow() {
	ctx := context.Background()
	if info != "" {
		fmt.println(info)
	} else {
		fmt.println("no info")
	}
}


// stopBurnCpu
func stopBurnCpu() (success bool, errs string) {
	fmt.println("stop testshow")
	return true, errs
}
