"""
Data model for a plot.
Author: Jared Paubel
"""

class Plot:
    """Represent a plot record returned by the repository layer."""

    def __init__(
        self,
        plot_id: int,
        plot_tag: str,
        location_description: str | None,
        size_sq_ft: int | None,
        is_raised_bed: bool = False,
        status_id: int | None = None
    ):
        """Initialize a plot model with location and size details.

        Args:
            plot_id: Unique plot identifier.
            plot_tag: Unique tag for the plot.
            location_description: Description of the plot's location.
            size_sq_ft: Size of the plot in square feet.
            is_raised_bed: Flag indicating if the plot is a raised bed (defaults to False).
            status_id: Optional ID referencing the plot status (defaults to None).

        Returns:
            None: Constructor initializes instance state.
        """
        self._plot_id = plot_id
        self._plot_tag = plot_tag
        self._location_description = location_description
        self._size_sq_f = size_sq_ft
        self._is_raised_bed = is_raised_bed
        self._status_id = status_id

    @property
    def plot_id(self) -> int:
        """Return the unique identifier for the plot.

        Returns:
            int: plot identifier.
        """
        return self._plot_id

    @property
    def plot_tag(self) -> str:
        """Return the unique tag for the plot.

        Returns:
            str: plot tag.
        """
        return self._plot_tag

    @property
    def location_description(self) -> str | None:
        """Return the location description for the plot.

        Returns:
            str | None: plot location.
        """
        return self._location_description

    @property
    def size_sq_ft(self) -> int | None:
        """Return the size of the plot.

        Returns:
            int | None: plot size.
        """
        return self._size_sq_f
    
    @property
    def is_raised_bed(self) -> bool:
        """Return True if the plot is a raised bed, and false otherwise.

        Returns:
            bool: whether plot has raised bed or not.
        """
        return self._is_raised_bed

    @property
    def status_id(self) -> int | None:
        """Return the optional ID referencing the plot status.

        Returns:
            int | None: optional status id.
        """
        return self._status_id
