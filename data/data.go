// package data currently contains all the underlying
// data models, data interface, and data itself.
// TODO: persistent storage
package data

var subscriptions map[string]Subscription = make(map[string]Subscription)
