// +build windows

/*
* Copyright 2014-2015 Jason Woods.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package main

import (
	"github.com/driskell/log-courier/Godeps/_workspace/src/github.com/op/go-logging"
	"os"
	"os/signal"
)

func (lc *LogCourier) registerSignals() {
	// Windows onyl supports os.Interrupt
	signal.Notify(lc.shutdown_chan, os.Interrupt)

	// No reload signal for Windows - implementation will have to wait
}

func (lc *LogCourier) configureLoggingPlatform(backends *[]logging.Backend) error {
	return nil
}