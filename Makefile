ARTIFACT = filehunt
SRC = main.go
OBJ = build/$(ARTIFACT)

buildg:
	@echo packing release......
	@goreleaser release --clean --skip=publish

