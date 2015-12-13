// Copyright (C) 2015 Bartłomiej Biernacki
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"

	"github.com/cloudwan/gohan/cli"
	"github.com/cloudwan/gohan/extension"

	"github.com/stianeikeland/go-rpio"
)

//RazuberiModule shows example javascript module
type RazuberiModule struct {
}

func (example *RazuberiModule) SetHigh(gpio int) {

	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		return
	}
	defer rpio.Close()

	var pin = rpio.Pin(gpio)
	pin.Output()

	pin.High()
}

func (example *RazuberiModule) SetLow(gpio int) {

	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		return
	}
	defer rpio.Close()

	var pin = rpio.Pin(gpio)
	pin.Output()

	pin.Low()
}

func main() {
	razuberiModule := &RazuberiModule{}

	//Register go based module for javascript
	extension.RegisterModule("razuberi", razuberiModule)
	cli.Run("razuberi", "razuberi", "0.0.1")
}
