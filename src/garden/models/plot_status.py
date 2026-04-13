"""
Data model for the plot status.
Author: Jared Paubel
"""

class Plot:
    """Represent a plot status record returned by the repository layer."""

    def __init__(
        self,
        status_id: int,
        status_name: str
    ):
        """Initialize a plot model with location and size details.

        Args:
            status_id: Unique status identifier.
            status_name: Name of the status.

        Returns:
            None: Constructor initializes instance state.
        """
        self._status_id = status_id
        self._status_name = status_name

    @property
    def status_id(self) -> int:
        """Return the unique identifier for the plot status.

        Returns:
            int: status id.
        """
        return self._status_id

    @property
    def status_name(self) -> str:
        """Return the descriptive name of the status.

        Returns:
            str: plot status.
        """
        return self._status_name
