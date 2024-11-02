@echo off
chcp 65001
set product_name=%1
set brand=%2
set commit_message=%3
set branch_name=%4
set env=%5

cd %product_name%
git checkout master
git pull
git checkout -b "%branch_name%"
git commit -m %commit_message%
git push origin %branch_name%
