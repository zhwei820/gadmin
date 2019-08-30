setlocal EnableDelayedExpansion

@echo off
echo !cd!

C:\Users\admin\go\bin\xorm.exe reverse mysql root:rootroot@tcp(localhost:3306)/gadmin?charset=utf8 .\templates\struct .\model
del /s /q /f  "./auth_*";
del /s /q /f  "./django*";
@pause