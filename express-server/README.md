# circles-server-node

<img src="https://img.shields.io/badge/TypeScript-3178C6?style=flat-square&logo=TypeScript&logoColor=white"/> <img src="https://img.shields.io/badge/Yarn-2C8EBB?style=flat-square&logo=Yarn&logoColor=white"/> <img src="https://img.shields.io/badge/NGINX-009639?style=flat-square&logo=NGINX&logoColor=white"/> <img src="https://img.shields.io/badge/aws-232F3E?style=flat-square&logo=Amazonaws&logoColor=white"/>

| <a id="a1"></a>목차         |
| --------------------------- |
| [1. 프로젝트 init](#1)<br/> |
| [2. 브랜치 ](#2)<br/>       |

<br/>

# <a id="1"></a>[1](#a1). 프로젝트 init

> ---
>
> # Yarn
>
> - yarn이 없다면 yarn 설치 후 진행해주세요
> - mac :: brew install yarn
>
> # local 세팅
>
>       git clone ${주소} && yarn install
>
> # 서버 구동 (with nodemon)
>
> ## yarn main
>
> - 배포 서버를 위한 구동
> - AWS db, domain setting으로 서버가 구동됩니다 (진행중)
>
> ## yarn dev
>
> - 로컬 서버를 위한 구동
> - localhost db, domain setting으로 서버가 구동됩니다.
> - 이 과정에서 mysql 오류가 날 수 있습니다.
>
> ---

# <a id="2"></a>[2](#a1). 브랜치

# <a id="3"></a>[3](#a1). 주의사항

# <a id="3"></a>[4](#a1). Trouble shooting

> ---
>
> # MYSQL
>
> ### erno: 1251
>
> code: 'ER_NOT_SUPPORTED_AUTH_MODE',
> errno: 1251,
> sqlMessage: 'Client does not support authentication protocol requested by server; consider upgrading MySQL client',
> sqlState: '08004',
> fatal: true
>
> 위와 같이 나오는 경우 mysql에 접속해서 아래와 같은 명령어를 입력
> `ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';` > `flush privileges;`
>
> https://stackoverflow.com/questions/50093144/mysql-8-0-client-does-not-support-authentication-protocol-requested-by-server
>
> ---
