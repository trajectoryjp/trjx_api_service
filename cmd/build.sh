if [[ ${PWD##*/} == "cmd" ]];then
    cd ../
fi
buf dep update
buf generate
