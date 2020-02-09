# bookstore_users-api
Users API

```
CREATE SCHEMA `users_db` DEFAULT CHARACTER SET utf8 COLLATE utf8_spanish2_ci;
```

```
CREATE TABLE `users_db`.`users` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `first_name` VARCHAR(45) NULL,
    `last_name`  VARCHAR(45) NULL,
    `email`     VARCHAR(45) NULL,
    `data_created` VARCHAR(45) NULL,
    PRIMARY KEY(`id`),
    UNIQUE INDEX `email_UNIQUE` (`email` ASC));
```

```
ALTER TABLE `users_db`.`users`
ADD COLUMN `status` VARCHAR(45) NOT NULL AFTER `email`,
ADD COLUMN `password` VARCHAR(32) NOT NULL AFTER `status`;
```


How deps work in go?

main.go it will try to find deps in this order 

require dep command

1. vendor (dep init / dep update ) -version, branch
2. GOPATH (got get .)
3. GOROOT

/usr/local/go/src

go get - download deps -> $GOPATH 

dep init - vendor

Gopack.toml - dependency management version ( 
sort of a pom.xml
)


Transitive dependencies

