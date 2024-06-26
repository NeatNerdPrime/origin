/*
Copyright 2023 The Kubernetes Authors.

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

package fnutil

func Map[T any, R any](f func(T) R, xs []T) []R {
	rv := make([]R, len(xs))
	for i, x := range xs {
		rv[i] = f(x)
	}
	return rv
}

func RemoveIf[T any](f func(T) bool, xs []T) []T {
	var rv []T
	for _, x := range xs {
		if !f(x) {
			rv = append(rv, x)
		}
	}
	return rv
}

func IsAll[T any](f func(T) bool, xs []T) bool {
	for _, x := range xs {
		if !f(x) {
			return false
		}
	}
	return true
}

func IndexSet[T comparable](xs []T) map[T]bool {
	rv := make(map[T]bool, len(xs))
	for _, x := range xs {
		rv[x] = true
	}
	return rv
}

func Intersection[T comparable](xs, ys []T) []T {
	ysSet := IndexSet(ys)
	var rv []T
	for _, x := range xs {
		if ysSet[x] {
			rv = append(rv, x)
		}
	}
	return rv
}
