<?php

namespace Usm;

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

        $content = file_get_contents($filePath);
        $secrets = json_decode($content, true);
        if (json_last_error() !== JSON_ERROR_NONE) {
            throw new \Exception('Invalid JSON in secrets file: ' . json_last_error_msg());
        }
        return new USM($secrets);
    }

    private static function locateSecretsFile(): string
    {
        // Implementation to find .secrets.json in current or parent directories
        // This is a simplified version
        $currentDir = getcwd();
        $root = dirname($currentDir);

        while ($currentDir !== $root) {
            $possiblePath = $currentDir . DIRECTORY_SEPARATOR . '.secrets.json';
            if (file_exists($possiblePath)) {
                return $possiblePath;
            }
            $currentDir = dirname($currentDir);
        }

        throw new \Exception('Could not locate .secrets.json file');
    }

    public function get(string $key): string
    {
        // Implementation to get and decrypt a secret
        // This is a placeholder
        return "decrypted_value_for_$key";
    }
}