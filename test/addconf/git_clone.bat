@echo off
chcp 65001
set product_name=%1

git clone "http://gitlab.paradise-soft.com.tw/configuration/%product_name%.git"

cd %product_name%
git pull
