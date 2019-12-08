#!/bin/sh

PROJECT_DIR=$1
INTERACTIVE=0
INIT_GIT=1
CREATE_GIT_HOOK=1
TEMPLATE_DIR=$(pwd)

# Determine, whether we use "interactive" mode or not
for  ARG in $@
do
    if [ ${ARG} != $1 ] 
    then
        if [ ${ARG} = "interactive" ] || [ ${ARG} = "interactively" ]
        then
            INTERACTIVE=1
        fi
    fi
done

# If project already exists, we should ask user, whether he/she want to continue or not
if [ -e ${PROJECT_DIR} ]
then
    if [ ${INTERACTIVE} = 1 ]
    then
        read -p "${PROJECT_DIR} is not empty, do you want to continue (Y/n)? " answer
        if [ $answer = "n" ] || [ $answer = "N" ]
        then
            echo "Project not created."
            exit 0
        fi
    else
        echo "${PROJECT_DIR} is not empty. Project not created."
        exit 0
    fi 
fi

# Copy project to PROJECT_DIR, add .gitignore and change to the project directory
cp -r ${TEMPLATE_DIR}/project ${PROJECT_DIR}
cp ${TEMPLATE_DIR}/resources/gitignore ${PROJECT_DIR}/.gitignore

cd ${PROJECT_DIR}

# Perform substitution
sed -i "s/{{PROJECT_DIR}}/${PROJECT_DIR}/g" "./README.md"

# Ask user, whether he/she want to init git repo or not
if [ ${INTERACTIVE} = 1 ]
then
    read -p "Do you want to init git repository (Y/n)? " answer
    if [ $answer = "n" ] || [ $answer = "N" ]
    then
        INIT_GIT=0
    fi
fi

# Init git and git-hooks
if [ ${INIT_GIT} = 1 ]
then
    git init

    # Ask user, whether want to create git hooks or not
    if [ ${INTERACTIVE} = 1 ]
    then
        read -p "Do you want to use git-hooks (Y/n)? " answer
        if [ $answer = "n" ] || [ $answer = "N" ]
        then
            CREATE_GIT_HOOK=0
        fi
    fi

    # Add githooks
    if [ ${CREATE_GIT_HOOK} = 1 ]
    then
        for event in $(ls ${TEMPLATE_DIR}/resources/hooks)
        do
            cp ${TEMPLATE_DIR}/resources/hooks/${event} ${event}
            echo "#!/bin/sh\n./${event}" > .git/hooks/${event}
        done
    fi
fi

echo "Project created. Perform 'cd "${PROJECT_DIR}"' and enjoy :)"