# Whitespace

[![golang](https://img.shields.io/badge/Language-Go-green.svg)](https://golang.org/)
[![GoDoc](https://godoc.org/github.com/mwat56/whitespace?status.svg)](https://godoc.org/github.com/mwat56/whitespace)
[![Go Report](https://goreportcard.com/badge/github.com/mwat56/whitespace)](https://goreportcard.com/report/github.com/mwat56/whitespace)
[![Issues](https://img.shields.io/github/issues/mwat56/whitespace.svg)](https://github.com/mwat56/whitespace/issues?q=is%3Aopen+is%3Aissue)
[![Size](https://img.shields.io/github/repo-size/mwat56/whitespace.svg)](https://github.com/mwat56/whitespace/)
[![Tag](https://img.shields.io/github/tag/mwat56/whitespace.svg)](https://github.com/mwat56/whitespace/tags)
[![License](https://img.shields.io/github/mwat56/whitespace.svg)](https://github.com/mwat56/whitespace/blob/main/LICENSE)

- [Whitespace](#whitespace)
	- [Purpose](#purpose)
	- [Installation](#installation)
	- [Usage](#usage)
		- [Runtime](#runtime)
	- [Libraries](#libraries)
	- [Licence](#licence)

----

## Purpose

Whitespace (`TABulators`, `NewLines`, and `SPaces`) are characters in e.g. HTML pages which don't have syntactic significance but are just supposed to ease the reading by humans.
Whether you put one space between `</p>` and the following `<p>` doesn't make a difference as far as parsing and rendering the HTML page is concerned – it just increases the filesize and therefore the amount of bandwidth needed, the time used for transfer and interpretation, the amount of memory used in both the sending server and the receiving user/browser.
In the end one could say: the more whitespace there is in your HTML pages the more expensive it is for all parties involved.

When writing [Nele](https://github.com/mwat56/Nele/blob/main/README.md) and [Kaliber](https://github.com/mwat56/kaliber/blob/main/README.md) (both of which are essentially web-servers) I realised that I basically implemented the same code for removing superfluous whitespace before delivering the HTML pages to the remote user.
So I extracted that code and refactored it to become the simple middleware plugin you can see here.

## Installation

You can use `Go` to install this package for you:

	go get -u github.com/mwat56/whitespace

## Usage

There are two ways to use this library:

1. as a middleware plugin which then will be used automatically;
2. by calling the `Remove(…)` function from your code whenever it suits you.

In both cases you have to use this library by

	import "github.com/mwat56/whitespace"

To use it as middleware for your web-server you call

	whitespace.Wrap(aHandler http.Handler) http.Handler

where `aHandler` is the page-handler you implemented to handle (generate and send) your web-pages.
The function's return value can then be used to set up your `http.Server` instance.

If instead you want to use the library manually you can call

	whitespace.Remove(aPage []byte) []byte

where `aPage` is the HTML page you prepared and the function's return value is that very page with all superfluous whitespace removed.

### Runtime

You can de-/activate the removal behaviour at runtime by setting the boolean `UseRemoveWhitespace` flag.
If it is `true` (i.e. the default) both the `Remove()` and `Wrap()` functions work as expected (see above).
However, if you set that flag to `false` no whitespace removal takes place and the HTML pages produced by your server will be send to the remote user as they were generated (including the whitespace you and/or your tools put there).

## Libraries

No external libraries were used building `whitespace`.

## Licence

        Copyright © 2020, 2022 M.Watermann, 10247 Berlin, Germany
                        All rights reserved
                    EMail : <support@mwat.de>

> This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.
>
> This software is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
>
> You should have received a copy of the GNU General Public License along with this program. If not, see the [GNU General Public License](http://www.gnu.org/licenses/gpl.html) for details.

----
