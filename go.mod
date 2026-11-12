module github.com/roihe/column

go 1.26.0

require (
	github.com/kelindar/bitmap v1.5.7
	github.com/kelindar/column v0.4.1
	github.com/kelindar/intmap v1.5.0
	github.com/kelindar/iostream v1.4.0
	github.com/kelindar/simd v1.2.0
	github.com/kelindar/smutex v1.0.0
	github.com/klauspost/compress v1.20.4
	github.com/stretchr/testify v1.11.1
	github.com/tidwall/btree v1.8.2
	github.com/zeebo/xxh3 v1.1.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dustin/go-humanize v1.1.0
	github.com/kelindar/async v1.6.0
	github.com/kelindar/xxrand v1.0.2
	github.com/klauspost/cpuid/v2 v2.4.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.50.0 // indirect
	golang.org/x/time v0.16.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/kelindar/column => github.com/roihe/column v0.4.1-roihe

replace github.com/kelindar/column/commit => github.com/roihe/column/commit v0.4.1-roihe
