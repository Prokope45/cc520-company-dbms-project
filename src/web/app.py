"""Flask application factory for the project.

Author: GPT-5.4
Reviewed by: Josh Weese
"""

from flask import Flask, jsonify, render_template, request
from src.data_access.RecordNotFoundException import RecordNotFoundException
from src.garden.i_plot_assignment_repository import IPlotAssignmentRepository
from src.garden.i_plot_repository import IPlotRepository
from src.garden.i_gardener_repository import IGardenerRepository


def create_app() -> Flask:
    """Create and configure the Flask application.

    Returns:
        Flask: Configured Flask application instance.
    """
    app = Flask(__name__, template_folder="templates")
    app.config["GARDENER_REPOSITORY"] = None
    app.config["PLOT_ASSIGNMENT_REPOSITORY"] = None
    app.config["PLOT_REPOSITORY"] = None

    def get_gardener_repository() -> IGardenerRepository:
        """Return the configured gardener repository or create the SQL implementation.

        Returns:
            IGardenerRepository: Gardener repository implementation for the current request.
        """
        injected_repository: IGardenerRepository | None = app.config.get("GARDENER_REPOSITORY")
        if injected_repository is not None:
            return injected_repository
        from src.garden.sql_gardener_repository import SqlGardenerRepository

        return SqlGardenerRepository(trusted=True)

    def get_plot_assignment_repository() -> IPlotAssignmentRepository:
        """Return the configured assignment repository or create the SQL implementation.

        Returns:
            IPlotAssignmentRepository: Assignment repository implementation for the current request.
        """
        injected_repository: IPlotAssignmentRepository | None = app.config.get("PLOT_ASSIGNMENT_REPOSITORY")
        if injected_repository is not None:
            return injected_repository
        from src.garden.sql_plot_assignment_repository import SqlPlotAssignmentRepository

        return SqlPlotAssignmentRepository(trusted=True)

    def get_plot_repository() -> IPlotRepository:
        """Return the configured plot repository or create the SQL implementation.

        Returns:
            IPlotRepository: Plot repository implementation for the current request.
        """
        injected_repository: IPlotRepository | None = app.config.get("PLOT_REPOSITORY")
        if injected_repository is not None:
            return injected_repository
        from src.garden.sql_plot_repository import SqlPlotRepository

        return SqlPlotRepository(trusted=True)

    @app.get("/")
    def index() -> str:
        """Render the application landing page.

        Returns:
            str: Rendered HTML for the index template.
        """
        return render_template("index.html")

    @app.get("/gardeners")
    def gardeners_page() -> str:
        """Render the gardeners page shell.

        Returns:
            str: Rendered HTML for the gardeners template.
        """
        return render_template("gardeners.html")

    @app.get("/api/gardeners")
    def retrieve_gardeners():
        """Return all gardeners as JSON payloads.

        Returns:
            Response: JSON response containing gardener records.
        """
        repository = get_gardener_repository()
        gardeners = repository.retrieve_gardeners() or []
        return jsonify(
            [
                {
                    "gardenerId": gardener.gardener_id,
                    "firstName": gardener.first_name,
                    "lastName": gardener.last_name,
                    "phone": gardener.phone,
                    "email": gardener.email,
                    "joinDate": str(gardener.join_date) if gardener.join_date is not None else None,
                }
                for gardener in gardeners
            ]
        )

    @app.get("/plots")
    def plots_page() -> str:
        """Render the plots page shell.

        Returns:
            str: Rendered HTML for the plots template.
        """
        return render_template("plots.html")

    @app.get("/api/gardeners/<int:gardener_id>/assignments")
    def retrieve_assignments_for_gardener(gardener_id: int):
        """Return a gardener's details together with plot assignments as JSON.

        Args:
            gardener_id: Gardener identifier from the route path.

        Returns:
            Response: JSON response containing gardener and assignment data or a 404 error payload.
        """
        gardener_repository = get_gardener_repository()
        assignment_repository = get_plot_assignment_repository()
        try:
            gardener = gardener_repository.fetch_gardener(gardener_id)
            if gardener is None:
                raise RecordNotFoundException(f"Gardener {gardener_id} was not found.")
        except RecordNotFoundException:
            return jsonify({"error": f"Gardener {gardener_id} was not found."}), 404

        assignments = assignment_repository.get_assignments(gardener_id) or []
        return jsonify(
            {
                "gardener": {
                    "gardenerId": gardener.gardener_id,
                    "firstName": gardener.first_name,
                    "lastName": gardener.last_name,
                    "phone": gardener.phone,
                    "email": gardener.email,
                    "joinDate": str(gardener.join_date) if gardener.join_date is not None else None,
                },
                "assignments": [
                    {
                        "assignmentId": assignment.assignment_id,
                        "plotId": assignment.plot_id,
                        "gardenerId": assignment.gardener_id,
                        "startDate": str(assignment.start_date),
                        "endDate": str(assignment.end_date) if assignment.end_date is not None else None,
                        "notes": assignment.notes,
                        "plotTag": assignment.plot_tag,
                        "locationDescription": assignment.location_description,
                        "sizeSqFt": assignment.size_sq_ft,
                        "isRaisedBed": assignment.is_raised_bed,
                        "status": assignment.status,
                    }
                    for assignment in assignments
                ],
            }
        )

    @app.get("/api/plots")
    def retrieve_plots():
        """Return all garden plots as JSON payloads."""
        repository = get_plot_repository()
        plots = repository.retrieve_plots()
        extracted_plots = []

        for plot in plots:
            plot_contents = {
                "plotId": plot.plot_id,
                "plotTag": plot.plot_tag,
                "locationDescription": plot.location_description,
                "sizeSqFt": plot.size_sq_ft,
                "isRaisedBed": plot.is_raised_bed,
                "status": "Unknown"
            }
            status = repository.retrieve_plot_status(plot.plot_id)
            if status is not None:
                plot_contents["status"] = status.status_name
            extracted_plots.append(plot_contents)
        return jsonify(extracted_plots)

    @app.get("/api/plots")
    def retrieve_available_plots():
        """Return all garden plots as JSON payloads."""
        repository = get_plot_repository()
        plots = repository.retrieve_available_plots()
        extracted_plots = []

        for plot in plots:
            plot_contents = {
                "plotId": plot.plot_id,
                "plotTag": plot.plot_tag,
                "locationDescription": plot.location_description,
                "sizeSqFt": plot.size_sq_ft,
                "isRaisedBed": plot.is_raised_bed,
                "status": "Unknown"
            }
            status = repository.retrieve_plot_status(plot.plot_id)
            if status is not None:
                plot_contents["status"] = status.status_name
            extracted_plots.append(plot_contents)
        return jsonify(extracted_plots)

    @app.put("/api/plots/<int:plot_id>/status")
    def update_plot_status(plot_id: int):
        """Update the status of a specific plot."""
        repository = get_plot_repository()
        data = request.get_json()
        new_status_id = int(data.get("status"))
        try:
            success = repository.update_plot_status(plot_id, new_status_id)

            if success:
                # Fetch updated plot to return the new status name
                updated_plot = next((p for p in repository.retrieve_plots() if p.plot_id == plot_id), None)
                if updated_plot:
                    status = repository.retrieve_plot_status(updated_plot.plot_id)
                    return jsonify({
                        "message": f"Plot {plot_id} status successfully updated to {'maintenance' if new_status_id == 2 else 'active'}.",
                        "newStatus": status.status_name
                    })
                else:
                     return jsonify({"error": "Status update succeeded, but could not retrieve updated plot details."}), 500
            else:
                return jsonify({"error": f"Failed to update status for plot {plot_id}."}), 500

        except Exception as e:
            # Catch any other unexpected errors during repository interaction
            return jsonify({"error": str(e)}), 500

    return app
