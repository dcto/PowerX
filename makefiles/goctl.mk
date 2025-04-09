# makefiles/goctl.mk

goctl-powerx-apis:
	goctl api go -api ./api/powerx.api -dir . --style=goZero

