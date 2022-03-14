FROM php:7.3-apache

RUN apt-get update && apt-get install -y \
    vim

# install xdebug
RUN pecl install xdebug \
    && docker-php-ext-enable xdebug