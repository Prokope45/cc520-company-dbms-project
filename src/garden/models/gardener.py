"""
Data model for a gardener.
Author: Josh Weese
"""
class Gardener:
    """Represent a gardener record returned by the repository layer."""

    def __init__(self, gardener_id: int, first_name: str, last_name: str, phone: str | None,
                 email: str | None, join_date: str | None):
        """Initialize a gardener model with contact and membership details.

        Args:
            gardener_id: Unique gardener identifier.
            first_name: Gardener first name.
            last_name: Gardener last name.
            phone: Optional phone number.
            email: Optional email address.
            join_date: Optional join date.

        Returns:
            None: Constructor initializes instance state.
        """
        self._gardener_id = gardener_id
        self._first_name = first_name
        self._last_name = last_name
        self._phone = phone
        self._email = email
        self._join_date = join_date

    @property
    def gardener_id(self):
        """Return the unique identifier for the gardener.

        Returns:
            int: Gardener identifier.
        """
        return self._gardener_id

    @property
    def first_name(self):
        """Return the gardener's first name.

        Returns:
            str: First name value.
        """
        return self._first_name

    @property
    def last_name(self):
        """Return the gardener's last name.

        Returns:
            str: Last name value.
        """
        return self._last_name

    @property
    def email(self):
        """Return the gardener's email address.

        Returns:
            str | None: Email address value when present.
        """
        return self._email

    @property
    def phone(self):
        """Return the gardener's phone number.

        Returns:
            str | None: Phone number when present.
        """
        return self._phone

    @property
    def join_date(self):
        """Return the gardener's join date.

        Returns:
            str | None: Join date when present.
        """
        return self._join_date
