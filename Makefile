version=1.0
maintainer="3onyc <3onyc@x3tech.com>"

build: envtpl

envtpl:
	go build github.com/3onyc/envtpl

package-deb: build
	install -Dm 0755 envtpl pkg/usr/bin
	fpm \
		--architecture native \
		--maintainer $(maintainer) \
		--license MIT \
		--name envtpl \
		--version $(version) \
		-s dir \
		-t deb \
		-C pkg/ \
		.

clean:
	rm -rf pkg
