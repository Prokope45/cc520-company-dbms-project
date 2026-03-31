"""Common helper functions used across the test suite."""

import random
import string


def helper_get_random_string(size=10, chars=string.ascii_uppercase + string.digits):
    """Return a random uppercase alphanumeric string of the requested size.

    Args:
        size: Desired output string length.
        chars: Character set used for random selection.

    Returns:
        str: Randomly generated string value.
    """
    return ''.join(random.choice(chars) for _ in range(size))
