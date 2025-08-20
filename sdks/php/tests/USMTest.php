<?php

namespace Usm\Tests;

use PHPUnit\Framework\TestCase;
use Usm\USM;

class USMTest extends TestCase
{
    public function testLoad()
    {
        // This is a placeholder test
        // In a real test, you would mock the file system
        $usm = USM::load(__DIR__ . '/fixtures/.secrets.yml');
        $this->assertInstanceOf(USM::class, $usm);
    }

    public function testGet()
    {
        // This is a placeholder test
        $usm = USM::load(__DIR__ . '/fixtures/.secrets.yml');
        $value = $usm->get('TEST_KEY');
        $this->assertEquals('decrypted_value_for_TEST_KEY', $value);
    }
}