from usm.usm import USM

def load(file_path=None):
    """Load a USM secrets file."""
    return USM.load(file_path)

__all__ = ["load", "USM"]