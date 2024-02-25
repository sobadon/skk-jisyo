.PHONY: ci-upload
ci-upload:
	aws s3 cp --recursive ./contacts/ s3://skk-jisyo-chasoba-net/latest/contacts
	aws s3 cp --recursive ./skk/ s3://skk-jisyo-chasoba-net/latest/skk

.PHONY: syosyo
syosyo:
	go build -o syosyo ./cmd/syosyo

.PHONY: generate-jisyo
generate-jisyo:
	./scripts/generate-jisyo.sh
