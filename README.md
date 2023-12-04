# AliyunOSSManager

## 1. 配置文件修改

配置文件为yaml格式，注意 ":" 后的空格，一定要保留

```yaml
# oss-admin config.yaml

Endpoint: "your-Endpoint"
AccessKeyId: "your-AccessKeyId"
AccessKeySecret: "your-AccessKeySecret"
BucketName: "your-bucket"
```

在配置文件中填入你要使用的OSS的Endpoint、AccessKeyId、AccessKeySecret、BucketName即可开始使用。

BucketName可设置为空值“”，通过--name进行手动指定



## 2. 查看有哪些Bucket

```bash
oss-admin bucket
```

注：若你需要查询多个账号，可复制config/config.yaml，然后通过“-f”指定配置文件查询不同账号，其他操作的同理

例：

```bash
cp config/config.yaml config/zhangsan.yaml
oss-admin bucket -f config/zhangsan.yaml
```

## 3. 查看bucket中的文件列表

```bash
oss-admin ls <Bucket>
oss-admin ls -f config/zhangsan.yaml
# 通过--name查看其他bucket
oss-admin ls -f config/zhangsan.yaml --name zhangsan_pro
oss-admin ls -f config/zhangsan.yaml --name zhangsan_dev
```

## 4. 查看Bucket所使用的存储空间大小

```bash
oss-admin du <Bucket>
oss-admin du -f config/zhangsan.yaml
```

## 5. 上传文件至OSS

```bash
oss-admin push --local oss.txt
# 如上传当前目录下oss.txt到bucket(zhangsan-pro)的default下
oss-admin push --local oss.txt --cloud default/oss.txt
```

## 6.从OSS下载文件

```bash
oss-admin pull --cloud default/aliyun.txt
# 如需要下载bucket(zhangsan-pro)的default下的oss.txt至/aliyun/下：
oss-admin pull --cloud default/aliyun.txt --local /aliyun/aliyun.txt
```

