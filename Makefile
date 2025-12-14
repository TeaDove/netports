VERSION ?= $(shell cat VERSION)

tag:
	git tag $(VERSION)
	git push origin --tags
