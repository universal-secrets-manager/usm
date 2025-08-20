from usm import load

def main():
    try:
        # Load the secrets file
        usm = load("./tests/fixtures/.secrets.yml")
        
        # Get a secret value
        value = usm.get("TEST_KEY")
        
        print(f"Secret value: {value}")
    except Exception as e:
        print(f"Error: {e}")

if __name__ == "__main__":
    main()