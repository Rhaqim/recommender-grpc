FROM python:3.8-slim

# Set the working directory
WORKDIR /APP

# Copy the current directory contents into the container at /APP
COPY . /APP

# Install the required packages
RUN pip install --no-cache-dir -r requirements.txt

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT [ "python", "app.py" ]