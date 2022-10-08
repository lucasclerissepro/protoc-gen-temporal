install:
	go install .


example: install
	pushd ./example > /dev/null \
	buf generate \
	popd > /dev/null
