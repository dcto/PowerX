# makefiles/build/app.mk

app-init: app-migrate app-seed app-run
app-init-db: app-migrate app-seed

app-migrate:
	go build -o $(POWERX_CTL_EXE_PATH) $(BUILD_CTL_PATH)
	$(POWERX_CTL_EXE_PATH) database migrate -f $(CONFIG_FILE)

app-seed:
	go build -o $(POWERX_CTL_EXE_PATH) $(BUILD_CTL_PATH)
	$(POWERX_CTL_EXE_PATH) database seed -f $(CONFIG_FILE)

app-run:
	go build -o $(POWERX_EXE_PATH) $(BUILD_EXE_PATH)
	$(POWERX_EXE_PATH) -f $(CONFIG_FILE)

