# Pull docker image

`docker pull mcr.microsoft.com/mssql/server:2017-latest`

# Docker run

`docker run -e 'ACCEPT_EULA=Y' -e 'SA_PASSWORD=1qaz@WSX3edc' -p 1433:1433 --name sql1  -d mcr.microsoft.com/mssql/server:2017-latest`

# Create sqlserver login & user

`create login sqladmin with PASSWORD='123456',CHECK_POLICY=OFF`
`ALTER SERVER Role [sysadmin] add Member [sqladmin]`
