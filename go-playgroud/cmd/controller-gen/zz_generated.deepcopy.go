//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package main

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in XdsResources) DeepCopyInto(out *XdsResources) {
	{
		in := &in
		*out = make(XdsResources, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XdsResources.
func (in XdsResources) DeepCopy() XdsResources {
	if in == nil {
		return nil
	}
	out := new(XdsResources)
	in.DeepCopyInto(out)
	return *out
}
