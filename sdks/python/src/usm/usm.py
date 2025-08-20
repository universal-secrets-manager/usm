import os
import yaml
import base64

class USM:
    def __init__(self, secrets):
        self.secrets = secrets
        self.cache = {}
        self.cache_expiry = 300  # 5 minutes

    @classmethod
    def load(cls, file_path=None):
        """Load a USM secrets file."""
        if not file_path:
            file_path = cls._locate_secrets_file()

        with open(file_path, 'r') as f:
            secrets = yaml.safe_load(f)

        return cls(secrets)

    @staticmethod
    def _locate_secrets_file():
        """Locate the .secrets.yml file in the current or parent directories."""
        current_dir = os.getcwd()
        while current_dir != os.path.dirname(current_dir):
            possible_path = os.path.join(current_dir, '.secrets.yml')
            if os.path.exists(possible_path):
                return possible_path
            current_dir = os.path.dirname(current_dir)
        raise FileNotFoundError('Could not locate .secrets.yml file')

    def get(self, key):
        """Get and decrypt a secret."""
        # Check if the secret exists
        if not self.secrets.get('secrets', {}).get('dev', {}).get(key):
            raise KeyError(f"Secret '{key}' not found")
        
        secret = self.secrets['secrets']['dev'][key]
        
        # For now, we'll just return the CT field as a string
        # In a real implementation, we would decrypt the secret using the project key
        # and the file key stored in the secret
        return base64.b64decode(secret['ct']).decode('utf-8')