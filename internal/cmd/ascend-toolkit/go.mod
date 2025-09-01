module github.com/innoai-tech/ascend-toolkit/cmd/ascend-toolkit

go 1.25.0

replace (
	ascend-common v0.0.0 => ../../ascend-common
	github.com/innoai-tech/ascend-toolkit v0.0.0 => ../../..
)

require (
	ascend-common v0.0.0
	github.com/innoai-tech/ascend-toolkit v0.0.0
)

require (
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	k8s.io/apimachinery v0.25.3 // indirect
)
