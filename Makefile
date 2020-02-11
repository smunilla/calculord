GO := go

COMMIT ?= $(if $(shell git status --porcelain --untracked-files=no),"${COMMIT_NO}-dirty","${COMMIT_NO}")
VERSION := ${shell cat ./VERSION}

calculord: ${SOURCES}
	${GO} build -ldflags "-X main.gitCommit=${COMMIT} -X main.version=${VERSION}"