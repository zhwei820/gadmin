# This is an auto-generated Django model module.
# You'll have to do the following manually to clean this up:
#   * Rearrange models' order
#   * Make sure each model has one field with primary_key=True
#   * Make sure each ForeignKey has `on_delete` set to the desired behavior.
#   * Remove `managed = False` lines if you wish to allow Django to create, modify, and delete the table
# Feel free to rename the models, but don't rename db_table values or field names.
from django.db import models


class Casbinrule(models.Model):
    p_type = models.CharField(max_length=255)
    v0 = models.CharField(max_length=255)
    v1 = models.CharField(max_length=255)
    v2 = models.CharField(max_length=255)
    v3 = models.CharField(max_length=255)
    v4 = models.CharField(max_length=255)
    v5 = models.CharField(max_length=255)

    class Meta:
        # managed = False
        #db_table = 'casbin_rule'
        pass


class PolicyConfig(models.Model):
    full_path = models.CharField(unique=True, max_length=255, db_index=True)
    name = models.CharField(max_length=255)
    label = models.CharField(max_length=255)
    descrption = models.CharField(max_length=255, blank=True, null=True)

    class Meta:
        # managed = False
        #db_table = 'policy_name'
        pass


class RoleConfig(models.Model):  # 角色配置表
    role_key = models.CharField(unique=True, max_length=255)
    name = models.CharField(max_length=255)
    descrption = models.CharField(max_length=255, blank=True, null=True)

    class Meta:
        # managed = False
        #db_table = 'role_name'
        pass

class User(models.Model):
    status = models.IntegerField()
    username = models.CharField(unique=True, max_length=255)
    nickname = models.CharField(max_length=255)
    password = models.CharField(max_length=255)
    email = models.CharField(max_length=255, blank=True, null=True)
    phone = models.CharField(max_length=255, blank=True, null=True)
    create_time = models.DateTimeField(auto_now_add=True)
    update_time = models.DateTimeField(auto_now=True)
    add_user_id = models.IntegerField()
    introduction = models.CharField(max_length=255, blank=True, null=True)
    avatar = models.CharField(max_length=255, blank=True, null=True)

    class Meta:
        # managed = False
        #db_table = 'user'
        pass

