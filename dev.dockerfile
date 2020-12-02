FROM golang:1.15.2-buster
WORKDIR /app
RUN go mod init go_course_cf
COPY . .
EXPOSE 3000
CMD bash
# docker run -d -p 3308:3306 -v mysql:/var/lib/mysql --name go_mysql --network go_red -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7
# docker run -d -p 8000:3000 --name go_base -w /app -v ${PWD}:/app --network go_red -it go_base bash
