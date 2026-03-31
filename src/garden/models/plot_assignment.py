"""
Data model for garden plot assignments.

Author: Josh Weese
"""
class PlotAssignment:
    """Represent a gardener's assignment to a plot."""

    def __init__(self, assignment_id: int, plot_id: int, gardener_id: int, start_date: str,
                 end_date: str | None, notes: str | None, plot_tag: str | None, location_description: str | None,
                 size_sq_ft: int | None, is_raised_bed: bool | None, status: str | None):
        """Initialize a plot assignment model with assignment and plot details.

        Args:
            assignment_id: Unique assignment identifier.
            plot_id: Plot identifier.
            gardener_id: Gardener identifier.
            start_date: Assignment start date.
            end_date: Optional assignment end date.
            notes: Optional assignment notes.
            plot_tag: Optional plot tag (e.g., A1, B2).
            location_description: Optional plot location description.
            size_sq_ft: Optional plot size in square feet.
            is_raised_bed: Optional raised-bed flag.
            status: Optional plot status.

        Returns:
            None: Constructor initializes instance state.
        """
        self._assignment_id = assignment_id
        self._plot_id = plot_id
        self._gardener_id = gardener_id
        self._start_date = start_date
        self._end_date = end_date
        self._notes = notes
        self._plot_tag = plot_tag
        self._location_description = location_description
        self._size_sq_ft = size_sq_ft
        self._is_raised_bed = is_raised_bed
        self._status = status

    @property
    def assignment_id(self):
        """Return the assignment identifier.

        Returns:
            int: Assignment identifier.
        """
        return self._assignment_id

    @property
    def plot_id(self):
        """Return the plot identifier.

        Returns:
            int: Plot identifier.
        """
        return self._plot_id

    @property
    def gardener_id(self):
        """Return the gardener identifier.

        Returns:
            int: Gardener identifier.
        """
        return self._gardener_id

    @property
    def start_date(self):
        """Return the assignment start date.

        Returns:
            str: Start date value.
        """
        return self._start_date

    @property
    def end_date(self):
        """Return the optional assignment end date.

        Returns:
            str | None: End date when present.
        """
        return self._end_date

    @property
    def notes(self):
        """Return optional notes for the assignment.

        Returns:
            str | None: Assignment notes when present.
        """
        return self._notes

    @property
    def plot_tag(self):
        """Return the optional plot tag.

        Returns:
            str | None: Plot tag (e.g., A1, B2) when present.
        """
        return self._plot_tag

    @property
    def location_description(self):
        """Return the optional plot location description.

        Returns:
            str | None: Plot location description when present.
        """
        return self._location_description

    @property
    def size_sq_ft(self):
        """Return the optional plot size in square feet.

        Returns:
            int | None: Plot size when present.
        """
        return self._size_sq_ft

    @property
    def is_raised_bed(self):
        """Return whether the plot is a raised bed.

        Returns:
            bool | None: Raised-bed flag when present.
        """
        return self._is_raised_bed

    @property
    def status(self):
        """Return the optional plot status.

        Returns:
            str | None: Plot status when present.
        """
        return self._status


