# Memo

# scss compiler won't create directory automacally
#
# scss and ts compilers create output files even an error occured
#
# tsc main.ts will compile a file and all its dependencies

# Variables ####################################################################

DEV_MODE ?=

CSS_EXT := .css
JS_EXT := .js
MAP_EXT := .map
PUG_EXT := .pug
SCSS_EXT := .scss
TMPL_EXT := .tmpl
TS_EXT := .ts

GO_CLI := go
NPM_CLI := npm
PUG_CLI := pug
SCSS_CLI := scss
TSC_CLI := tsc
TYPINGS_CLI := typings
UGLIFYJS_CLI := uglifyjs

NPM_INSTALL_FLAGS :=
PUG_FLAGS := --silent
SCSS_FLAGS :=
TSC_FLAGS :=
UGLIFYJS_FLAGS := --screw-ie8 --compress --mangle

ifdef DEV_MODE
PUG_FLAGS += --pretty
SCSS_FLAGS += --sourceMap=auto --style=expanded
TSC_FLAGS += --sourceMap
else
NPM_INSTALL_FLAGS += --production
SCSS_FLAGS += --sourceMap=none --style=compressed
TSC_FLAGS += --removeComments
endif

define LF


endef

## UI

UI_DIR := ui
UI_APP_DIR := $(UI_DIR)/app
UI_CSS_DIR := $(UI_DIR)/css
UI_JS_DIR := $(UI_DIR)/js
UI_NPM_DIR := $(UI_DIR)/node_modules
UI_SCSS_CACHE_DIR := $(UI_DIR)/.sass-cache
UI_TMPL_DIR := $(UI_DIR)/tmpl
UI_TYPINGS_DIR := $(UI_DIR)/typings

UI_NPM_CONFIG_FILE := $(UI_DIR)/package.json
UI_TYPINGS_CONFIG_FILE := $(UI_DIR)/typings.json
UI_TYPINGS_INDEX_FILE := $(UI_TYPINGS_DIR)/index.d.ts

UI_PUG_FILES := $(shell find $(UI_TMPL_DIR) -name "[!_]*$(PUG_EXT)" -print)
UI_SCSS_FILES := $(shell find $(UI_CSS_DIR) -name "[!_]*$(SCSS_EXT)" -print)
UI_TS_FILES := $(shell find $(UI_APP_DIR) $(UI_JS_DIR) -name "[!_]*$(TS_EXT)" -print)
UI_CSS_DIST_FILES := $(patsubst %$(SCSS_EXT),%$(CSS_EXT),$(UI_SCSS_FILES))
UI_JS_DIST_FILES := $(patsubst %$(TS_EXT),%$(JS_EXT),$(UI_TS_FILES))
UI_TMPL_DIST_FILES := $(patsubst %$(PUG_EXT),%$(TMPL_EXT),$(UI_PUG_FILES))
UI_DIST_FILES := $(UI_TMPL_DIST_FILES) $(UI_CSS_DIST_FILES) $(UI_JS_DIST_FILES) 

UI_PUG_FLAGS := $(PUG_FLAGS) --extension $(patsubst .%,%,$(TMPL_EXT))
UI_SCSS_FLAGS := $(SCSS_FLAGS) --cache-location $(UI_SCSS_CACHE_DIR)

## CMD

CMD_PULSE_NAME := pulse

CMD_DIR := cmd
CMD_NAMES := $(CMD_PULSE_NAME)
CMD_FILES := $(addprefix $(GOPATH)/bin/,$(CMD_NAMES))

## Deployment

DPL_DIR := wwwroot
DPL_PUBLIC_DIR := $(DPL_DIR)/static

DPL_UI_APP_DIR_LINK := $(DPL_PUBLIC_DIR)/app
DPL_UI_CSS_DIR_LINK := $(DPL_PUBLIC_DIR)/css
DPL_UI_JS_DIR_LINK := $(DPL_PUBLIC_DIR)/js
DPL_UI_NPM_DIR_LINK := $(DPL_PUBLIC_DIR)/node_modules
DPL_UI_TMPL_DIR_LINK := $(DPL_DIR)/tmpl
DPL_LINKS := \
	$(DPL_UI_APP_DIR_LINK) \
	$(DPL_UI_CSS_DIR_LINK) \
	$(DPL_UI_JS_DIR_LINK) \
	$(DPL_UI_NPM_DIR_LINK) \
	$(DPL_UI_TMPL_DIR_LINK)

