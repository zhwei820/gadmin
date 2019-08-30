# Generated by Django 2.1.8 on 2019-08-30 05:52

from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='Casbinrule',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('p_type', models.CharField(max_length=255)),
                ('v0', models.CharField(max_length=255)),
                ('v1', models.CharField(max_length=255)),
                ('v2', models.CharField(max_length=255)),
                ('v3', models.CharField(max_length=255)),
                ('v4', models.CharField(max_length=255)),
                ('v5', models.CharField(max_length=255)),
            ],
        ),
        migrations.CreateModel(
            name='PolicyConfig',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('full_path', models.CharField(db_index=True, max_length=255, unique=True)),
                ('name', models.CharField(max_length=255)),
                ('descrption', models.CharField(blank=True, max_length=255, null=True)),
            ],
        ),
        migrations.CreateModel(
            name='RoleConfig',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('role_key', models.CharField(max_length=255, unique=True)),
                ('name', models.CharField(max_length=255)),
                ('descrption', models.CharField(blank=True, max_length=255, null=True)),
            ],
        ),
        migrations.CreateModel(
            name='RolePolicy',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('role_key', models.CharField(db_index=True, max_length=255)),
                ('policy_path', models.CharField(db_index=True, max_length=255)),
            ],
        ),
        migrations.CreateModel(
            name='User',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('status', models.IntegerField()),
                ('user_name', models.CharField(max_length=255, unique=True)),
                ('nick_name', models.CharField(max_length=255)),
                ('password', models.CharField(max_length=255)),
                ('email', models.CharField(blank=True, max_length=255, null=True)),
                ('phone', models.CharField(blank=True, max_length=255, null=True)),
                ('sex', models.IntegerField()),
                ('create_time', models.DateTimeField(auto_now_add=True)),
                ('update_time', models.DateTimeField(auto_now=True)),
                ('add_user_id', models.IntegerField()),
                ('introduction', models.CharField(blank=True, max_length=255, null=True)),
                ('avatar', models.CharField(blank=True, max_length=255, null=True)),
                ('role_keys', models.CharField(db_index=True, default='', max_length=255)),
            ],
        ),
    ]
