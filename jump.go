package plugin

import (
	"fmt"
	"log"

	"github.com/gdbu/jump/permissions"
	"github.com/mojura/mojura"

	"github.com/gdbu/jump"
	"github.com/gdbu/scribe"
	"github.com/hatchify/errors"
	"github.com/vroomy/vroomy"
)

const (
	// ErrInvalidSetUserArguments is returned when an invalid number of set user arguments are provided
	ErrInvalidSetUserArguments = errors.Error("invalid set user arguments, expecting no or one argument (redirectOnFail, optional)")
	// ErrInvalidCheckPermissionsArguments is returned when an invalid number of check permissions arguments are provided
	ErrInvalidCheckPermissionsArguments = errors.Error("invalid check permissions arguments, expecting two arguments (resource name and parameter key)")
	// ErrInvalidGrantPermissionsArguments is returned when an invalid number of grant permissions arguments are provided
	ErrInvalidGrantPermissionsArguments = errors.Error("invalid check permissions arguments, expecting three arguments (resource name, user actions, admin actions)")
)

const (
	permRWD = permissions.ActionRead | permissions.ActionWrite | permissions.ActionDelete
)

var p plugin

func init() {
	if err := vroomy.Register("jump", &p); err != nil {
		log.Fatalf("error loading jump plugin: %v", err)
	}
}

type plugin struct {
	vroomy.BasePlugin

	out  *scribe.Scribe
	jump *jump.Jump

	Opts mojura.Opts `vroomy:"mojura-opts"`
}

// Load will be called by Vroomy on initialization
func (p *plugin) Load(env vroomy.Environment) (err error) {
	p.out = scribe.New("Jump")
	if p.jump, err = jump.New(p.Opts); err != nil {
		err = fmt.Errorf("error initializing jump: %v", err)
		return
	}

	if p.Opts.IsMirror {
		return
	}

	if err = p.Seed(); err != nil {
		err = fmt.Errorf("error seeding users: %v", err)
		return
	}

	return
}

// Backend will return the plugin's backend
func (p *plugin) Backend() interface{} {
	return p.jump
}

// Close will close the Jump plugin and underlying Jump library
func (p *plugin) Close() error {
	return p.jump.Close()
}

func (p *plugin) Seed() (err error) {
	var apiKey string
	if _, err = p.jump.GetUser("00000000"); err == nil {
		return
	}

	// Set initial core permissions for users resource
	resourceKey := newResourceKey("users", "")
	if err = p.jump.SetPermission(resourceKey, "users", permissions.ActionNone, permRWD); err != nil {
		err = fmt.Errorf("error setting permissions for users group: %v", err)
		return
	}

	// Create a seed user
	if _, apiKey, err = p.jump.CreateUser("admin", "admin", "users", "admins"); err != nil {
		err = fmt.Errorf("error creating admin user: %v", err)
		return
	}

	if err = p.jump.Users().UpdateVerified("00000000", true); err != nil {
		err = fmt.Errorf("error verifying admin user: %v", err)
		return
	}

	p.out.Successf("Successfully created admin with api key of: %s", apiKey)
	return

}
