const { USM } = require('@usm/secrets');

async function main() {
  try {
    const usm = await USM.load();
    const dbUrl = await usm.get('DB_URL');
    console.log('DB_URL:', dbUrl);
  } catch (error) {
    console.error('Error:', error.message);
  }
}

main();