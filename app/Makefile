.PHONY: test cover serve clean

# 指定ディレクトリでテスト実行
TEST_DIR=.

test:
	cd $(TEST_DIR) && go test -v -coverprofile=coverage.out .

cover: test
	cd $(TEST_DIR) && go tool cover -html=coverage.out -o coverage.html

serve: cover
	@COVER_DIR=$$(mktemp -d) && \
	mv $(TEST_DIR)/coverage.html "$$COVER_DIR" && \
	echo "Coverage report available at: http://127.0.0.1:8000/coverage.html" && \
	(cd "$$COVER_DIR" && python3 -m http.server); \
	rm -rf "$$COVER_DIR"

clean:
	rm -f $(TEST_DIR)/coverage.out $(TEST_DIR)/coverage.html
