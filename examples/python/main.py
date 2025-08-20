from usm import load

def main():
    try:
        usm = load()
        db_url = usm.get("DB_URL")
        print(f"DB_URL: {db_url}")
    except Exception as e:
        print(f"Error: {e}")

if __name__ == "__main__":
    main()