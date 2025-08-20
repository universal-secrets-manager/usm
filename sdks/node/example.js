const { load } = require('./dist/index.js');

async function main() {
  try {
    // Load the secrets file
    const usm = await load('./test/fixtures/.secrets.yml');
    
    // Get a secret value
    const value = await usm.get('TEST_KEY');
    
    console.log(`Secret value: ${value}`);
  } catch (error) {
    console.error('Error:', error.message);
  }
}

main();