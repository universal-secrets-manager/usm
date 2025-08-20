import pytest
from usm import USM


def test_load():
    """Test loading a secrets file."""
    # This is a placeholder test
    # In a real test, you would mock the file system
    usm = USM.load("./tests/fixtures/.secrets.yml")
    assert isinstance(usm, USM)


def test_get():
    """Test getting a secret."""
    # This is a placeholder test
    usm = USM.load("./tests/fixtures/.secrets.yml")
    value = usm.get("TEST_KEY")
    assert value == "decrypted_value_for_TEST_KEY"