all: asm.s
	ls *.go | xargs -i ../blog_fmt.sh {}
	../blog_fmt.sh asm.s

asm.s: asm.go
	go build -gcflags=-S $< > $@ 2>&1
	go clean

