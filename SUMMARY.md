## Manually Written Code

I manually wrote code in:
- Garden.RetrievePlots.sql
- Garden.RetrieveAvailablePlots.sql
- Garden.RetrievePlotStatus.sql
- Garden.UpdatePlotStatus.sql
- i_plot_repository.py
- sql_plot_repository.py
- plot.py
- plot_status.py

I also closely review the AI code and updated it where it got things wrong. In particular, the API endpoints
it created did not properly supply data to the backend repository to be processed in the query. So I ended up
fixing a lot of bugs it created.

## AI Written Code

I used RooCode connected to LM Studio running Gemma-4-e2b. I later switched to qwen/quen3.5-9b since the Gemma model kept failing at writing the HTML files correctly, preferring to use djinja templates rather than writing
javascript and updating the elements directly from the javascript code.

The local model mostly wrote:
- plots.html
- assign_plots.html
- some functions in app.py (although I made major updates to fix some bugs).
- some unit tests in test_web_app.py

I added [RooCodeTasks](RooCodeTasks) which contains all the Markdown files of the tasks I gave the agent.

## Notice for Late Submission

I am submitting this assignment on time, but not entirely finished due to time constraints. I participated in DataFest and had no available time to dedicate to working on this assignment over the weekend. I also have an evening class on Mondays, so I apologize for not submitting my assignment with all the required features completed.
