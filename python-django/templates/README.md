# {{name}}

## Setup

1. Build the docker image.

   ```bash
   make docker
   ```

2. Run the docker container.

   ```bash
   make docker-run
   ```

## Development

1. Activate a virtual environment.

   **Linux/MacOS**

   ```bash
   source venv/bin/activate
   ```

   **Windows** (*PowerShell*)

   ```cmd
   .\venv\Scripts\activate;
   ```

2. Install dependencies.

   ```bash
   pip install -r requirements.txt
   ```

3. Run schema migrations.

   ```bash
   python manage.py migrate
   ```

4. Run Tests.

   ```bash
   python manage.py test
   ```

5. Run the development server.

   ```bash
   python manage.py runserver
   ```
