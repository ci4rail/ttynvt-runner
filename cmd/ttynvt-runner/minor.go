/*
Copyright © 2022 Ci4Rail GmbH <engineering@ci4rail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import "fmt"

const maxMinorNumbers = 256

var (
	minorMap [maxMinorNumbers]bool
)

func initMinorMap() {
	for i := 0; i < maxMinorNumbers; i++ {
		minorMap[i] = true
	}
}

func getMinor() (minor int, err error) {
	for minor, isFree := range minorMap {
		if isFree {
			minorMap[minor] = false
			return minor + 1, nil
		}
	}
	return 0, fmt.Errorf("no free minor numbers are left")
}

func releaseMinor(minor int) {
	if minor > 0 {
		minorMap[minor-1] = true
	}
}
