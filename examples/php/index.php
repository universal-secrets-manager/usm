<?php

require_once 'vendor/autoload.php';

use Usm\USM;

try {
    $usm = USM::load();
    $dbUrl = $usm->get('DB_URL');
    echo "DB_URL: $dbUrl\n";
} catch (Exception $e) {
    echo "Error: " . $e->getMessage() . "\n";
}