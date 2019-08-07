# coding: utf-8
# 19-4-3 下午3:40

import logging
from celery.schedules import crontab

LOG_LEVEL = logging.INFO

DEBUG = True
ALLOWED_HOSTS = ['*']

MYSQL_HOST = '127.0.0.1'
MYSQL_PORT = 3306
MYSQL_USERNAME = 'root'
MYSQL_PASSWORD = 'rootroot'


REDIS_LOCATION = 'redis://redis:6379/0'
CELERY_BROKER_URL = REDIS_LOCATION
CELERY_RESULT_BACKEND = REDIS_LOCATION

# RAVEN
RAVEN_CONFIG = {
    'dsn': '',
}