# Rules ########################################################################

.PHONY: run all clean mostlyclean

run: all
	cd $(DPL_DIR) && $(GOPATH)/bin/$(CMD_PULSE_NAME)

all: ui_all cmd_all dpl_all

clean: ui_clean cmd_clean dpl_clean
	rm -Rf $(SCSS_CACHE_DIR)

mostlyclean: ui_mostlyclean cmd_mostlyclean dpl_mostlyclean

# UI

.PHONY: ui_all ui_clean ui_mostlyclean

.DELETE_ON_ERROR: $(UI_CSS_DIST_FILES) $(UI_JS_DIST_FILES)

ui_all: $(UI_DIST_FILES)

ui_clean: ui_mostlyclean
	rm -Rf $(UI_NPM_DIR) $(UI_TYPINGS_DIR) $(UI_SCSS_CACHE_DIR)

ui_mostlyclean:
	find $(UI_APP_DIR) $(UI_CSS_DIR) $(UI_JS_DIR) $(UI_TMPL_DIR) \( \
		-name "*$(CSS_EXT)" \
		-o -name "*$(JS_EXT)" \
		-o -name "*$(MAP_EXT)" \
		-o -name "*$(TMPL_EXT)" \
	\) -exec rm "{}" +

$(UI_TMPL_DIST_FILES): %$(TMPL_EXT) : %$(PUG_EXT)
	$(PUG_CLI) $(UI_PUG_FLAGS) --out $(dir $@) $<

# TODO: Compile from bootstrap file
$(UI_CSS_DIST_FILES): %$(CSS_EXT) : %$(SCSS_EXT)
	$(SCSS_CLI) $(UI_SCSS_FLAGS) $< $@

$(UI_JS_DIST_FILES): $(UI_TS_FILES) | $(UI_NPM_DIR) $(UI_TYPINGS_DIR)
	$(TSC_CLI) $(TSC_FLAGS) --project $(UI_DIR)
ifndef DEV_MODE
	$(foreach F,$(UI_JS_DIST_FILES),$(UGLIFYJS_CLI) $(UGLIFYJS_FLAGS) --output $F $F$(LF))
endif

$(UI_NPM_DIR): $(UI_NPM_CONFIG_FILE)
	cd $(UI_DIR) && $(NPM_CLI) install $(NPM_INSTALL_FLAGS)

$(UI_TYPINGS_DIR): $(TYPINGS_CONFIG_FILE)
	cd $(UI_DIR) && $(TYPINGS_CLI) install

## CMD

.PHONY: cmd_all cmd_clean cmd_mostlyclean

cmd_all: $(CMD_FILES)

cmd_clean: cmd_mostlyclean

cmd_mostlyclean: ; rm -Rf $(CMD_FILES)

$(CMD_FILES): $(GOPATH)/bin/% : $(CMD_DIR)/% handler
	$(GO_CLI) get ./$<
	$(GO_CLI) install ./$<

## Deployment

.PHONY: dpl_all dpl_clean dpl_mostlyclean

dpl_all: | $(DPL_DIR) $(DPL_LINKS)

dpl_clean: dpl_mostlyclean

dpl_mostlyclean: ; rm -Rf $(DPL_DIR)

$(DPL_UI_APP_DIR_LINK): | $(DPL_PUBLIC_DIR)
	ln -s $(abspath $(UI_APP_DIR)) $@

$(DPL_UI_CSS_DIR_LINK): | $(DPL_PUBLIC_DIR)
	ln -s $(abspath $(UI_CSS_DIR)) $@

$(DPL_UI_JS_DIR_LINK): | $(DPL_PUBLIC_DIR)
	ln -s $(abspath $(UI_JS_DIR)) $@

$(DPL_UI_NPM_DIR_LINK): | $(DPL_PUBLIC_DIR)
	ln -s $(abspath $(UI_NPM_DIR)) $@

$(DPL_UI_TMPL_DIR_LINK): | $(DPL_DIR)
	ln -s $(abspath $(UI_TMPL_DIR)) $@

$(DPL_DIR) $(DPL_PUBLIC_DIR):
	mkdir -p $@
