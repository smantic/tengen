// Package tengen is used for creating command line apps from a config struct.
// It uses the names of the fields of the struct to create flags.
// It can be configured to use other sources of values
//
//  There are 3 levels of setting a config
//  1. Filling in the struct
//  2. Setting flags
//  3. Loading configs (e.g env var, remote)
//
// if a config is set in step 1 and found again in step 2 or 3 it will overrite the config value.
//
// The idea is that you can set the values of config manually by filling in the struct
// in code for the most basic set up of your application (default values are useful here)
// Then for more advance setups we can use env variables and flags to set our configs
// And further then that is a remote config setup using something like etcd
// for full manageability and observibility into what our configs are.
package tengen
