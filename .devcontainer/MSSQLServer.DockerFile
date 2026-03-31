# Use the official Microsoft SQL Server 2025 image from the Docker Hub
FROM mcr.microsoft.com/mssql/server:2025-latest

# Set environment variables https://hub.docker.com/r/microsoft/mssql-server
ENV ACCEPT_EULA=Y
ENV MSSQL_SA_PASSWORD=sup3rs3cur3P@ssword
ENV MSSQL_PID=Express


# Create a directory for SQL Server backups
RUN mkdir /var/opt/mssql/backup

USER root
RUN apt-get -y update 
RUN apt-get install -y curl

# Create the directory with root user and change ownership (before writing to .bashrc)
RUN mkdir -p /home/mssql && chown mssql:mssql /home/mssql

USER mssql

RUN curl -L -o /var/opt/mssql/backup/wwi.bak 'https://github.com/Microsoft/sql-server-samples/releases/download/wide-world-importers-v1.0/WideWorldImporters-Full.bak'

# Copy SQL files to a directory in the container
COPY sql-init/createCC520.sql /home/mssql/sql-init/
COPY sql-init/worldWideImporters.sql /home/mssql/sql-init/


# Copy the entrypoint script
COPY entrypoint.sh /usr/src/app/entrypoint.sh

USER root
# Make sure the setup script is in Unix format
RUN sed -i 's/\r$//' /usr/src/app/entrypoint.sh
# Grant execution rights to the entrypoint script
RUN chmod +x /usr/src/app/entrypoint.sh
USER mssql
# Expose the SQL Server port
EXPOSE 1433

# Run the entrypoint script
ENTRYPOINT ["/usr/src/app/entrypoint.sh"]