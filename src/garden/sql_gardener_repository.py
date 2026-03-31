"""SQL Server implementation of the IGardenerRepository interface.

Author: Josh Weese

"""

from typing import Optional, List, Any

from src.data_access.RecordNotFoundException import RecordNotFoundException
from src.data_access.sql_command_executor import SqlCommandExecutor
from src.garden.i_gardener_repository import IGardenerRepository
from src.garden.models.gardener import Gardener


class SqlGardenerRepository(IGardenerRepository):
    """Persist and retrieve gardeners through SQL Server stored procedures."""

    def __init__(self, server: str = "", database: str = "", trusted: bool = True):
        """Initialize the repository with a SQL command executor.

        Args:
            server: Optional SQL Server host override.
            database: Optional database name override.
            trusted: Whether trusted authentication should be used.

        Returns:
            None: Constructor initializes repository state.
        """
        self.executor = SqlCommandExecutor(server=server, database=database, trusted=trusted)

    def create_gardener(self, first_name: str, last_name: str, phone: str | None,
                        email: str | None, join_date: str | None) -> Optional[Gardener]:
        """Create a new gardener record and return the resulting model.

        Args:
            first_name: Gardener first name.
            last_name: Gardener last name.
            phone: Optional phone number.
            email: Optional email address.
            join_date: Optional join date.

        Returns:
            Optional[Gardener]: Created gardener model, or ``None`` when no output row is returned.
        """
        if first_name is None or first_name == "":
            raise ValueError("First name cannot be empty.")
        if last_name is None or last_name == "":
            raise ValueError("Last name cannot be empty.")
        sp_name = "Garden.CreateGardener"
        inp_param_names = ['FirstName', 'LastName', 'Phone', 'Email', 'JoinDate']
        inp_param_values = [first_name, last_name, phone, email, join_date]
        out_param = {
            'sp_local': ['GardenerId'],
            'sp_local_types': ['int'],
            'sp_out': ["GardenerId"],
        }

        results = self.executor.execute_stored_procedure(sp_name,
                                                         input_param_names=inp_param_names,
                                                         input_param_values=inp_param_values,
                                                         output_param=out_param)
        if len(results) == 1:
            return Gardener(results[0].GardenerId_var, first_name, last_name, phone, email, join_date)
        else:
            return None

    def retrieve_gardeners(self) -> Optional[List[Gardener]]:
        """Return all gardeners currently stored in the backing database.

        Returns:
            Optional[List[Gardener]]: Gardener models when rows exist; otherwise ``None``.
        """
        sp_name = "Garden.RetrieveGardeners"
        results = self.executor.execute_stored_procedure(sp_name)

        if len(results) >= 1:
            return self.translate_gardeners(results)
        else:
            return None

    def fetch_gardener(self, gardener_id: int) -> Optional[Gardener]:
        """Return one gardener by identifier or raise when the record is absent.

        Args:
            gardener_id: Gardener identifier.

        Returns:
            Optional[Gardener]: Matching gardener model.
        """
        sp_name = "Garden.FetchGardener"
        inp_param_names = ['GardenerId']
        inp_param_values = [gardener_id]

        results = self.executor.execute_stored_procedure(sp_name,
                                                         input_param_names=inp_param_names,
                                                         input_param_values=inp_param_values)
        if len(results) == 1:
            return self.translate_gardener(results[0])
        else:
            raise RecordNotFoundException(gardener_id)

    def get_gardener_by_email(self, email: str) -> Optional[Gardener]:
        """Return one gardener matching the supplied email address.

        Args:
            email: Email address to search by.

        Returns:
            Optional[Gardener]: Matching gardener model, or ``None`` when absent.
        """
        sp_name = "Garden.GetGardenerByEmail"
        inp_param_names = ['Email']
        inp_param_values = [email]

        results = self.executor.execute_stored_procedure(sp_name,
                                                         input_param_names=inp_param_names,
                                                         input_param_values=inp_param_values)
        if len(results) == 1:
            return self.translate_gardener(results[0])
        else:
            return None

    def translate_gardener(self, row: Any) -> Gardener:
        """Map a database row object into a ``Gardener`` model.

        Args:
            row: Database row object with gardener fields.

        Returns:
            Gardener: Translated gardener model.
        """
        return Gardener(row.GardenerId, row.FirstName, row.LastName, row.Phone, row.Email, row.JoinDate)

    def translate_gardeners(self, rows: List[Any]) -> List[Gardener]:
        """Map multiple database rows into a list of ``Gardener`` models.

        Args:
            rows: Database row objects with gardener fields.

        Returns:
            List[Gardener]: Translated gardener models.
        """
        gardeners = []
        for row in rows:
            gardeners.append(self.translate_gardener(row))
        return gardeners
