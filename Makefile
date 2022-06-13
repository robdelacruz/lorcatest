# Usage:
# 'make dep' and 'make webtools' to install dependencies.
# 'make clean' to clear all work files
# 'make' to build bin, css and js
# 'make serve' to start dev webserver

NODE_VER = 14

JSFILES = index.js Index.svelte

all: t www/style.css www/index.html www/bundle.js

dep:
	go env -w GO111MODULE=auto
	go get github.com/zserge/lorca

webtools:
	npm install --save-dev tailwindcss
	npm install --save-dev postcss-cli
	npm install --save-dev cssnano
	npm install --save-dev svelte
	npm install --save-dev rollup
	npm install --save-dev rollup-plugin-svelte
	npm install --save-dev @rollup/plugin-node-resolve

www/index.html: index.html
	mkdir -p www
	cp index.html www/index.html

www/style.css: twsrc.css
	#npx tailwind build twsrc.css -o twsrc.o 1>/dev/null
	#npx postcss twsrc.o > www/style.css
	mkdir -p www
	npx tailwind -i twsrc.css -o www/style.css 1>/dev/null

www/bundle.js: $(JSFILES)
	mkdir -p www
	npx rollup -c

t: t.go
	go build -o t t.go

clean:
	rm -rf t www

serve:
	python -m SimpleHTTPServer

