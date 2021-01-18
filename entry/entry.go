// Copyright (c) 2020 rookie-ninja
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
package rk_entry

import (
	"github.com/rookie-ninja/rk-query"
	"go.uber.org/zap"
)

// An entry could be any kinds of services or pieces of codes which
// needs to be start/initialized while application starts
//
// A third party entry could be implemented and inject to rk-boot via rk-boot.yaml file
//
// How to create a new custom entry? Please see example/ for details
// Step 1:
// Construct your own entry YAML struct as needed
// example:
// ---
// myEntry:
//   enabled: true
//   key: value
//
// Step 2:
// Create a struct which implements Entry interface
//
// Step 3:
// Implements EntryRegFunc
//
// Step 4:
// Register your reg function in init() in order to register your entry while application starts
//
// How entry interact with rk-boot.Bootstrapper?
// 1: Entry will be created and registered into rk_ctx.GlobalAppCtx
// 2: Bootstrap will be called from Bootstrapper.Bootstrap() function
// 3: Application will wait for shutdown signal
// 4: Shutdown will be called from Bootstrapper.Shutdown() function

// New entry function which must be implemented
type EntryRegFunc func(string, *rk_query.EventFactory, *zap.Logger) map[string]Entry

// Entry interface which must be implemented for bootstrapper to bootstrap
type Entry interface {
	// bootstrap entry
	Bootstrap(rk_query.Event)

	// shutdown entry
	Shutdown(rk_query.Event)

	// get name of entry
	GetName() string

	// get type of entry
	GetType() string

	// print entry as string
	String() string
}
