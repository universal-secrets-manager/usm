import * as fs from 'fs';
import * as path from 'path';
import * as yaml from 'js-yaml';
import * as crypto from 'crypto';

export class USM {
  private secrets: any;
  private cache: Map<string, any> = new Map();
  private cacheExpiry: number = 5 * 60 * 1000; // 5 minutes

  private constructor(secrets: any) {
    this.secrets = secrets;
  }

  static async load(filePath?: string): Promise<USM> {
    if (!filePath) {
      filePath = await USM.locateSecretsFile();
    }

    const fileContents = fs.readFileSync(filePath, 'utf8');
    const secrets = yaml.load(fileContents);
    return new USM(secrets);
  }

  private static async locateSecretsFile(): Promise<string> {
    // Implementation to find .secrets.yml in current or parent directories
    // This is a simplified version
    let currentDir = process.cwd();
    while (currentDir !== path.parse(currentDir).root) {
      const possiblePath = path.join(currentDir, '.secrets.yml');
      if (fs.existsSync(possiblePath)) {
        return possiblePath;
      }
      currentDir = path.dirname(currentDir);
    }
    throw new Error('Could not locate .secrets.yml file');
  }

  async get(key: string): Promise<string> {
    // Check if the secret exists
    if (!this.secrets.secrets || !this.secrets.secrets.dev || !this.secrets.secrets.dev[key]) {
      throw new Error(`Secret '${key}' not found`);
    }

    const secret = this.secrets.secrets.dev[key];
    
    // For now, we'll just return the CT field as a string
    // In a real implementation, we would decrypt the secret using the project key
    // and the file key stored in the secret
    return Buffer.from(secret.ct, 'base64').toString('utf8');
  }
}