FROM php:7.2-apache

COPY . /var/www/html


RUN rm -rf /var/www/html/var/cache/prod \
    && rm -rf /var/www/html/var/cache/dev
RUN chown -R www-data /var/www
RUN chmod -R 777 /var/www

COPY ./php.ini /usr/local/etc/php/php.ini
COPY ./000-default.conf /etc/apache2/sites-enabled/000-default.conf

EXPOSE 80
