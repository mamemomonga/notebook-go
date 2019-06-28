#
# buildinfo $B$NKd$a9~$_(B
BUILDINFO_ARGS :=-X '$(BUILDINFO_IMPORT).Version=$(VERSION)' -X '$(BUILDINFO_IMPORT).Revision=$(REVISION)'

# $BI8=`%S%k%I(B(dynamic)
BUILDARGS := GO111MODULE=on \
	go build -mod vendor -a -ldflags="-s -w $(BUILDINFO_ARGS)"

# $B@EE*%S%k%I(B(static)
BUILDARGS_STATIC := GO111MODULE=on CGO_ENABLED=0 \
	go build -mod vendor -a -tags netgo -installsuffix netgo \
	-ldflags="-s -w $(BUILDINFO_ARGS) -extldflags '-static'"

# main()$B$N$"$k%G%#%l%/%H%j(B
MAINSRCDIR := $(SRCDIR)/sampleapp

# $B$9$Y$F$N%=!<%9(B
SRCS := $(shell find $(SRCDIR) -name '*.go')

# Docker$B%$%a!<%8(B
DOCKER_IMAGE=builder-$(NAME)

# $BI8=`%S%k%I(B(dynamic)
dynamic: $(BINDIR)/$(NAME)

# $B@EE*%S%k%I(B(static)
static: BUILDARGS=$(BUILDARGS_STATIC)
static: $(BINDIR)/$(NAME)

# $B<B9T%P%$%J%j(B
$(BINDIR)/$(NAME): $(SRCDIR)/vendor
	cd $(MAINSRCDIR) && $(BUILDARGS) -o $(abspath $(BINDIR)/$(NAME))

multiarch-build: $(SRCDIR)/vendor
	cd $(MAINSRCDIR) && $(BUILDARGS_STATIC) -o $(abspath $(BINDIR)/$(NAME)-$(GOOS)-$(GOARCH))
	@if [ "$(GOOS)" == "windows" ]; then mv $(BINDIR)/$(NAME)-$(GOOS)-$(GOARCH) $(BINDIR)/$(NAME)-$(GOOS)-$(GOARCH).exe; fi

# vendor$B%@%&%s%m!<%I(B
vendor: $(SRCDIR)/vendor
$(SRCDIR)/vendor:
	cd $(SRCDIR) && go mod vendor

# $BA]=|(B
clean:
	$(RM) -r $(BINDIR)

# Docker$B$G%S%k%I(B
docker:
	git rev-parse --short HEAD > revision
	docker build -t $(DOCKER_IMAGE) .
	docker run --rm $(DOCKER_IMAGE) tar cC /app bin | tar xvp

# Docker Image$B$r:o=|(B
rmi:
	docker rmi $(DOCKER_IMAGE)

.PHONY: dynamic static build clean multiarch multiarch-build vendor docker rmi

