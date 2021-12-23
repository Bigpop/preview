if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        xdg-open $1
elif [[ "$OSTYPE" == "darwin"* ]]; then
        open $1
elif [[ "$OSTYPE" == "cygwin" ]]; then
        start $1
elif [[ "$OSTYPE" == "win32" ]]; then
        start $1
else
        python -m webbrowser $1
fi