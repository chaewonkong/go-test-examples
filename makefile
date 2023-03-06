# mock 코드 생성
mock:
	@/bin/sh -c 'echo "${GREEN}[mocking을 시작합니다.]${NC}"'
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