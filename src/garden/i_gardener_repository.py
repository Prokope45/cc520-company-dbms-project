"""
Repository interface for gardener-related data operations.

Author: Josh Weese
"""
from abc import ABC, abstractmethod
from typing import Optional, List

from src.garden.models.gardener import Gardener


class IGardenerRepository(ABC):
    """Define the contract for gardener repository implementations."""

    @abstractmethod
    def create_gardener(self, first_name: str, last_name: str, phone: str | None,
                        email: str | None, join_date: str | None) -> Optional[Gardener]:
        """Create a gardener record and return the created model when successful.

        Args:
            first_name: Gardener first name.
            last_name: Gardener last name.
            phone: Optional phone number.
            email: Optional email address.
            join_date: Optional join date.

        Returns:
            Optional[Gardener]: Created gardener model, or ``None`` when creation fails.
        """
        pass

    @abstractmethod
    def retrieve_gardeners(self) -> Optional[List[Gardener]]:
        """Return all known gardeners or ``None`` when no rows exist.

        Returns:
            Optional[List[Gardener]]: Gardener models when rows exist; otherwise ``None``.
        """
        pass

    @abstractmethod
    def fetch_gardener(self, gardener_id: int) -> Optional[Gardener]:
        """Return a single gardener by identifier.

        Args:
            gardener_id: Gardener identifier.

        Returns:
            Optional[Gardener]: Matching gardener model.
        """
        pass

    @abstractmethod
    def get_gardener_by_email(self, email: str) -> Optional[Gardener]:
        """Return a gardener matching the supplied email address.

        Args:
            email: Email address to search by.

        Returns:
            Optional[Gardener]: Matching gardener model, or ``None`` when absent.
        """
        pass

