set currentdir=%~dp0
cd /d %currentdir%

@echo off
if "%1" == "h" goto begin
start mshta vbscript:createobject("wscript.shell").run("%~nx0 h",0)(window.close)&&exit
:begin


python recycle.py

