import { describe, it, expect } from 'vitest';
import { USM } from '../src/usm';

describe('USM', () => {
  it('should load secrets', async () => {
    // This is a placeholder test
    // In a real test, you would mock the file system
    const usm = await USM.load('./test/fixtures/.secrets.yml');
    expect(usm).toBeInstanceOf(USM);
  });

  it('should get a secret', async () => {
    // This is a placeholder test
    const usm = await USM.load('./test/fixtures/.secrets.yml');
    const value = await usm.get('TEST_KEY');
    expect(value).toBe('decrypted_value_for_TEST_KEY');
  });
});