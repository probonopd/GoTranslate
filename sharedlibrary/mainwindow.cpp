#include "mainwindow.h"
#include "ui_mainwindow.h"
#include<QDebug>
#include<QProcess>
#include "gotranslate.h"

MainWindow::MainWindow(QWidget *parent) :
    QMainWindow(parent),
    ui(new Ui::MainWindow)
{
    ui->setupUi(this);

    // Add available languages to the dropdown menu
    QStringList langs = { "ar", "bg", "ca", "cs", "da", "de", "el", "en", "es", "fi", "fr",
                          "he", "hi", "hr", "hu", "id", "it", "ja", "ko", "ms", "nl", "nb",
                          "pl", "pt", "ro", "ru", "sk", "sl", "sv", "ta", "te", "th", "tr",
                          "vi", "zh-Hans", "zh-Hant", "yue" };

    foreach (const QString &str, langs) {
           ui->comboBox->addItem (str);
     }
     ui->comboBox->setCurrentText("de");
}

MainWindow::~MainWindow()
{
    delete ui;
}

// Right-clicked on the button ni the UI editor, "Go to slot..." created this method
// and declared it in the header file
void MainWindow::on_translateButton_clicked()
{
        // Get the string to be translated from the input text box
        QString toBeTranslated =  ui->input->toPlainText();

        // Call into shared library written in Go to do the actual translation
        GoString language = {ui->comboBox->currentText().toUtf8() + "\0", ui->comboBox->currentText().toUtf8().length()};
        GoString text = {toBeTranslated.toUtf8() + "\0", toBeTranslated.toUtf8().length()};
        GoString translated = DoTranslate(language, text);
        QString output = QString::fromUtf8((char*)(translated.p));

        // TODO: Error handling

        // Output the translated string into the output text box
        ui->output->setText(output);

}
