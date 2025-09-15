genproto:
	@ echo "start to gen the proto"
	@./scripts/genproto.sh
	@ echo "finish to gen the proto"

.PHONY: genproto


genopenapi:
	@ echo "start to gen the openapi"
	@./scripts/genopenapi.sh
	@ echo "finish to gen the openapi"

.PHONY: genopenapi

gen: genproto genopenapi

.PHONY: gen