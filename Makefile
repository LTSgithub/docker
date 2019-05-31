EXEC_FILE = manager
POLICY_EXEC_FILE = ./bin/$(EXEC_FILE)

.PHONY: all clean

all: clean  $(POLICY_EXEC_FILE)


$(POLICY_EXEC_FILE):
	@echo "all ..."
	go build -race -v $(EXEC_FILE)
	#go build -tags "debug" -v $(EXEC_FILE)
	@mkdir -p ./bin
	@mv $(EXEC_FILE) ./bin/
router:
	@echo "generate router"
	@cd ./src/manager && bee run -gendoc=true

clean:
	@echo "clean ..."
	rm -f $(POLICY_EXEC_FILE)


