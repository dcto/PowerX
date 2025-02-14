CURRENT_DIR := $(shell pwd)
CONFIG_FILE := $(CURRENT_DIR)/etc/powerx.yaml

# 设定需要编译的go文件目录
BUILD_EXE_PATH := $(CURRENT_DIR)/cmd/server/powerx.go
BUILD_CTL_PATH := $(CURRENT_DIR)/cmd/ctl/powerxctl.go


PATH_BUILD := $(CURRENT_DIR)/.build

# 将编译好的执行文件，放入根目录下
POWERX_EXE_PATH:=$(CURRENT_DIR)/powerx
POWERX_CTL_EXE_PATH:=$(CURRENT_DIR)/powerxctl

export CURRENT_DIR
export CONFIG_FILE
export BUILD_EXE_PATH
export BUILD_CTL_PATH
export PATH_BUILD
export POWERX_EXE_PATH
export POWERX_CTL_EXE_PATH

include makefiles/goctl.mk
include makefiles/app.mk
include makefiles/build/linux.mk
include makefiles/build/macos.mk
include makefiles/build/windows.mk
