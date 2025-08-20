# PHP SDK

This directory contains the PHP SDK for USM.

## Installation

```bash
composer require usm/secrets
```

## Usage

```php
<?php
require_once 'vendor/autoload.php';

use Usm\USM;

$usm = USM::load();
$dbUrl = $usm->get('DB_URL');
```