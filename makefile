# mock 코드 생성
mock:
	@unset LANG LC_ALL LC_MESSAGES \
		&& go list -f '{{.Dir}}' ./... \
		| tail -n +2 \
		| grep -Ev 'vendor|cmd|fx' \
		| xargs -n1 mockery \
			--all \
			--case underscore \
			--keeptree \
			--disable-version-string \
			--with-expecter \
			--note "NOTE: run 'make mocks' to update this file and generate new ones." \
			--dir
.PHONY: mock

# tidy & vendor
install:
	go mod tidy && go mod vendor
.PHONY: install