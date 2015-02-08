# vultr
Vultr CLI and library

[![GoDoc](https://godoc.org/github.com/JamesClonk/vultr/lib?status.png)](https://godoc.org/github.com/JamesClonk/vultr/lib)
[![Build Status](https://travis-ci.org/JamesClonk/vultr.png?branch=master)](https://travis-ci.org/JamesClonk/vultr)

#### Screenshot

![Screenshot](https://github.com/JamesClonk/vultr/raw/master/screenshot.png "Screenshot")

#### TODO

* make custom json unmarshaller that maps everything to a string value (Vultr has a weird habit of giving switching types on some JSON fields without warning. (for example: pending charges. yesterday it was a string, today its a number))
* do that by first unmarshalling everything into a map[string]interface{}, then parse that into &struct
* switch all types in all structs to strings (ofc)

* add usage guide for command line tool
* add documentation on how to use library in other projects
