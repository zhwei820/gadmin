import os

ENV = os.getenv('MARS_API_ENV', 'dev')


try:
    if ENV == 'dev':
        from .local_config import *  # noqa
    elif ENV == 'test':
        from .test_config import *  # noqa
    elif ENV == 'test_ci':
        from .test_ci_config import *  # noqa
    elif ENV == 'pro':
        from .pro_config import *  # noqa
except Exception as e:
    print('*** Import local config failed ***')
    print(e)


export = [
    DEBUG,
    ALLOWED_HOSTS,
    CELERY_BROKER_URL,
    CELERY_RESULT_BACKEND,
    RAVEN_CONFIG,
]

# 腾讯云 CDN 配置
# COS_URL = 'https://newmedia.pinpaibao.com.cn'  # 自定义域名， 不配置将使用COS默认域名
COS_URL = 'https://new-ppb-1251001081.coscd.myqcloud.com'  # 自定义域名， 不配置将使用COS默认域名
COS_BUCKET = 'new-ppb'
REGION = 'ap-chengdu'

# ua
DEFAULT_MOCK_UA = """Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) \
AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36"""
