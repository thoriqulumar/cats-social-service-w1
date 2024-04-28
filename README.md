The code structure that used on this project is follow this convention: https://github.com/golang-standards/project-layout

the framework that used on this project: 
- web framework https://gin-gonic.com/
- db migration https://github.com/golang-migrate/migrate


# How To Run on Local
1. Install Go 1.22.2 
2. run postgre on local
3. run migration script to sync your local database (if needed)
4. ```make run-app```
5. Test the app
   ```curl localhost:8080/ping```


# Project Requirements 
https://openidea-projectsprint.notion.site/Cats-Social-9e7639a6a68748c38c67f81d9ab3c769
