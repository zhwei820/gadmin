# This is an auto-generated Django model module.
# You'll have to do the following manually to clean this up:
#   * Rearrange models' order
#   * Make sure each model has one field with primary_key=True
#   * Make sure each ForeignKey has `on_delete` set to the desired behavior.
#   * Remove `managed = False` lines if you wish to allow Django to create, modify, and delete the table
# Feel free to rename the models, but don't rename db_table values or field names.
from django.db import models


class CasbinRule(models.Model):
    p_type = models.CharField(max_length=255)
    v0 = models.CharField(max_length=255)
    v1 = models.CharField(max_length=255)
    v2 = models.CharField(max_length=255)
    v3 = models.CharField(max_length=255)
    v4 = models.CharField(max_length=255)
    v5 = models.CharField(max_length=255)

    class Meta:
        # managed = False
        db_table = 'casbin_rule'


class Menu(models.Model):
    menu_path = models.CharField(max_length=255)
    component = models.CharField(max_length=255)
    redirect = models.CharField(max_length=255)
    name = models.CharField(unique=True, max_length=255)
    hidden = models.IntegerField(blank=True, null=True)
    alwaysshow = models.IntegerField(blank=True, null=True)
    sort = models.IntegerField(blank=True, null=True)
    parent_name = models.CharField(max_length=255, blank=True, null=True)
    auto_create = models.IntegerField(blank=True, null=True)

    class Meta:
        # managed = False
        db_table = 'menu'


class MenuMeta(models.Model):
    menu_name = models.CharField(unique=True, max_length=255)
    title = models.CharField(max_length=255)
    icon = models.CharField(max_length=255, blank=True, null=True)
    nocache = models.IntegerField(blank=True, null=True)

    class Meta:
        # managed = False
        db_table = 'menu_meta'


class PolicyName(models.Model):
    full_path = models.CharField(unique=True, max_length=255)
    name = models.CharField(max_length=255)
    descrption = models.CharField(max_length=255, blank=True, null=True)

    class Meta:
        # managed = False
        db_table = 'policy_name'


class RoleMenu(models.Model):
    role_key = models.CharField(max_length=255)
    menu_name = models.CharField(max_length=255)

    class Meta:
        # managed = False
        db_table = 'role_menu'


class RoleName(models.Model):
    role_key = models.CharField(unique=True, max_length=255)
    name = models.CharField(max_length=255)
    descrption = models.CharField(max_length=255, blank=True, null=True)

    class Meta:
        # managed = False
        db_table = 'role_name'


class User(models.Model):
    status = models.IntegerField()
    user_name = models.CharField(unique=True, max_length=255)
    nick_name = models.CharField(max_length=255)
    password = models.CharField(max_length=255)
    email = models.CharField(max_length=255, blank=True, null=True)
    phone = models.CharField(max_length=255, blank=True, null=True)
    sex = models.IntegerField()
    age = models.IntegerField()
    add_time = models.DateTimeField()
    update_time = models.DateTimeField()
    add_user_id = models.IntegerField()
    introduction = models.CharField(db_column='Introduction', max_length=255, blank=True, null=True)  # Field name made lowercase.
    avatar = models.CharField(max_length=255, blank=True, null=True)

    class Meta:
        # managed = False
        db_table = 'user'
