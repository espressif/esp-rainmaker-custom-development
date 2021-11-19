SILENT-$(VERBOSE) := @

# In the short-term, we will support packages as well as executables

# Service Name
SERVICE-NAME := rainmaker

# Versioning
MAJOR_VERSION:=1
MINOR_VERSION:=0
REVISION:=0
BUILD_DATE ?= $(shell date +%Y-%m-%dT%H:%M)
BUILD_VERSION ?= $(shell git log --pretty=format:"%h" -1)

help:
	@echo ""
	#command to build package
	@echo "make <PACKAGE_NAME> "
	@echo ""
	#command to build all packages
	@echo "make all "
	@echo ""
	#command to deploy package
	@echo "make <PACKAGE_NAME>-deploy S3-BUCKET=<YOUR_BUCKET_NAME> [STAGE-NAME=<YOUR_STAGE_NAME>]  "
	@echo ""
	#command to build and deploy all packages
	@echo "make deploy S3-BUCKET=<YOUR_BUCKET_NAME> [STAGE-NAME=<YOUR_STAGE_NAME>]  "
	@echo ""
	#command to remove deployed package
	@echo "make <PACKAGE_NAME>-remove "
	@echo ""

# executables - the previous structure
EXE_SRCS   := $(wildcard src/handlers/controllers/*)
EXE_TARGETS := $(patsubst src/handlers/controllers/%,%,$(EXE_SRCS))
STAGE-NAME := dev
VERSION := 1.0.0
REGION := us-east-1

# all packages
ALL_PKGS            := espcustombaseapi espcustombase espcustomservice

# all packages
DEPLOY_PKGS           := espcustombaseapi espcustombase espcustomservice

# Add the App version, build number, build date and customer deployment using LDFlags
LDFLAGS += "-s -w -X main.appName=$(SERVICE-NAME) -X main.appVersion=$(MAJOR_VERSION).$(MINOR_VERSION).$(REVISION)-$(BUILD_VERSION) -X main.appBuildDate=$(BUILD_DATE) \
             -X main.customerDeployment=$(CUSTOMER_DEPLOYMENT) -X main.rainmakerTrial=$(RAINMAKER_TRIAL)"

all: $(ALL_PKGS)

deploy: $(ALL_PKGS)

dep_ensure:
	#dep ensure

define GetValueFromConfig
$(shell node -p "require('./$(STAGE-NAME)-config.json').$(1)")
endef

define pkg_targets
-include src/handlers/$(1)/build.mk

$(1): $$(wildcard src/handlers/$(1)/executables/*)

$(1)-deploy:
	@[ "$(S3-BUCKET)" ] || ( echo ">> Please enter valid s3 bucket name or create a bucket using command: aws s3 mb <YOUR_BUCKET_NAME>"; exit 1)
ifneq ($(filter $(1),$(DEPLOY_PKGS)),)
	$$(SILENT-)sam package --template-file src/handlers/*/$(1).yml --output-template-file $(1)_package.yml --s3-bucket "$(S3-BUCKET)" --region "$(REGION)"

	$$(SILENT-)sam deploy --template-file $(1)_package.yml --stack-name $(1) --capabilities CAPABILITY_NAMED_IAM --no-fail-on-empty-changeset --parameter-overrides "StageName=$(STAGE-NAME)"  \
	$(2) --region "$(REGION)"
endif

deploy: $(1)-deploy

$(1)-remove:
	$$(SILENT-)aws cloudformation delete-stack --stack-name $(1)

remove: $(1)-remove

$$(wildcard src/handlers/$(1)/executables/*): dep_ensure
	@echo "[go] $(1) => $$(notdir $$@)"
	$$(SILENT-)GOOS=linux go build -ldflags $(LDFLAGS)  -o bin/handlers/$(1)/$$(notdir $$@) $$@/*.go
endef

config := "LogLevel=$(call GetValueFromConfig,LogLevel)" \
	"BuildVersion=$(BUILD_VERSION)"

$(foreach pkg,$(ALL_PKGS),$(eval $(call pkg_targets,$(pkg),$(config))))
