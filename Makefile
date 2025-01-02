.PHONY: build run watch

build:
	go build -o=/tmp/bin/main .
	
run: build
	/tmp/bin/main

watch:
	go run github.com/air-verse/air@latest \
		--build.cmd "make build" --build.bin "/tmp/bin/main" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"
