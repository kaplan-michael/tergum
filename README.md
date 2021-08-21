<p align="center">
  <h1 align="center">Tergum: Universal Backup Tool</h1>
  <p align="center">
    <a href="https://opensource.sikalabs.com"><img alt="SikaLabs" src="https://img.shields.io/badge/OPENSOURCE BY-SIKALABS-131480?style=for-the-badge"></a>
    <a href="https://sikalabs.com"><img alt="SikaLabs" src="https://img.shields.io/badge/-sikalabs.com-gray?style=for-the-badge"></a>
    <a href="mailto://opensource@sikalabs.com"><img alt="SikaLabs" src="https://img.shields.io/badge/-opensource@sikalabs.com-gray?style=for-the-badge"></a>
  </p>
</p>

## Why Tergum?

Tergum is simple tool provides centralized backup solution with multiple sources (databases, files, S3, ...) and multiple backup storages (S3, filesystem, ...). Tergum has native backup monitoring and alerts you when backup fails. Tergum also support backup encryption, compression and automatic recovery testing.

__Tergum is under active development, not all features are already implemented. Check [current project state](#current-project-state)__

## What "Tergum" means?

Tergum means backup in latin.

## Install

Install using Brew:

```
brew install sikalabs/tap/tergum
```

## Usage

See help:

```
tergum --help
```

### Tergum Config File

Tergum supports only JSON config file, but we're working on YAML support.

Config file examples are in [misc/example/config](./misc/example/config) directory

#### Basic Config Structure

```yaml
Meta:
  SchemaVersion: 3
Notification: <Notification>
Backups:
  - <Backup>
  - <Backup>
  - ...
```

#### Backup Block

```yaml
ID: <UniqueBackupID>
Source:
  Mysql: <BackupSourceMysqlConfiguration>
  Postgres: <BackupSourcePostgresConfiguration>
Middlewares:
  - <MiddlewareConfiguration>
  - ...
Destinations:
  - ID: <UniqueBackupDestinationID>
    Middlewares:
      - <MiddlewareConfiguration>
      - ...
    FilePath: <BackupDestinationFilePathConfiguration>
    File: <BackupDestinationFileConfiguration>
    S3: <BackupDestinationS3Configuration>
  - ...
```

#### GzipMiddlewareConfiguration

```yaml
Gzip: {}
```

#### Example BackupSourceMysqlConfiguration Block

```yaml
Host: "127.0.0.1"
Port: "3306"
User: "root"
Password: "root"
Database: "default"
```

#### Example BackupSourcePostgresConfiguration Block

```yaml
Host: "127.0.0.1"
Port: "15432"
User: "postgres"
Password: "pg"
Database: "postgres"
```

#### Example BackupDestinationFilePathConfiguration Block

```yaml
Path: "/backup/mysql-default.sql"
```

#### Example BackupDestinationFileConfiguration Block

```jsx
Dir: "/backup/"
Prefix: "mysql-default"
Suffix: "sql"
```

#### Example BackupDestinationS3Configuration Block

AWS:

```yaml
AccessKey: "admin"
SecretKey: "asdfasdf"
Endpoint: "https://minio.example.com"
BucketName: "tergum-backups"
Prefix: "mysql-default"
Suffix: "sql"
```

Minio:

```yaml
accessKey: "aws_access_key_id"
secretKey: "aws_secret_access_key"
region: "eu-central-1"
bucketName: "tergum-backups"
prefix: "mysql-default"
suffix: "sql"
```

#### Notification Block

```yaml
Backends: {
  Email:  <NotificationBackendEmail>
Target:
  - <NotificationTarget>
  - <NotificationTarget>
  - ...
```

#### Example NotificationBackendEmail Block

```yaml
SmtpHost: "mail.example.com"
SmtpPort: "25"
Email: "tergum@example.com"
Password: "asdfasdf"
```

#### NotificationTarget Block

```yaml
Email: <NotificationEmailTarget>
```

#### Example NotificationEmailTarget Block

```yaml
Emails:
  - ondrej@example.com
  - monitoring@example.com
```

## Current Project State

### Backup Sources

- [ ] Files
- [x] Postgres
- [x] MySQL
- [ ] S3
- [ ] MongoDB

### Backup Processors

- [x] GZIP Compression
- [ ] Symetric Encryption
- [ ] Asymetric Encryption
- [ ] GPG Encryption
- [ ] GPG Signatures

### Backup Storage

- [x] Files
- [x] S3

### Notification

- [x] Email
- [ ] Slack
- [ ] Pagerduty
