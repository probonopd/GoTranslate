#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>

namespace Ui {
class MainWindow;
}

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    explicit MainWindow(QWidget *parent = 0);
    ~MainWindow();

// This method was automatically created using Qt Creator,
    //
private slots:
    void on_translateButton_clicked();

private:
    Ui::MainWindow *ui;
};

#endif // MAINWINDOW_H
