<?php

namespace Usm;

use Symfony\Component\Yaml\Yaml;

class USM
{
    private array $secrets;
    private array $cache = [];
    private int $cacheExpiry = 300; // 5 minutes

    public function __construct(array $secrets)
    {
        $this->secrets = $secrets;
    }

    public static function load(?string $filePath = null): USM
    {
        if (!$filePath) {
            $filePath = self::locateSecretsFile();
        }

        $secrets = Yaml::parseFile($filePath);
        return new USM($secrets);
    }

    private static function locateSecretsFile(): string
    {
        // Implementation to find .secrets.yml in current or parent directories
        // This is a simplified version
        $currentDir = getcwd();
        $root = dirname($currentDir);

        while ($currentDir !== $root) {
            $possiblePath = $currentDir . DIRECTORY_SEPARATOR . '.secrets.yml';
            if (file_exists($possiblePath)) {
                return $possiblePath;
            }
            $currentDir = dirname($currentDir);
        }

        throw new \Exception('Could not locate .secrets.yml file');
    }

    public function get(string $key): string
    {
        // Implementation to get and decrypt a secret
        // This is a placeholder
        return "decrypted_value_for_$key";
    }
}