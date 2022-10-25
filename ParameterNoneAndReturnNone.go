// no argument
package Delegate

import "reflect"

// Delegate with no Parameter and no Return Value
type ParameterNoneAndReturnNone struct {
	array []func()
}

// Add Function To This Delegate
func (P *ParameterNoneAndReturnNone) Add(funcname func()) {
	P.array = append(P.array, funcname)
}

// Remove Function From This Delegate
func (d *ParameterNoneAndReturnNone) Remove(funcname func()) {
	var temp []func()
	for _, v := range d.array {
		if reflect.ValueOf(funcname).Pointer() == reflect.ValueOf(v).Pointer() {
			continue
		} else {
			temp = append(temp, v)
		}
	}
	d.array = temp
}

// Run This Delegate
func (d *ParameterNoneAndReturnNone) Run() {
	for _, v := range d.array {
		v()
	}
}
