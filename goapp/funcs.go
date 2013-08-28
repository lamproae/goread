/*
 * Copyright (c) 2013 Matt Jibson <matt.jibson@gmail.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package goapp

import (
	"html/template"
	"reflect"
	"time"
)

// eq reports whether the first argument is equal to
// any of the remaining arguments.
func eq(args ...interface{}) bool {
	if len(args) == 0 {
		return false
	}
	x := args[0]
	switch x := x.(type) {
	case string, int, int64, byte, float32, float64:
		for _, y := range args[1:] {
			if x == y {
				return true
			}
		}
		return false
	}

	for _, y := range args[1:] {
		if reflect.DeepEqual(x, y) {
			return true
		}
	}
	return false
}

func routeUrl(name string, pairs ...string) string {
	u, err := router.Get(name).URL(pairs...)
	if err != nil {
		return err.Error()
	}
	return u.String()
}

func htmlattr(s string) template.HTMLAttr {
	return template.HTMLAttr(s)
}

func since(t time.Time) time.Duration {
	return time.Since(t)
}

func nanotime(i int64) time.Time {
	return time.Unix(0, i)
}

var funcs = template.FuncMap{
	"eq":       eq,
	"htmlattr": htmlattr,
	"nanotime": nanotime,
	"url":      routeUrl,
	"since":    since,
}
