# Generated by Django 2.1.4 on 2019-08-22 12:47

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('gadmin', '0001_initial'),
    ]

    operations = [
        migrations.CreateModel(
            name='RolePolicy',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('role_key', models.CharField(db_index=True, max_length=255)),
                ('policy_path', models.CharField(db_index=True, max_length=255)),
            ],
        ),
        migrations.RemoveField(
            model_name='menu',
            name='name',
        ),
        migrations.RemoveField(
            model_name='menumeta',
            name='menu_name',
        ),
        migrations.RemoveField(
            model_name='rolemenu',
            name='menu_name',
        ),
        migrations.AddField(
            model_name='menumeta',
            name='menu_id',
            field=models.IntegerField(db_index=True, default=0),
        ),
        migrations.AddField(
            model_name='rolemenu',
            name='menu_id',
            field=models.IntegerField(db_index=True, default=0),
        ),
        migrations.AlterField(
            model_name='policyconfig',
            name='full_path',
            field=models.CharField(db_index=True, max_length=255, unique=True),
        ),
        migrations.AlterField(
            model_name='rolemenu',
            name='role_key',
            field=models.CharField(db_index=True, max_length=255),
        ),
    ]