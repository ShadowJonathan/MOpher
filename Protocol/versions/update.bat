for /d %%i in (*) do cd %%i && go generate && cd ..