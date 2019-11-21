# GoTranslate

This is a simple example on how to combine a Go CLI tool with a lightweight Qt GUI application that calls the Go CLI tool.

In this example, the Go CLI tool and the Qt GUI application are separate and are not linked together, so that we can still use the cross-compiling capabilities built into Go and do not need to use CGo.

## Cross-compile the CLI tool for Windows

### Install Go

This is described elsewhere.

### Build the CLI tool for Windows

```bash
cd gotranslate/

# Get and locally bundle all dependencies in the source tree
go mod init gotranslate
go mod tidy
go mod vendor

# Build for Windows
GOOS=windows go build *.go
mv bing.exe gotranslatecli.exe
```

No cross-compiler involved so far besides what is built natively into Go. Neat!

## Cross-compile the GUI application for Windows

This has been tested using Xubuntu 18.04.2 and QtCreator-5.8.0-x86_64.AppImage.

### Install crosscompiler

Based on http://mexuaz.com/mxe-in-qt-creator/

```bash
# Probably this list can be reduced
sudo apt-get --no-install-recommends -qq -y install mxe-i686-w64-mingw32.shared-qt3d mxe-i686-w64-mingw32.shared-qtactiveqt mxe-i686-w64-mingw32.shared-qtbase mxe-i686-w64-mingw32.shared-qtcanvas3d mxe-i686-w64-mingw32.shared-qtcharts mxe-i686-w64-mingw32.shared-qtconnectivity mxe-i686-w64-mingw32.shared-qtdatavis3d mxe-i686-w64-mingw32.shared-qtdeclarative mxe-i686-w64-mingw32.shared-qtgamepad mxe-i686-w64-mingw32.shared-qtgraphicaleffects mxe-i686-w64-mingw32.shared-qtimageformats mxe-i686-w64-mingw32.shared-qtlocation mxe-i686-w64-mingw32.shared-qtmultimedia mxe-i686-w64-mingw32.shared-qtofficeopenxml mxe-i686-w64-mingw32.shared-qtpurchasing mxe-i686-w64-mingw32.shared-qtquickcontrols mxe-i686-w64-mingw32.shared-qtquickcontrols2 mxe-i686-w64-mingw32.shared-qtscript mxe-i686-w64-mingw32.shared-qtscxml mxe-i686-w64-mingw32.shared-qtsensors mxe-i686-w64-mingw32.shared-qtserialbus mxe-i686-w64-mingw32.shared-qtserialport mxe-i686-w64-mingw32.shared-qtservice mxe-i686-w64-mingw32.shared-qtsvg mxe-i686-w64-mingw32.shared-qtsystems mxe-i686-w64-mingw32.shared-qttools mxe-i686-w64-mingw32.shared-qttranslations mxe-i686-w64-mingw32.shared-qtvirtualkeyboard mxe-i686-w64-mingw32.shared-qtwebchannel mxe-i686-w64-mingw32.shared-qtwebkit mxe-i686-w64-mingw32.shared-qtwebsockets mxe-i686-w64-mingw32.shared-qtwinextras mxe-i686-w64-mingw32.shared-qtxlsxwriter mxe-i686-w64-mingw32.shared-qtxmlpatterns mxe-i686-w64-mingw32.make
```

### Configure Qt Creator to crosscompile for Windows

* In Qt Creator, click on the 'Projects' icon in the left-hand bar
* Edit build configuration -> Add -> Clone Selected -> "Windows"
* Build Steps -> Add Build Step -> Custom Build Step
Command: `/usr/lib/mxe/usr/i686-w64-mingw32.shared/qt5/bin/qmake`
Arguments: `%{sourceDir}`
Working directory: `%{buildDir}`

Move this build step to the first position

Disable the original qmake build step

Build Environment -> PATH -> append to the end:

`:/usr/lib/mxe/usr/i686-w64-mingw32.shared/qt5/bin::/usr/lib/mxe/usr/i686-w64-mingw32.shared/bin:/usr/lib/mxe/usr/bin/`

### Build for Windows in Qt Creator

Click the 'Build' icon (hammer icon in the bottom left of the screen)

Check 'Compile Output', it will tell where the .exe has been created, e.g.,

`/home/me/build-GoTranslate-Desktop_Qt_5_8_0_GCC_64bit-Debug`

### Deploy Qt to the application directory

Need to copy the needed Qt libraries into that directory (or possibly use `windeployqt`)


```bash
# Probably this list can be reduced
cp /usr/lib/mxe/usr/i686-w64-mingw32.shared/qt5/bin/*.dll /usr/lib/mxe/usr/i686-w64-mingw32.shared/bin/*.dll /home/me/build-GoTranslate-Desktop_Qt_5_8_0_GCC_64bit-Debug/release/
cp -r /usr/lib/mxe/usr/i686-w64-mingw32.shared/qt5/plugins/{platforms,imageformats,styles} /home/me/build-GoTranslate-Desktop_Qt_5_8_0_GCC_64bit-Debug/release/
# cp -r /usr/lib/mxe/usr/i686-w64-mingw32.shared/qt5/qml/ /home/me/build-GoTranslate-Desktop_Qt_5_8_0_GCC_64bit-Debug/release/ # If Qml was used
```

### Zip it

```bash
zip -r QtApp.zip /home/me/build-GoTranslate-Desktop_Qt_5_8_0_GCC_64bit-Debug/release
```