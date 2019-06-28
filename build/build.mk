# buildinfo $B$N(B $B%$%s%]!<%H%Q%9(B
# $B%W%m%8%'%/%H$N(BImport Path$B$O(Bgo.mod$B$+$i<hF@$9$k(B
BUILDINFO_IMPORT := $(shell cat go.mod | grep module | awk '{ print $2 }')/$(NAME)/buildinfo

# $B=PNO%G%#%l%/%H%j(B
BINDIR := bin

# $B%=!<%9$N$"$k%G%#%l%/%H%j(B
SRCDIR := .

# $B%P!<%8%g%s(B
VERSION   := v$(shell cat version)

# $B%j%S%8%g%s(B
# revision$B%U%!%$%k$,$"$l$P$=$3$+$i<hF@$9$k(B
REVISION  := $(shell if [ -e revision ]; then cat revision; else git rev-parse --short HEAD; fi)

# main$B$N$"$k%G%#%l%/%H%j(B
MAINSRCDIR := $(SRCDIR)/$(NAME)

# buildinfo $B$NKd$a9~$_(B
BUILDINFO_ARGS :=-X '$(BUILDINFO_IMPORT).Version=$(VERSION)' -X '$(BUILDINFO_IMPORT).Revision=$(REVISION)'

# $BI8=`%S%k%I(B(dynamic)
BUILDARGS := GO111MODULE=on \
	go build -mod vendor -a -ldflags="-s -w $(BUILDINFO_ARGS)"

# $B@EE*%S%k%I(B(static)
BUILDARGS_STATIC := GO111MODULE=on CGO_ENABLED=0 \
	go build -mod vendor -a -tags netgo -installsuffix netgo \
	-ldflags="-s -w $(BUILDINFO_ARGS) -extldflags '-static'"

# $B$9$Y$F$N%=!<%9(B
SRCS := $(shell find $(SRCDIR) -name '*.go')

# Docker$B%$%a!<%8(B
DOCKER_IMAGE=builder-$(NAME)

# ------------------------------------------------------------------

# $B%G%U%)%k%H$NF0:n(B
default: dynamic

# $BI8=`%S%k%I(B(dynamic)
dynamic: $(BINDIR)/$(NAME)

# $B@EE*%S%k%I(B(static)
static: BUILDARGS=$(BUILDARGS_STATIC)
static: $(BINDIR)/$(NAME)

# $B<B9T%P%$%J%j(B
$(BINDIR)/$(NAME): vendor
	cd $(MAINSRCDIR) && $(BUILDARGS) -o $(abspath $(BINDIR)/$(NAME))

multiarch-build: vendor
	cd $(MAINSRCDIR) && $(BUILDARGS_STATIC) -o $(abspath $(BINDIR)/$(NAME)-$(GOOS)-$(GOARCH))
	@if [ "$(GOOS)" == "windows" ]; then mv $(BINDIR)/$(NAME)-$(GOOS)-$(GOARCH) $(BINDIR)/$(NAME)-$(GOOS)-$(GOARCH).exe; fi

# vendor$B%@%&%s%m!<%I(B
vendor:
	cd $(SRCDIR) && go mod vendor

# $BA]=|(B
clean:
	$(RM) -r $(BINDIR)

# Docker$B$G%S%k%I(B
docker:
	git rev-parse --short HEAD > revision
	docker build -t $(DOCKER_IMAGE) .
	docker run --rm $(DOCKER_IMAGE) tar cC /g bin | tar xvp
	# docker run --rm -it $(DOCKER_IMAGE)

# Docker Image$B$r:o=|(B
rmi:
	docker rmi $(DOCKER_IMAGE)

# $B%^%k%A%"!<%-%F%/%A%c(B
# $BBP1~%j%9%H(B https://golang.org/doc/install/source#environment
multiarch:
	GOOS=darwin  GOARCH=amd64 $(MAKE) multiarch-build
	GOOS=windows GOARCH=amd64 $(MAKE) multiarch-build
	GOOS=linux   GOARCH=amd64 $(MAKE) multiarch-build
	GOOS=linux   GOARCH=arm64 $(MAKE) multiarch-build
	GOOS=linux   GOARCH=arm   $(MAKE) multiarch-build

.PHONY: dynamic static build clean multiarch multiarch-build vendor docker rmi

