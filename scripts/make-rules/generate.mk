# ==============================================================================
# 用来进行代码生成的 Makefile
#

.PHONY: gen.add-copyright
gen.add-copyright: ## 添加版权头信息.
	@addlicense -v -f $(ROOT_DIR)/scripts/boilerplate.txt $(ROOT_DIR) --skip-dirs=third_party,vendor,$(OUTPUT_DIR)

