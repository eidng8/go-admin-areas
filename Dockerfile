FROM mysql:8

EXPOSE 3306 33060 33061

ENV MYSQL_DATABASE=admin_areas
ENV MYSQL_USER=du
ENV MYSQL_ROOT_HOST=%

RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' > /etc/timezone

RUN echo 'CREATE DATABASE IF NOT EXISTS `admin_areas` DEFAULT CHARACTER SET `utf8mb4` COLLATE `utf8mb4_0900_ai_ci`;' > /docker-entrypoint-initdb.d/init.sql
