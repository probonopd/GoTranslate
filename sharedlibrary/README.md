# GoTranslate

This is a simple example on how to combine Go code compiled as a shared library with a lightweight Qt GUI application that calls the Go CLI tool. This has been tested using Xubuntu 18.04.2 and QtCreator-5.8.0-x86_64.AppImage.

## Build the Go code as a shared library

### Install Go

This is described elsewhere.

### Modify the Go code

* The package must be amain package
* The source must import the pseudo-package “C”
* Use the `//export` comment to annotate functions you wish to make accessible to other languages
* It must `return C.CString(...)` if it wants to return a string. The function must be declared as returning a `*C.char`
* An empty `main()` function must be declared
* The package is compiled using the `-buildmode=c-shared` build flag to create the shared object binary:
  `go build -o libgotranslate.so -buildmode=c-shared *.go` - note the "lib" in the library name, which is important

### Modify the Qt application to call into the shared library

* The Qt `.cpp` file needs: `#include "gotranslate.h"`
* **FIXME:** How can I make it to accept: #include `"gotranslate/libgotranslate.h"`
* The Qt `.pro` file needs: `LIBS += -L/home/me/GoTranslate/bridged/gotranslate -lgotranslate # libgotranslate.so`
* The code needs to convert between Go and Qt data types, e.g.,:

```c++
// Get the string to be translated from the input text box
QString toBeTranslated =  ui->input->toPlainText();

// Call into shared library written in Go to do the actual translation
GoString language = {ui->comboBox->currentText().toUtf8() + "\0", ui->comboBox->currentText().toUtf8().length()};
GoString text = {toBeTranslated.toUtf8() + "\0", toBeTranslated.toUtf8().length()};
GoString translated = DoTranslate(language, text);
QString output = QString::fromUtf8((char*)(translated.p));
```

Now compile. It should work.