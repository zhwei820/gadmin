import inspect

from django.contrib import admin

from . import models as app_models

for attr in dir(app_models):
    model = getattr(app_models, attr)
    if not inspect.isclass(model):
        continue

    try:
        admin.site.register(model)
    except:
        pass
