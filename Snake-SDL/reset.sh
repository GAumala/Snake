./snake > error.log 2>&1
for i in {1..5} ; do
    zenity --question --text="Do you want to launch the game again?"
    if [ "$?" = 1 ] ; then
        break
    fi
    echo "Launching Snake Game..."
    ./snake > error.log 2>&1
done
