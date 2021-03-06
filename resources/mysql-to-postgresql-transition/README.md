# MySQL to Postgresql transition

1. get pgloader version 3.6.1 (or above?).
   For Arch Linux, it exists in AUR.
   For Debian or others, compile it through docker. `git clone https://github.com/dimitri/pgloader.git && cd pgloader && git checkout v3.6.1 && docker build -t pgloader:3.6.1 .`
2. modify transition files (stage1, stage2, stage3) to fit your database needs
3. run pgloader transitions files (ensure you've set up permissions for docker in `pg_hba.conf`):
   `docker run --rm --name pgloader --net host --mount type=bind,source="$PWD",destination=/tmp/transition -p 5432:5432 pgloader:3.6.1 pgloader /tmp/transition/db-transfer-stage1`
   `docker run --rm --name pgloader --net host --mount type=bind,source="$PWD",destination=/tmp/transition -p 5432:5432 pgloader:3.6.1 pgloader /tmp/transition/db-transfer-stage2`
4. install the python virtual environment: `python3 -m venv venv`
5. activate the virtual environment: `source venv/bin/activate`
6. install sql clients: `pip install psycopg2 pymysql`
7. run the "stage3.py" python script: `./stage3.py`
