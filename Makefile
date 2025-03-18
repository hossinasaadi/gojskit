BUILD_DIR=build
IOS_ARTIFACT=$(BUILD_DIR)/GoJsKit.xcframework

ANDROID_ARTIFACT=$(BUILD_DIR)/GoJsKit.aar
LDFLAGS="-s -w"
IMPORT_PATH=github.com/hossinasaadi/gojskit

build_android:
	mkdir -p $(BUILD_DIR)
	gomobile bind  -a -ldflags $(LDFLAGS) -tags android -target=android -o $(ANDROID_ARTIFACT) $(IMPORT_PATH)

build_apple:
	go get golang.org/x/mobile/cmd/gomobile
	gomobile init
	mkdir -p $(BUILD_DIR)
	GOOS=ios gomobile bind -a -v -ldflags $(LDFLAGS) -tags maccatalyst -target=ios,iossimulator,macos,maccatalyst -iosversion=14 -o $(IOS_ARTIFACT) $(IMPORT_PATH)
