@echo off
chcp 65001
set product_name=%1
set brand=%2
set file_name=%3
set env=%4

cd %product_name%
cd %env%
git add ./%brand%/%file_name%


