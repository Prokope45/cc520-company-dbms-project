"""Custom exception for handling cases where a requested record is not found in the database.
Author: Josh Weese
"""
class RecordNotFoundException(Exception):
    """Represent an error raised when a requested database record does not exist."""

    def __init__(self, key):
        """Initialize the exception with the missing record key.

        Args:
            key: Identifier value used to look up the missing record.

        Returns:
            None: Constructor initializes exception state.
        """
        self.key = key
        self.message = f"The requested record with key [{key}] does not exist."
        super().__init__(self.message)
