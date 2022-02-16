#! /bin/sh
echo "------------- initialize mysql -------------"

# mysql 체크
mysql -V

if [ $? -eq 127 ]; then
    echo "mysql이 존재하지 않습니다"
    echo "mysql을 설치합니다."
    brew install mysql
    if [ $? -ne 0]; then
        echo "설치에 실패했습니다."
        exit 127
    fi
fi

brew services start mysql
# mysql_secure_installation