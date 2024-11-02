@echo off
chcp 65001
set product_name=%1

if exist "%product_name%" (
    rmdir /s /q "%product_name%"
    echo Product folder and environment variable have been deleted.
) else (
    echo Product folder does not exist.
)
