package main

import (
	"context"
	"log"

	"github.com/gdbu/jump/permissions"

	"github.com/gdbu/jump"
	"github.com/hatchify/errors"

	"github.com/hatchify/scribe"
)

const (
	// ErrResourceIDIsEmpty is returned when resource id is expected but not found withing a permissions hook
	ErrResourceIDIsEmpty = errors.Error("resourceID is empty")
	// ErrInvalidSetUserArguments is returned when an invalid number of set user arguments are provided
	ErrInvalidSetUserArguments = errors.Error("invalid set user arguments, expecting no or one argument (redirectOnFail, optional)")
	// ErrInvalidCheckPermissionsArguments is returned when an invalid number of check permissions arguments are provided
	ErrInvalidCheckPermissionsArguments = errors.Error("invalid check permissions arguments, expecting two arguments (resource name and parameter key)")
	// ErrInvalidGrantPermissionsArguments is returned when an invalid number of grant permissions arguments are provided
	ErrInvalidGrantPermissionsArguments = errors.Error("invalid check permissions arguments, expecting three arguments (resource name, user actions, admin actions)")
	// ErrAlreadyLoggedOut is returned when a logout is attempted for a user whom has already logged out of the system.
	ErrAlreadyLoggedOut = errors.Error("already logged out")
)

const (
	permR   = permissions.ActionRead
	permRW  = permissions.ActionRead | permissions.ActionWrite
	permRWD = permissions.ActionRead | permissions.ActionWrite | permissions.ActionDelete
	permWD  = permissions.ActionWrite | permissions.ActionDelete
)

var p plugin

// Init will be called by vroomy on initialization
func Init(env map[string]string) (err error) {
	p.out = scribe.New("Auth")

	if p.jump, err = jump.New(env["dataDir"]); err != nil {
		log.Fatalf("error initializing jump: %v", err)
	}

	// TODO: Move this to seed
	if err = p.seed(); err != nil {
		log.Fatalf("error seeding users: %v", err)
	}

	return
}

type plugin struct {
	out  *scribe.Scribe
	jump *jump.Jump
}

func (p *plugin) seed() (err error) {
	var apiKey string
	if _, err = p.jump.GetUser("00000000"); err == nil {
		return
	}

	// Set initial core permissions for users resource
	resourceKey := newResourceKey("users", "")
	// TODO: use httpserve.Context here (needs PR's on httpserver Context type, out of scope currently)
	if err = p.jump.SetPermission(context.Background(), resourceKey, "users", permissions.ActionNone, permRWD); err != nil {
		return
	}

	// Create a seed user
	// TODO: use httpserve.Context here (needs PR's on httpserver Context type, out of scope currently)
	if _, apiKey, err = p.jump.CreateUser("admin", "admin", "users", "admins"); err != nil {
		return
	}

	p.out.Successf("Successfully created admin with api key of: %s", apiKey)
	return

}

// Jump will return the plugin's backend
func Jump() *jump.Jump {
	return p.jump
}

// Backend will return the plugin's backend
func Backend() interface{} {
	return p.jump
}

// Close will close the Jump plugin and underlying Jump library
func Close() error {
	return p.jump.Close()
}
